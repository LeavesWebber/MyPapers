package logic

import (
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
)

// GetAllReviews 查询审核列表
func GetAllReviews(filter string, userId uint) (out []*response.GetPaper, err error) {
	// 1. 去reviews表找自己对应的审核paperId
	paperIds, err := mysql.GetReviewIds(userId)
	if err != nil {
		return
	}
	for _, v := range paperIds {
		// 2. 去papers表找对应的paper
		paper, err := mysql.GetPaper(v)
		if err != nil {
			return nil, err
		}
		if filter == "Reviewed" {
			if paper.Status != "UnReview" && paper.Status != "InReview" {
				// 3. 去reviews表找对应的reviews
				reviews, err := mysql.GetReviews(v)
				if err != nil {
					return nil, err
				}
				out = append(out, &response.GetPaper{
					Paper: paper,
				})
				// 4. 去users表找对应的user
				for _, v := range reviews {
					user, err := mysql.GetUserInfoByID(v.ReviewerId)
					if err != nil {
						return nil, err
					}
					out[len(out)-1].ReviewInfos = append(out[len(out)-1].ReviewInfos, &response.ReviewInfo{
						ReviewerName: user.Username,
						Comment:      v.Comment,
						Status:       v.Status,
					})
				}
			}
		} else {
			if paper.Status == "UnReview" || paper.Status == "InReview" {

				out = append(out, &response.GetPaper{
					Paper: paper,
				})
				// 3. 去reviews表找对应的reviews
				reviews, err := mysql.GetReviews(v)
				if err != nil {
					return nil, err
				}
				// 4. 去users表找对应的user
				for _, v := range reviews {
					user, err := mysql.GetUserInfoByID(v.ReviewerId)
					if err != nil {
						return nil, err
					}
					out[len(out)-1].ReviewInfos = append(out[len(out)-1].ReviewInfos, &response.ReviewInfo{
						ReviewerName: user.Username,
						Comment:      v.Comment,
						Status:       v.Status,
					})
				}
			}
		}
	}
	for _, v := range out {
		// 通过conferenceId或者journalId找到自己的level
		if v.Paper.ConferenceId != 0 {
			userConference, err := mysql.GetUserConferenceByIdAndUserId(v.Paper.ConferenceId, userId)
			if err != nil {
				return nil, err
			}
			if userConference != nil {
				v.Level = userConference.Level
			}
		} else {
			userJournal, err := mysql.GetUserJournalByIdAndUserId(v.Paper.JournalId, userId)
			if err != nil {
				return nil, err
			}
			if userJournal != nil {
				v.Level = userJournal.Level
			}
		}
	}
	return
}

// AllotReviewers 分配审核人
func AllotReviewers(in *request.AllotReviewers) (err error) {
	// 如果paper已审核，则不允许分配
	paper, err := mysql.GetPaper(in.PaperId)
	if err != nil {
		return err
	}
	if paper.Status == "Reviewed" {
		return
	}
	// 先通过审核人name找到对应的id
	reviewIds, err := mysql.GetReviewIdsByName(in.ReviewerNames)
	if err != nil {
		return err
	}
	if len(reviewIds) != len(in.ReviewerNames) {
		return global.ErrorUserNotExist
	}
	// 再将paperId和reviewerId存入reviews表
	return mysql.InsertReview(in.PaperId, reviewIds)
}

// SubmitReview 提交审核
func SubmitReview(in *request.SubmitReview, currentId uint) (out *tables.Review, err error) {
	// 如果paper已审核，则不允许提交
	paper, err := mysql.GetPaper(in.PaperId)
	if err != nil {
		return nil, err
	}
	if paper.Status == "Reviewed" {
		return nil, global.ErrPaperReviewed
	}
	var level string
	if paper.ConferenceId != 0 {
		// 从user_conferences表中找到userId对应的position
		userConference, err := mysql.GetUserConferenceByIdAndUserId(paper.ConferenceId, currentId)
		if err != nil {
			return nil, err
		}
		level = userConference.Level
	}
	if paper.JournalId != 0 {
		// 从user_journals表中找到userId对应的position
		userJournal, err := mysql.GetUserJournalByIdAndUserId(paper.JournalId, currentId)
		if err != nil {
			return nil, err
		}
		level = userJournal.Level
	}
	if level == "president" {
		if in.Status == "Accept" {
			// 获取审核文章的路径
			paper, err = mysql.GetPaper(in.PaperId)

			//把文件存入ipfs并且返回cid
			cid, err := saveToIPFS(paper.Filepath)
			if err != nil {
				return nil, err
			}
			//cid := "cidcidcidcid"
			// 更新cid
			if err = mysql.SetPaperCid(in.PaperId, cid); err != nil {
				return nil, err
			}
		}
		// 更新paper的status
		if err = mysql.SetPaperStatus(in.PaperId, in.Status); err != nil {
			return nil, err
		}
	} else {
		// 更新paper的status
		if err = mysql.SetPaperStatus(in.PaperId, "InReview"); err != nil {
			return nil, err
		}
	}
	// 将in存入reviews表
	return mysql.SubmitReview(in, currentId)
}

// saveToIPFS 把文件存入ipfs并且返回cid
func saveToIPFS(filePath string) (cid string, err error) {
	// 读取本地文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	// 连接到ipfs
	sh := shell.NewShell(global.MPS_CONFIG.IPFS.Host + ":" + global.MPS_CONFIG.IPFS.Port)
	//sh := shell.NewShell("127.0.0.1:5001")
	// 将文件添加到 IPFS，并返回文件的cid
	return sh.Add(bytes.NewReader(data))
}

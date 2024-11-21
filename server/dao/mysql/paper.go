package mysql

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"strconv"
	"time"
)

const (
	President = "president"
)

// SubmitPaper 投稿
func SubmitPaper(in *tables.Paper) (out *tables.Paper, err error) {
	if in.ID != 0 { // 如果是修改稿版本
		// 记录原来paperId
		oldPaperId := in.ID
		// 插入papers表
		in.ID = 0
		if err = global.MPS_DB.Create(in).Error; err != nil {
			return
		}
		// 更新原来的user_paper表行数据的oldVersion标志
		for _, user := range in.Users {
			if err = global.MPS_DB.Model(&tables.UserPaper{}).Where("user_id = ? and paper_id = ?", user.ID, oldPaperId).Update("old_version", true).Error; err != nil {
				return
			}
		}
		// 更新reviews表中的oldVersion标志
		if err = global.MPS_DB.Model(&tables.Review{}).Where("paper_id = ?", oldPaperId).Update("old_version", true).Error; err != nil {
			return
		}
		// 查出原来的paperId对应的reviewerId，插入reviews表
		var review []*tables.Review
		if err = global.MPS_DB.Model(&tables.Review{}).Select("reviewer_id").Where("paper_id = ?", oldPaperId).Find(&review).Error; err != nil {
			return
		}
		for _, v := range review {
			v.PaperId = in.ID
		}
		if err = global.MPS_DB.Create(review).Error; err != nil {
			return
		}
	} else { // 如果是新稿
		// 插入papers表
		if err = global.MPS_DB.Create(in).Error; err != nil {
			return
		}
		var review []*tables.Review
		if in.JournalId != 0 { // 如果是期刊投稿
			// 从user_journal表中获取主编的user_id
			var userJournal []tables.UserJournal
			if err = global.MPS_DB.Select("user_id").Where("journal_id = ? and level = ?", in.JournalId, President).Find(&userJournal).Error; err != nil {
				return
			}
			review = make([]*tables.Review, len(userJournal))
			for i, v := range userJournal {
				review[i] = new(tables.Review)
				review[i].ReviewerId = v.UserId
				review[i].PaperId = in.ID
			}
			if err = global.MPS_DB.Create(review).Error; err != nil {
				return
			}
		} else { // 如果是会议投稿
			// 从user_conference表中获取主编的user_id
			var userConference []tables.UserConference
			if err = global.MPS_DB.Select("user_id").Where("conference_id = ? and level = ?", in.ConferenceId, President).Find(&userConference).Error; err != nil {
				return
			}
			review = make([]*tables.Review, len(userConference))
			for i, v := range userConference {
				review[i] = new(tables.Review)
				review[i].ReviewerId = v.UserId
				review[i].PaperId = in.ID
			}
			if err = global.MPS_DB.Create(review).Error; err != nil {
				return
			}
		}
	}
	return in, nil
}

// GetPaper 投稿详情
func GetPaper(paperID uint) (detail *tables.Paper, err error) {
	detail = new(tables.Paper)
	// 获取投稿信息,预加载users
	err = global.MPS_DB.Preload("Users").Where("id = ?", paperID).First(detail).Error
	return
}

// GetAllSelfPapers 查询自己的投稿列表
func GetAllSelfPapers(userId uint) (out []*tables.Paper, err error) {
	// 先从中间表找到user的paper
	var userPaper []*tables.UserPaper
	if err = global.MPS_DB.Where("user_id = ? and (old_version is null or old_version = 'false')", userId).Find(&userPaper).Error; err != nil {
		return
	}
	paperIds := make([]uint, len(userPaper))
	for k, v := range userPaper {
		paperIds[k] = v.PaperId
	}
	if err = global.MPS_DB.Preload("Users").Where("id in (?)", paperIds).Find(&out).Error; err != nil {
		return
	}
	return
}

// GetAllPapers 查询投稿列表
func GetAllPapers(journalId, conferenceId int) (out []*tables.Paper, err error) {
	if journalId != 0 {
		err = global.MPS_DB.Where("journal_id = ?", journalId).Find(&out).Error
	} else {
		err = global.MPS_DB.Where("conference_id = ?", conferenceId).Find(&out).Error
	}
	return
}

// UpdatePaper 更新投稿
func UpdatePaper(usersId []uint, in *request.UpdatePaper, paper *tables.Paper) (out *tables.Paper, err error) {
	// 更新user_paper表
	// 先删掉原来的
	if err = global.MPS_DB.Where("paper_id = ?", in.PaperId).Delete(&tables.UserPaper{}).Error; err != nil {
		return
	}
	// 再插入新的
	userPaper := make([]*tables.UserPaper, len(usersId))
	for k, v := range usersId {
		userPaper[k] = new(tables.UserPaper)
		userPaper[k].UserId = v
		userPaper[k].PaperId = in.PaperId
	}
	if err = global.MPS_DB.Create(userPaper).Error; err != nil {
		return
	}
	// 更新papers表
	out = new(tables.Paper)
	err = global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", in.PaperId).Updates(paper).First(out).Error
	return
}

// UpdatePaper2 更新投稿
func UpdatePaper2(userId uint, in *request.UpdatePaper, paper *tables.Paper) (out *tables.Paper, err error) {
	// 更新user_paper表
	// 先删掉原来的
	if err = global.MPS_DB.Where("paper_id = ?", in.PaperId).Delete(&tables.UserPaper{}).Error; err != nil {
		return
	}
	// 再插入新的
	userPaper := &tables.UserPaper{
		UserId:  userId,
		PaperId: in.PaperId,
	}
	if err = global.MPS_DB.Create(userPaper).Error; err != nil {
		return
	}
	// 更新papers表
	out = new(tables.Paper)
	err = global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", in.PaperId).Updates(paper).First(out).Error
	return
}

// DeletePaper 删除投稿
func DeletePaper(id uint) (err error) {
	tx := global.MPS_DB.Begin()
	if err = tx.Where("id = ?", id).Delete(&tables.Paper{}).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Where("paper_id = ?", id).Delete(&tables.Review{}).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Where("paper_id = ?", id).Delete(&tables.UserPaper{}).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

// SetPaperStatus 修改投稿状态
func SetPaperStatus(id uint, status string) error {
	return global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", id).Update("status", status).Error
}

// SetPaperCid 修改投稿cid
func SetPaperCid(id uint, cid string) error {
	return global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", id).Update("cid", cid).Error
}

// GetAllAcceptPapers 查询所有已通过的投稿
func GetAllAcceptPapers() (out []*tables.Paper, err error) {
	return out, global.MPS_DB.Preload("Users").Where("status = ?", "Published").Find(&out).Error
}

// GetAllAcceptPapersByJournalAndTime 按期刊和时间查询已经审核通过的投稿列表
func GetAllAcceptPapersByJournalAndTime(journalId int, startTime, endTime time.Time) (out []*tables.Paper, err error) {
	return out, global.MPS_DB.Preload("Users").Where("journal_id = ? and status = ? and created_at between ? and ?", journalId, "Published", startTime, endTime).Find(&out).Error
}

// GetAllAcceptPapersByConferenceAndTime 按会议和时间查询已经审核通过的投稿列表
func GetAllAcceptPapersByConferenceAndTime(conferenceId int, startTime, endTime time.Time) (out []*tables.Paper, err error) {
	return out, global.MPS_DB.Preload("Users").Where("conference_id = ? and status = ? and created_at between ? and ?", conferenceId, "Published", startTime, endTime).Find(&out).Error
}

// GetPaperCountToday 查询数据库中当天的稿件数量
func GetPaperCountToday(versionPre string) (count int64, err error) {
	// 20231221000101中查询20231221xxxx00的数量
	return count, global.MPS_DB.Unscoped().Model(&tables.Paper{}).Where("version_id like ?", versionPre+"%00").Count(&count).Error
}

// GetPaperVersions 获取投稿的所有版本
func GetPaperVersions(versionId uint) (out []*tables.Paper, err error) {
	// 截取前12位
	versionPre := strconv.Itoa(int(versionId))[:12]
	return out, global.MPS_DB.Where("version_id like ?", versionPre+"%").Find(&out).Error
}

// SetPaperInfo 设置投稿信息
func SetPaperInfo(paperId uint, uri, url, cid, metadataUri string) error {
	return global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", paperId).Updates(map[string]interface{}{"image_uri": uri, "image_url": url, "image_cid": cid, "json_uri": metadataUri}).Error
}

// PublishPaper 发布投稿
func PublishPaper(in *request.PublishPaper) (err error) {
	return global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", in.PaperId).Updates(map[string]interface{}{"status": "Published", "download_price": in.DownloadPrice, "copyright_trading_price": in.CopyrightTradingTrice, "transaction_hash": in.TransactionHash, "token_id": in.TokenId}).Error
}

// AddPaperViewer 增加投稿可查看者
func AddPaperViewer(in *request.AddPaperViewer) (err error) {
	paperViewers := &tables.PaperViewers{
		PaperId:  in.PaperId,
		ViewerId: in.ViewerId,
	}
	return global.MPS_DB.Create(paperViewers).Error
}

// CheckPaperViewer 检查投稿可查看者
func CheckPaperViewer(paperId, viewerId uint) (isViewer bool, err error) {
	if err = global.MPS_DB.Where("paper_id = ? and viewer_id = ?", paperId, viewerId).First(&tables.PaperViewers{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

// GetMyNFTs 获取我的NFT
func GetMyNFTs(userId uint) (out []*response.GetMyNFTs, err error) {
	var paperIds []int
	if err = global.MPS_DB.Model(&tables.UserPaper{}).Select("paper_id").Where("user_id = ?", userId).Find(&paperIds).Error; err != nil {
		return
	}
	if len(paperIds) == 0 {
		return
	}
	var paper []tables.Paper
	if err = global.MPS_DB.Where("token_id IS NOT NULL").Find(&paper, paperIds).Error; err != nil {
		return
	}
	out = make([]*response.GetMyNFTs, 0)
	for _, v := range paper {
		if v.TokenId == "" {
			continue
		}
		out = append(out, &response.GetMyNFTs{
			PaperId:               v.ID,
			TokenId:               v.TokenId,
			ImageUri:              v.ImageUri,
			ImageUrl:              v.ImageUrl,
			ImageCid:              v.ImageCid,
			CopyrightTradingPrice: v.CopyrightTradingPrice,
			TransactionHash:       v.TransactionHash,
		})
	}
	return
}

// UpdatePrice 更新价格
func UpdatePrice(in *request.UpdatePrice) (err error) {
	return global.MPS_DB.Model(&tables.Paper{}).Where("id = ?", in.PaperId).Updates(map[string]interface{}{"download_price": in.DownloadPrice, "copyright_trading_price": in.CopyrightTradingTrice}).Error
}

// GetNFTInfoByTokenId 根据tokenIds获取NFT信息
func GetNFTInfoByTokenId(tokenIds []int) (out []*response.GetMyNFTs, err error) {
	var paper []tables.Paper
	err = global.MPS_DB.Debug().Where("token_id in (?)", tokenIds).Find(&paper).Error
	out = make([]*response.GetMyNFTs, len(paper))
	for k, v := range paper {
		out[k] = &response.GetMyNFTs{
			PaperId:               v.ID,
			TokenId:               v.TokenId,
			ImageUri:              v.ImageUri,
			ImageUrl:              v.ImageUrl,
			ImageCid:              v.ImageCid,
			CopyrightTradingPrice: v.CopyrightTradingPrice,
			TransactionHash:       v.TransactionHash,
		}
	}
	return
}

// UpdatePaperUserId 修改投稿对应的user_id
func UpdatePaperUserId(paperId, userId uint) (err error) {
	return global.MPS_DB.Model(&tables.UserPaper{}).Where("paper_id = ?", paperId).Update("user_id", userId).Error
}

// GetConferenceOrJournal 获取会议或期刊
func GetConferenceOrJournal(conferenceId, journalId uint) (name string, err error) {
	if conferenceId != 0 {
		var conference tables.Conference
		if err = global.MPS_DB.Select("name").Where("id = ?", conferenceId).First(&conference).Error; err != nil {
			return
		}
		name = conference.Name
	} else {
		var journal tables.Journal
		if err = global.MPS_DB.Select("name").Where("id = ?", journalId).First(&journal).Error; err != nil {
			return
		}
		name = journal.Name
	}
	return
}

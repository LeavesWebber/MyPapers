package logic

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"go.uber.org/zap"
)

func getPaperVersionID() (uint, error) {
	//	version_id格式: 20231221000100,年月日8位,当天的第几个稿件4位，最后两位表示稿件第几个版本
	//	获取当前时间
	now := time.Now()
	//	获取当前时间的年月日
	year, month, day := now.Date()
	//	当前时间的年月日8位
	versionPre := fmt.Sprintf("%d%02d%02d", year, month, day)
	// 	获取当天的第几个稿件
	count, err := mysql.GetPaperCountToday(versionPre)
	if err != nil {
		return 0, err
	}
	//	第几个稿件4位，最后两位表示稿件第几个版本
	countStr := fmt.Sprintf("%04d%02d", count+1, 0)
	//	拼接version_id
	versionID, err := strconv.ParseUint(versionPre+countStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(versionID), nil
}

// SubmitPaper 投稿
func SubmitPaper(c *gin.Context, in *request.SubmitPaper) (out *tables.Paper, err error) {
	//author, _ := json.Marshal(in.Authors)
	//keywords, _ := json.Marshal(in.Keywords)
	//subjectCategory, _ := json.Marshal(in.SubjectCategory)
	// 如果没有id，就说明是新投稿，需要生成version_id格式: 20231221000100,年月日8位,当天的第几个稿件4位，最后两位表示稿件第几个版本
	if in.Id == 0 {
		in.VersionId, err = getPaperVersionID()
		if err != nil {
			return nil, err
		}
	} else {
		// 说明是修改稿件，把version_id+1
		in.VersionId = in.VersionId + 1
	}
	// 根据用户名查找用户id
	authors := utils.StringToSlice(in.Authors)
	usersId, err := mysql.GetUsersIdByNames(authors)
	if err != nil {
		return nil, err
	}
	if len(usersId) != len(authors) {
		return nil, global.ErrorUserNotExist
	}
	users := make([]tables.User, len(usersId))
	for i, v := range usersId {
		users[i] = tables.User{
			MPS_MODEL: global.MPS_MODEL{ID: v},
		}
	}
	// 2. 生成文件名和保存路径
	filename := filepath.Base(in.Data.Filename)
	finalName := fmt.Sprintf("%d_%s", in.VersionId, filename)
	saveFile := filepath.Join("./public/", finalName)
	//saveFile := filepath.Join("./public/", filename)
	// 保存文件
	if err := c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		return nil, err
	}

	// 投稿信息
	paper := &tables.Paper{
		JournalId:    in.JournalId,
		ConferenceId: in.ConferenceId,
		VersionId:    in.VersionId,
		PaperType:    in.PaperType,
		Title:        in.Title,
		Abstract:     in.Abstract,
		//KeyWords:        string(keywords),
		//SubjectCategory: string(subjectCategory),
		KeyWords:        in.Keywords,
		SubjectCategory: in.SubjectCategory,
		//ManuscriptID:       in.ManuscriptID,
		//InformedConsent:    in.InformedConsent,
		//AnimalSubjects:     in.AnimalSubjects,
		CorAuthor: in.CorAuthor,
		//ManuscriptType:     in.ManuscriptType,
		UniqueContribution: in.UniqueContribution,
		BlockAddress:       in.BlockAddress,
		Hash:               in.Hash,
		Filepath:           saveFile,
		Status:             "UnReview", // 默认为未审核
		Users:              users,      // gorm在中间表中自动插入数据
	}
	paper.ManuscriptID = strconv.Itoa(rand.Intn(1000000000))
	if in.Id != 0 {
		paper.ID = in.Id
	}
	// 存入数据库
	return mysql.SubmitPaper(paper)
}

// SubmitPaper2 投稿2作者只有一个
func SubmitPaper2(c *gin.Context, in *request.SubmitPaper, userId uint) (out *tables.Paper, err error) {
	// 如果没有id，就说明是新投稿，需要生成version_id格式: 20231221000100,年月日8位,当天的第几个稿件4位，最后两位表示稿件第几个版本
	if in.Id == 0 {
		in.VersionId, err = getPaperVersionID()
		if err != nil {
			return nil, err
		}
	} else {
		// 说明是修改稿件，把version_id+1
		in.VersionId = in.VersionId + 1
	}
	// 根据用户名查找用户id
	authors := utils.StringToSlice(in.Authors)
	usersId, err := mysql.GetUsersIdByNames(authors)
	if err != nil {
		return nil, err
	}
	if len(usersId) != len(authors) {
		return nil, global.ErrorUserNotExist
	}
	users := make([]tables.User, 1)
	users[0] = tables.User{
		MPS_MODEL: global.MPS_MODEL{ID: userId},
	}
	in.Authors = ""
	for _, uId := range usersId {
		userInfo, err := mysql.GetUserInfoByID(uId)
		if err != nil {
			return nil, err
		}
		in.Authors = in.Authors + "," + userInfo.FirstName + " " + userInfo.LastName
	}
	// 去掉第一个逗号
	in.Authors = in.Authors[1:]
	// 2. 生成文件名和保存路径
	filename := filepath.Base(in.Data.Filename)
	finalName := fmt.Sprintf("%d_%s", in.VersionId, filename)
	saveFile := filepath.Join("./public/", finalName)
	//saveFile := filepath.Join("./public/", filename)
	// 保存文件
	if err := c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		return nil, err
	}

	// 投稿信息
	paper := &tables.Paper{
		JournalId:    in.JournalId,
		ConferenceId: in.ConferenceId,
		VersionId:    in.VersionId,
		PaperType:    in.PaperType,
		Title:        in.Title,
		Authors:      in.Authors,
		Abstract:     in.Abstract,
		//KeyWords:        string(keywords),
		//SubjectCategory: string(subjectCategory),
		KeyWords:        in.Keywords,
		SubjectCategory: in.SubjectCategory,
		//ManuscriptID:       in.ManuscriptID,
		//InformedConsent:    in.InformedConsent,
		//AnimalSubjects:     in.AnimalSubjects,
		CorAuthor: in.CorAuthor,
		//ManuscriptType:     in.ManuscriptType,
		UniqueContribution:   in.UniqueContribution,
		BlockAddress:         in.BlockAddress,
		PaperTransactionHash: in.PaperTransactionAddress,
		Hash:                 in.Hash,
		Filepath:             saveFile,
		Status:               "UnReview", // 默认为未审核
		Users:                users,      // gorm在中间表中自动插入数据
	}
	//paper.ManuscriptID = strconv.Itoa(rand.Intn(1000000000))
	paper.ManuscriptID = strconv.FormatInt(int64(in.VersionId), 10)
	if in.Id != 0 {
		paper.ID = in.Id
	}
	// 存入数据库
	return mysql.SubmitPaper(paper)
}

// GetPaper 投稿详情
func GetPaper(paperID uint) (detail *response.GetPaper, err error) {
	// 获取投稿信息
	detail = new(response.GetPaper)
	if detail.Paper, err = mysql.GetPaper(paperID); err != nil {
		global.MPS_LOG.Error("mysql.Getpaper error", zap.Error(err))
		return
	}
	detail.Paper.Filepath = "http://" + global.MPS_CONFIG.Nginx.Host + ":" + global.MPS_CONFIG.Nginx.Port + "/" + detail.Paper.Filepath
	// 获取各个审核人的审核结果
	reviews, err := mysql.GetReviews(paperID)
	if err != nil {
		global.MPS_LOG.Error("mysql.getreviews error", zap.Error(err))
		return nil, err
	}
	// 获取各个审核人的信息
	for _, v := range reviews {
		user, err := mysql.GetUserInfoByID(v.ReviewerId)
		if err != nil {
			global.MPS_LOG.Error("lmysql.getuserinfobyid error", zap.Error(err))
			return nil, err
		}
		detail.ReviewInfos = append(detail.ReviewInfos, &response.ReviewInfo{
			ReviewerName: user.Username,
			Comment:      v.Comment,
			Status:       v.Status,
		})
	}
	return
}

// GetAllSelfPapers 查询自己的投稿列表
func GetAllSelfPapers(filter string, userId uint) (out []*tables.Paper, err error) {
	// 获取投稿信息
	papers, err := mysql.GetAllSelfPapers(userId)
	if err != nil {
		return
	}
	if filter == "Reviewed" {
		for _, v := range papers {
			if v.Status != "UnReview" && v.Status != "InReview" {
				out = append(out, v)
			}
		}
	} else {
		for _, v := range papers {
			if v.Status == "UnReview" || v.Status == "InReview" {
				out = append(out, v)
			}
		}
	}
	// 获取作者id
	//var authorsIds []uint
	//// 从中间表中获取作者id
	//for _, v := range out {
	//	if authorsIds, err = mysql.GetUserIDsByPaperID(v.ID); err != nil {
	//		return nil, err
	//	}
	//}
	//// 获取作者信息
	//for _, v := range out {
	//	var authors []tables.User
	//	for _, id := range authorsIds {
	//		user, err := mysql.GetUserInfoByID(id)
	//		if err != nil {
	//			return nil, err
	//		}
	//		authors = append(authors, *user)
	//	}
	//	v.Users = authors
	//}
	return
}

// GetAllPapers 查询投稿列表
func GetAllPapers(journalId, conferenceId string) (out []*tables.Paper, err error) {
	journalID, _ := strconv.Atoi(journalId)
	conferenceID, _ := strconv.Atoi(conferenceId)
	return mysql.GetAllPapers(journalID, conferenceID)
}

// UpdatePaper 更新投稿
func UpdatePaper(filePath string, in *request.UpdatePaper) (out *tables.Paper, err error) {
	// 先查询投稿信息
	//sourcePaper, err := mysql.GetPaper(in.PaperId)
	//if err != nil {
	//	return nil, err
	//}
	// 如果投稿的状态不为待审核，则不允许修改
	//if sourcePaper.Status != "UnReview" {
	//	return nil, global.ErrPaperReviewed
	//}
	//author, _ := json.Marshal(in.Authors)
	//keywords, _ := json.Marshal(in.Keywords)
	//subjectCategory, _ := json.Marshal(in.SubjectCategory)

	// 根据用户名查找用户id
	authors := utils.StringToSlice(in.Authors)
	usersId, err := mysql.GetUsersIdByNames(authors)
	if err != nil {
		return nil, err
	}

	//users := make([]tables.User, len(usersId))
	//for i, v := range usersId {
	//	users[i] = tables.User{
	//		MPS_MODEL: global.MPS_MODEL{ID: v},
	//	}
	//}
	// 投稿信息
	paper := &tables.Paper{
		PaperType: in.PaperType,
		Title:     in.Title,
		Abstract:  in.Abstract,
		//KeyWords:        string(keywords),
		//SubjectCategory: string(subjectCategory),
		KeyWords:        in.Keywords,
		SubjectCategory: in.SubjectCategory,
		//ManuscriptID:       in.ManuscriptID,
		//InformedConsent:    in.InformedConsent,
		//AnimalSubjects:     in.AnimalSubjects,
		CorAuthor:          in.CorAuthor,
		ManuscriptType:     in.ManuscriptType,
		UniqueContribution: in.UniqueContribution,
		BlockAddress:       in.BlockAddress,
		Hash:               in.Hash,
		Filepath:           filePath,
		Status:             "UnReview", // 默认为未审核
		//Users:              users,      // gorm在中间表中自动插入数据
	}
	// 存入数据库
	return mysql.UpdatePaper(usersId, in, paper)
}

// UpdatePaper2 更新投稿
func UpdatePaper2(userId uint, filePath string, in *request.UpdatePaper) (out *tables.Paper, err error) {
	// 根据用户名查找用户id
	authors := utils.StringToSlice(in.Authors)
	usersId, err := mysql.GetUsersIdByNames(authors)
	if err != nil {
		return nil, err
	}
	if len(usersId) != len(authors) {
		return nil, global.ErrorUserNotExist
	}
	in.Authors = ""
	for _, uId := range usersId {
		userInfo, err := mysql.GetUserInfoByID(uId)
		if err != nil {
			return nil, err
		}
		in.Authors = in.Authors + "," + userInfo.FirstName + " " + userInfo.LastName
	}
	// 去掉第一个逗号
	in.Authors = in.Authors[1:]
	// 投稿信息
	paper := &tables.Paper{
		PaperType: in.PaperType,
		Title:     in.Title,
		Authors:   in.Authors,
		Abstract:  in.Abstract,
		//KeyWords:        string(keywords),
		//SubjectCategory: string(subjectCategory),
		KeyWords:        in.Keywords,
		SubjectCategory: in.SubjectCategory,
		//ManuscriptID:       in.ManuscriptID,
		//InformedConsent:    in.InformedConsent,
		//AnimalSubjects:     in.AnimalSubjects,
		CorAuthor:            in.CorAuthor,
		ManuscriptType:       in.ManuscriptType,
		UniqueContribution:   in.UniqueContribution,
		BlockAddress:         in.BlockAddress,
		PaperTransactionHash: in.PaperTransactionAddress,
		Hash:                 in.Hash,
		Filepath:             filePath,
		Status:               "UnReview", // 默认为未审核
		//Users:              users,      // gorm在中间表中自动插入数据
	}
	// 存入数据库
	return mysql.UpdatePaper2(userId, in, paper)
}

// DeletePaper 删除投稿
func DeletePaper(paperId uint) (err error) {
	// 先查询投稿信息
	sourcePaper, err := mysql.GetPaper(paperId)
	if err != nil {
		return err
	}
	// 如果投稿的状态不为待审核，则不允许删除
	if sourcePaper.Status != "UnReview" {
		return global.ErrPaperReviewed
	}
	return mysql.DeletePaper(paperId)
}

// GetAllAcceptPapers 查询所有已通过的投稿
func GetAllAcceptPapers() (out []*tables.Paper, err error) {
	return mysql.GetAllAcceptPapers()
}

// GetAllAcceptPapersByJournalAndTime 按期刊和时间查询已经审核通过的投稿列表
func GetAllAcceptPapersByJournalAndTime(journalId string, startTime, endTime time.Time) (out []*tables.Paper, err error) {
	journalID, _ := strconv.Atoi(journalId)
	return mysql.GetAllAcceptPapersByJournalAndTime(journalID, startTime, endTime)
}

// GetAllAcceptPapersByConferenceAndTime 按会议和时间查询已经审核通过的投稿列表
func GetAllAcceptPapersByConferenceAndTime(conferenceId string, startTime, endTime time.Time) (out []*tables.Paper, err error) {
	conferenceID, _ := strconv.Atoi(conferenceId)
	return mysql.GetAllAcceptPapersByConferenceAndTime(conferenceID, startTime, endTime)
}

// GetPaperVersions 获取投稿的所有版本
func GetPaperVersions(versionId uint) (out []*tables.Paper, err error) {
	return mysql.GetPaperVersions(versionId)
}

var honoraryCertificateUri string

// GetHonoraryCertificate 获取荣誉证书
func GetHonoraryCertificate(paperId uint) (honoraryCertificateInfo *response.HonoraryCertificateInfo, err error) {
	// 根据paperId查Paper的信息
	paper := new(tables.Paper)
	honoraryCertificateInfo = new(response.HonoraryCertificateInfo)
	if paper, err = mysql.GetPaper(paperId); err != nil {
		return
	}
	// 生成证书
	honoraryCertificateInfo.Url, err = createHonoraryCertificate(paper)
	// 把图片存入ipfs并且返回cid
	cid, err := saveToIPFS(honoraryCertificateUri)
	if err != nil {
		return
	}
	honoraryCertificateInfo.Cid = cid
	honoraryCertificateInfo.Uri = "http://" + global.MPS_CONFIG.IPFS.Host + ":" + global.MPS_CONFIG.IPFS.GatewayPort + global.MPS_CONFIG.IPFS.GatewayPath + cid
	// 创建json元数据
	//metadata := make(map[string]string)
	//metadata["name"] = paper.Title
	//metadata["description"] = paper.ManuscriptID
	//metadata["image"] = honoraryCertificateInfo.Uri
	metadata := new(struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
	})
	metadata.Name = paper.Title
	metadata.Description = paper.ManuscriptID
	metadata.Image = honoraryCertificateInfo.Uri
	// 转为json格式，保存到文件
	metadataPath := "./public/" + paper.ManuscriptID + ".json"

	{ // 可以不用逐行写入
		// map转json字符串
		//metadataByte, err := json.Marshal(metadata)
		//if err != nil {
		//	return
		//}
		// 写入文件
		// 先格式化json字符串
		//err = ioutil.WriteFile(metadataPath, metadataByte, 0644)
		//if err != nil {
		//	return
		//}
		// 打开文件以进行写入
		//jsonData, err := json.MarshalIndent(metadata, "", "    ")
		//err = ioutil.WriteFile(metadataPath, jsonData, 0644)
		//if err != nil {
		//	return
		//}
	}
	file, err := os.Create(metadataPath)
	if err != nil {
		return
	}
	defer file.Close()

	// 写入文件
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(metadata); err != nil {
		return
	}
	// 把json文件存入ipfs并且返回cid
	metadataCid, err := saveToIPFS(metadataPath)
	if err != nil {
		return
	}
	honoraryCertificateInfo.MetadataUri = "http://" + global.MPS_CONFIG.IPFS.Host + ":" + global.MPS_CONFIG.IPFS.GatewayPort + global.MPS_CONFIG.IPFS.GatewayPath + metadataCid
	// 更新paper表中的image_uri和json_uri
	if err = mysql.SetPaperInfo(paperId, honoraryCertificateInfo.Uri, honoraryCertificateInfo.Url, cid, honoraryCertificateInfo.MetadataUri); err != nil {
		return
	}
	// 返回证书路径
	return honoraryCertificateInfo, err
}

var (
	fontKai *truetype.Font // 字体
	fontTtf *truetype.Font // 字体
)

// 根据路径加载字体文件
// path 字体的路径
func loadFont(path string) (font *truetype.Font, err error) {
	var fontBytes []byte
	fontBytes, err = ioutil.ReadFile(path) // 读取字体文件
	if err != nil {
		err = fmt.Errorf("加载字体文件出错:%s", err.Error())
		return
	}
	font, err = freetype.ParseFont(fontBytes) // 解析字体文件
	if err != nil {
		err = fmt.Errorf("解析字体文件出错,%s", err.Error())
		return
	}
	return
}
func contentAuthors(content *freetype.Context, authors string) {
	content.SetFontSize(80) // 设置字体大小
	// 一个字母大小占40宽，图片宽2000,高1414
	// 计算authors所占的宽度，然后计算出居中的x坐标
	authorsWidth := len(authors) * 40
	authorsX := (2000 - authorsWidth) / 2
	content.DrawString(authors, freetype.Pt(authorsX, 520))
}
func contentData(content *freetype.Context, title, name string) {
	content.SetFontSize(30) // 设置字体大小
	data := "Have Successfully Published a Paper Titled" + title + " in the " + name
	dataX := 480
	dataY := 720
	for i := 0; i < len(data); i += 90 {
		if i == 0 {
			content.DrawString(data[i:i+56], freetype.Pt(dataX+110, dataY))
			dataY += 75
			i -= 10
			continue
		}
		content.DrawString(data[i:i+56], freetype.Pt(dataX+110, dataY))
		if i+90 > len(data) {
			content.DrawString(data[i:], freetype.Pt(dataX, dataY))
			break
		}
		content.DrawString(data[i:i+85], freetype.Pt(dataX, dataY))
		dataY += 75
	}
}

func contentIPFS(content *freetype.Context, IPFS string) {
	content.SetFontSize(30) // 设置字体大小
	content.DrawString("107.155.56.166:8080/ipfs/"+IPFS, freetype.Pt(430, 980))
}

func contentHash(content *freetype.Context, transactionAddress, blockAddress string) {
	content.SetFontSize(30) // 设置字体大小
	content.DrawString(transactionAddress, freetype.Pt(500, 930))
	content.DrawString(blockAddress, freetype.Pt(440, 865))
}

func contentDate(content *freetype.Context) {
	date := time.Now().Format("2006.01.02")
	content.SetFontSize(30) // 设置字体大小
	content.DrawString(date, freetype.Pt(440, 1050))
}
func createHonoraryCertificate(paper *tables.Paper) (honoraryCertificateUrl string, err error) {
	// 根据路径打开模板文件
	templateFile, err := os.Open("./image/newcertificate.png")
	if err != nil {
		return
	}
	defer templateFile.Close()
	// 解码
	templateFileImage, err := png.Decode(templateFile)
	if err != nil {
		return
	}
	// 新建一张和模板文件一样大小的画布
	newTemplateImage := image.NewRGBA(templateFileImage.Bounds())
	// 将模板图片画到新建的画布上
	draw.Draw(newTemplateImage, templateFileImage.Bounds(), templateFileImage, templateFileImage.Bounds().Min, draw.Over)

	// 加载字体文件  这里我们加载两种字体文件
	fontKai, err = loadFont("./times.ttf")
	if err != nil {
		return
	}
	// 向图片中写入文字
	// 在写入之前有一些准备工作
	content := freetype.NewContext()
	content.SetClip(newTemplateImage.Bounds())
	content.SetDst(newTemplateImage)
	content.SetSrc(image.Black) // 设置字体颜色
	content.SetDPI(72)          // 设置字体分辨率

	content.SetFontSize(80)  // 设置字体大小
	content.SetFont(fontKai) // 设置字体样式，就是我们上面加载的字体

	contentAuthors(content, paper.Authors) // 写入作者信息
	// 根据paperId查投的是哪个会议或者期刊
	conferenceOrJournalName, err := mysql.GetConferenceOrJournal(paper.ConferenceId, paper.JournalId)
	contentData(content, paper.Title, conferenceOrJournalName)
	contentHash(content, paper.PaperTransactionHash, paper.BlockAddress) // 写入hash信息
	contentIPFS(content, paper.Cid)
	contentDate(content) // 写入日期信息

	// 保存图片
	honoraryCertificateUri = "./image/" + paper.ManuscriptID + ".png"
	dstFile, err := os.Create(honoraryCertificateUri)
	if err != nil {
		return
	}
	defer dstFile.Close()
	png.Encode(dstFile, newTemplateImage)
	// 返回路径
	return global.MPS_CONFIG.ImagePath + paper.ManuscriptID + ".png", err
}

// PublishPaper 发布投稿
func PublishPaper(in *request.PublishPaper) (err error) {
	// 先查询投稿信息
	sourcePaper, err := mysql.GetPaper(in.PaperId)
	if err != nil {
		return
	}
	// 如果投稿的状态不为接受，则不允许发布
	if sourcePaper.Status != "Accept" {
		return global.ErrPaperNotAccepted
	}
	return mysql.PublishPaper(in)
}

// AddPaperViewer 设置投稿可查看者
func AddPaperViewer(in *request.AddPaperViewer) (err error) {
	return mysql.AddPaperViewer(in)
}

// CheckPaperViewer 查看用户是否有权限查看投稿
func CheckPaperViewer(paperId, viewerId uint) (isViewer bool, err error) {
	return mysql.CheckPaperViewer(paperId, viewerId)
}

// GetMyNFTs 获取我的NFT
func GetMyNFTs(userId uint) (out []*response.GetMyNFTs, err error) {
	return mysql.GetMyNFTs(userId)
}

// UpdatePrice 更新价格
func UpdatePrice(in *request.UpdatePrice) (err error) {
	return mysql.UpdatePrice(in)
}

// GetNFTInfoByTokenId 根据tokenId获取NFT信息
func GetNFTInfoByTokenId(tokenIds string) (out []*response.GetMyNFTs, err error) {
	// tokenIds = 8,9;将tokenIds转为int切片
	tokenIdsSlice := strings.Split(tokenIds, ",")
	tokenIdsInt := make([]int, len(tokenIdsSlice))
	for i, v := range tokenIdsSlice {
		if v == "0" {
			continue
		}
		tokenIdsInt[i], err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	return mysql.GetNFTInfoByTokenId(tokenIdsInt)
}

// UpdatePaperUserId 修改投稿对应的user_id
func UpdatePaperUserId(paperId, userId uint) (err error) {
	return mysql.UpdatePaperUserId(paperId, userId)
}

package logic

import (
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
)

// CreateJournal 创建期刊
func CreateJournal(in *request.CreateJournal) (out *tables.Journal, err error) {
	// 1. 根据主席名字获取主席ID
	// 主席名字可能有多个，所以把主席名字放在一个切片里
	var presidentsName []string
	for _, v := range in.Presidents {
		presidentsName = append(presidentsName, v.Name)
	}
	// 根据主席名字获取主席ID
	presidentsIDs, err := mysql.GetUsersIdByNames(presidentsName)
	if err != nil {
		return
	}
	// 2. 根据副主席名字获取副主席ID
	// 副主席名字可能有多个，所以把副主席名字放在一个切片里
	var vicePresidentsName []string
	for _, v := range in.VicePresidents {
		vicePresidentsName = append(vicePresidentsName, v.Name)
	}
	vicePresidentsIDs, err := mysql.GetUsersIdByNames(vicePresidentsName)
	if err != nil {
		return
	}
	// 3. 根据成员名字获取成员ID
	// 成员名字可能有多个，所以把成员名字放在一个切片里
	var membersName []string
	for _, v := range in.Members {
		membersName = append(membersName, v.Name)
	}
	membersIDs, err := mysql.GetUsersIdByNames(membersName)
	if err != nil {
		return
	}
	// 4. 创建期刊
	out, err = mysql.CreateJournal(in, presidentsIDs, vicePresidentsIDs, membersIDs)
	return
}

// UpdateJournal 更新期刊
func UpdateJournal(in *request.UpdateJournal) (out *tables.Journal, err error) {
	out = new(tables.Journal)
	// 1. 根据主席名字获取主席ID
	// 主席名字可能有多个，所以把主席名字放在一个切片里
	var presidentsName []string
	for _, v := range in.Presidents {
		presidentsName = append(presidentsName, v.Name)
	}
	// 根据主席名字获取主席ID
	presidentsID, err := mysql.GetUsersIdByNames(presidentsName)
	if err != nil {
		return
	}
	// 2. 根据副主席名字获取副主席ID
	// 副主席名字可能有多个，所以把副主席名字放在一个切片里
	var vicePresidentsName []string
	for _, v := range in.VicePresidents {
		vicePresidentsName = append(vicePresidentsName, v.Name)
	}
	// 根据副主席名字获取副主席ID
	vicePresidentsID, err := mysql.GetUsersIdByNames(vicePresidentsName)
	if err != nil {
		return
	}
	// 3. 根据成员名字获取成员ID
	// 成员名字可能有多个，所以把成员名字放在一个切片里
	var membersName []string
	for _, v := range in.Members {
		membersName = append(membersName, v.Name)
	}
	// 根据成员名字获取成员ID
	membersID, err := mysql.GetUsersIdByNames(membersName)
	if err != nil {
		return
	}
	// 4. 更新期刊
	out, err = mysql.UpdateJournal(in, presidentsID, vicePresidentsID, membersID)
	return
}

// GetJournalsByUser 查询自己所在的期刊列表
func GetJournalsByUser(userID uint) (out []*tables.Journal, err error) {
	// 1. 根据用户ID获取用户所在的期刊ID
	journalIDs, err := mysql.GetJournalIDsByUserID(userID)
	if err != nil {
		return
	}
	// 2. 根据期刊ID获取期刊信息
	out, err = mysql.GetJournalsByJournalIDs(journalIDs)
	return
}

// GetJournal 获取期刊详情
func GetJournal(id uint) (out *response.JournalDetail, err error) {
	// 1. 根据期刊ID获取期刊信息
	journalInfo, err := mysql.GetJournal(id)
	if err != nil {
		return
	}
	// 2. 根据期刊ID获取主席信息
	presidents, err := mysql.GetUserJournalInfoByUserID(id, "president")
	if err != nil {
		return nil, err
	}
	// 3. 根据副主席ID获取副主席信息
	vicePresident, err := mysql.GetUserJournalInfoByUserID(id, "vice_president")
	if err != nil {
		return nil, err
	}
	// 4. 根据成员ID获取成员信息
	members, err := mysql.GetUserJournalInfoByUserID(id, "member")
	if err != nil {
		return nil, err
	}
	// 5. 将期刊信息转换成JournalDetail结构体
	out = &response.JournalDetail{
		CreatorId:      journalInfo.CreatorId,
		Name:           journalInfo.Name,
		CreateAt:       journalInfo.CreatedAt,
		Description:    journalInfo.Description,
		Category:       journalInfo.Category,
		Presidents:     presidents,
		VicePresidents: vicePresident,
		Members:        members,
	}
	return
}

// GetAllJournals 获取期刊列表
func GetAllJournals() (journalList []*tables.Journal, err error) {
	return mysql.GetAllJournals()
}

// GetAllJournalsByCommittee 根据委员会查询期刊列表
func GetAllJournalsByCommittee(committeeID uint) (journalList []*tables.Journal, err error) {
	// 1. 根据委员会ID获取期刊
	return mysql.GetAllJournalsByCommittee(committeeID)

}

// DeleteJournal 删除期刊
func DeleteJournal(journalID uint) (err error) {
	// 1. 删除期刊
	return mysql.DeleteJournal(journalID)
}

// CreateJournalIssue 创建Issue
func CreateJournalIssue(in *request.CreateJournalIssue) error {
	// 1. 从投稿开始时间来填充Issue的年份字段
	year := in.SubmissionStartTime.Year()
	// 2. 创建Issue
	return mysql.CreateJournalIssue(year, in)

}

// GetAllJournalIssues  查询期刊Issue列表
func GetAllJournalIssues(journalID uint) (out []*tables.JournalIssue, err error) {
	// 1. 根据期刊ID获取期刊Issue列表
	out, err = mysql.GetAllJournalIssues(journalID)
	return
}

// UpdateJournalIssue 更新期刊Issue
func UpdateJournalIssue(in *request.UpdateJournalIssue) error {
	// 2. 更新期刊Issue
	return mysql.UpdateJournalIssue(in)
}

// DeleteJournalIssue 删除期刊Issue
func DeleteJournalIssue(issueID uint) error {
	// 1. 查看该期刊Issue是否有文章
	articleCount, err := mysql.GetArticleCountByIssueID(issueID)
	if err != nil {
		return err
	}
	if articleCount > 0 {
		return global.ErrIssueHasPaper
	}
	// 2. 删除期刊Issue
	return mysql.DeleteJournalIssue(issueID)
}

// GetLevelInJournal 获取用户在期刊的level
func GetLevelInJournal(journalID, userID uint) (level string, err error) {
	return mysql.GetLevelInJournal(journalID, userID)
}

package mysql

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model/request"
	"server/model/tables"
)

// CreateJournal 创建期刊
func CreateJournal(in *request.CreateJournal, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Journal, err error) {
	journal := &tables.Journal{
		CreatorId:   in.CreatorId,
		CommitteeId: in.CommitteeId,
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 3. 插入journal表
	if err = tx.Create(journal).Error; err != nil {
		tx.Rollback()
		return
	}
	// 2. 将期刊人员信息信息转换成UserJournal结构体
	var presidents, vicePresidents, members []*tables.UserJournal
	if presidents, err = userJournal(in.Presidents, presidentsID, journal.ID); err != nil {
		return
	}
	if vicePresidents, err = userJournal(in.VicePresidents, vicePresidentsID, journal.ID); err != nil {
		return
	}
	if members, err = userJournal(in.Members, membersID, journal.ID); err != nil {
		return
	}
	// 3. 插入user_journal表
	// 有数据才插入
	if len(presidents) > 0 {
		if err = tx.Create(presidents).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if len(vicePresidents) > 0 {
		if err = tx.Create(vicePresidents).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if len(members) > 0 {
		if err = tx.Create(members).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	// 设置权限为成员
	// 从users表查出人员的权限 // 一共有4个权限，分别是101超级管理员，102可创建委员会、期刊、会议者，103成员，104普通用户
	userIds := make([]uint, 0, len(presidentsID)+len(vicePresidentsID)+len(membersID))
	for _, v := range presidentsID {
		userIds = append(userIds, v)
	}
	for _, v := range vicePresidentsID {
		userIds = append(userIds, v)
	}
	for _, v := range membersID {
		userIds = append(userIds, v)
	}
	var users []*tables.User
	if err = tx.Where("id in ?", userIds).Find(&users).Error; err != nil {
		tx.Rollback()
		return
	}
	for i := range users {
		if users[i].AuthorityId > 103 {
			users[i].AuthorityId = 103
		}
	}
	if err = tx.Save(&users).Error; err != nil {
		tx.Rollback()
		return
	}
	// 4. 提交事务
	if err = tx.Commit().Error; err != nil {
		return
	}
	return journal, nil
}

// userJournal 将主席/副主席/成员信息转换成userJournal结构体
func userJournal(members []*request.Info, IDs []uint, journalId uint) (memberInfos []*tables.UserJournal, err error) {
	memberInfos = make([]*tables.UserJournal, 0, len(members))
	//var startTime, endTime time.Time
	for i, v := range members {
		//startTime, err = utils.ParseTime(v.StartTime)
		//startTime, err = time.ParseInLocation("2006-01-02 00:00:00", v.StartTime, time.Local)
		//if err != nil {
		//	return
		//}
		//endTime, err = utils.ParseTime(v.EndTime)
		//endTime, err = time.ParseInLocation("2006-01-02 00:00:00", v.EndTime, time.Local)
		//if err != nil {
		//	return
		//}
		memberInfos = append(memberInfos, &tables.UserJournal{
			UserId:    IDs[i],
			JournalId: journalId,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
			Position:  v.Position,
			Level:     v.Level,
		})
	}
	return
}

// UpdateJournal 更新期刊
func UpdateJournal(in *request.UpdateJournal, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Journal, err error) {
	journal := &tables.Journal{
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 2. 更新journal表
	if err = tx.Model(&tables.Journal{}).Where("id = ?", in.ID).Updates(journal).First(journal).Error; err != nil {
		tx.Rollback()
		return
	}
	// 3. 删除user_journal表中的期刊人员的相关信息
	if err = tx.Where("journal_id = ?", in.ID).Delete(&tables.UserJournal{}).Error; err != nil {
		tx.Rollback()
		return
	}
	// 4. 将期刊人员信息信息转换成UserJournal结构体
	var presidents, vicePresidents, members []*tables.UserJournal
	if presidents, err = userJournal(in.Presidents, presidentsID, journal.ID); err != nil {
		tx.Rollback()
		return
	}
	if vicePresidents, err = userJournal(in.VicePresidents, vicePresidentsID, journal.ID); err != nil {
		tx.Rollback()
		return
	}
	if members, err = userJournal(in.Members, membersID, journal.ID); err != nil {
		tx.Rollback()
		return
	}

	// 5. 插入user_journal表
	// 有数据才插入
	if len(presidents) > 0 {
		if err = tx.Create(presidents).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if len(vicePresidents) > 0 {
		if err = tx.Create(vicePresidents).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if len(members) > 0 {
		if err = tx.Create(members).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	// 设置权限为成员
	// 从users表查出人员的权限 // 一共有4个权限，分别是101超级管理员，102可创建委员会、期刊、会议者，103成员，104普通用户
	userIds := make([]uint, 0, len(presidentsID)+len(vicePresidentsID)+len(membersID))
	for _, v := range presidentsID {
		userIds = append(userIds, v)
	}
	for _, v := range vicePresidentsID {
		userIds = append(userIds, v)
	}
	for _, v := range membersID {
		userIds = append(userIds, v)
	}
	var users []*tables.User
	if err = tx.Where("id in ?", userIds).Find(&users).Error; err != nil {
		tx.Rollback()
		return
	}
	for i := range users {
		if users[i].AuthorityId > 103 {
			users[i].AuthorityId = 103
		}
	}
	if err = tx.Save(&users).Error; err != nil {
		tx.Rollback()
		return
	}
	// 4. 提交事务
	if err = tx.Commit().Error; err != nil {
		return
	}
	return journal, nil
}

// GetJournalIDsByUserID 根据userID获取期刊ID
func GetJournalIDsByUserID(id uint) (journalIDs []uint, err error) {
	userJournals := make([]*tables.UserJournal, 0)
	if err = global.MPS_DB.Where("user_id = ?", id).Find(&userJournals).Error; err != nil {
		return
	}
	for _, userJournal := range userJournals {
		journalIDs = append(journalIDs, userJournal.JournalId)
	}
	return
}

// GetJournalsByJournalIDs 根据journalIDs获取期刊
func GetJournalsByJournalIDs(journalIDs []uint) (journals []*tables.Journal, err error) {
	if err = global.MPS_DB.Where("id IN ?", journalIDs).Find(&journals).Error; err != nil {
		return
	}
	return
}

// GetJournal 根据期刊ID获取期刊详情
func GetJournal(journalID uint) (journal *tables.Journal, err error) {
	journal = new(tables.Journal)
	err = global.MPS_DB.Where("id = ?", journalID).First(journal).Error
	return
}

// GetUserJournalInfoByUserID 根据userID获取成员在期刊的信息
func GetUserJournalInfoByUserID(id uint, position string) (members []*request.Info, err error) {
	memberInfos := make([]*tables.UserJournal, 0)
	if err = global.MPS_DB.Where("journal_id = ? AND level = ?", id, position).Find(&memberInfos).Error; err != nil {
		return
	}
	for _, memberInfo := range memberInfos {
		userInfo, err := GetUserInfoByID(memberInfo.UserId)
		if err != nil {
			return nil, err
		}
		member := &request.Info{
			Name:      userInfo.Username,
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			HeaderImg: userInfo.HeaderImg,
			Position:  memberInfo.Position,
			StartTime: memberInfo.StartTime,
			EndTime:   memberInfo.EndTime,
			Level:     memberInfo.Level,
		}
		members = append(members, member)
	}
	return
}

// GetAllJournals 期刊列表
func GetAllJournals() (journalList []*tables.Journal, err error) {
	journalList = make([]*tables.Journal, 0)
	// 1. 查询期刊
	err = global.MPS_DB.Model(&tables.Journal{}).Find(&journalList).Error
	return
}

// GetAllJournalsByCommittee 根据委员会ID获取期刊列表
func GetAllJournalsByCommittee(committeeID uint) (journalList []*tables.Journal, err error) {
	journalList = make([]*tables.Journal, 0)
	err = global.MPS_DB.Where("committee_id = ?", committeeID).Find(&journalList).Error
	return
}

// DeleteJournal 删除期刊
func DeleteJournal(id uint) (err error) {
	tx := global.MPS_DB.Begin()
	if err = tx.Where("id = ?", id).Delete(&tables.Journal{}).Error; err != nil {
		return
	}
	if err = tx.Where("journal_id = ?", id).Delete(&tables.UserJournal{}).Error; err != nil {
		return
	}
	return tx.Commit().Error
}

// GetUserJournalByIdAndUserId 根据期刊ID和用户ID获取用户期刊信息
func GetUserJournalByIdAndUserId(Id, userId uint) (userJournal *tables.UserJournal, err error) {
	userJournal = new(tables.UserJournal)
	err = global.MPS_DB.Where("journal_id = ? AND user_id = ?", Id, userId).First(userJournal).Error
	//if err == gorm.ErrRecordNotFound {
	//	return nil,nil
	//}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// CreateJournalIssue 创建Issue
func CreateJournalIssue(year int, in *request.CreateJournalIssue) error {
	issue := &tables.JournalIssue{
		JournalId:           in.JournalId,
		Name:                in.Name,
		SubmissionStartTime: in.SubmissionStartTime,
		SubmissionEndTime:   in.SubmissionEndTime,
		Description:         in.Description,
		Year:                year,
		Volume:              in.Volume,
	}
	return global.MPS_DB.Create(issue).Error
}

// GetAllJournalIssues  查询期刊Issue列表
func GetAllJournalIssues(journalId uint) (issues []*tables.JournalIssue, err error) {
	issues = make([]*tables.JournalIssue, 0)
	err = global.MPS_DB.Where("journal_id = ?", journalId).Find(&issues).Error
	return
}

// UpdateJournalIssue 更新期刊Issue
func UpdateJournalIssue(in *request.UpdateJournalIssue) error {
	issue := &tables.JournalIssue{
		Name:                in.Name,
		SubmissionStartTime: in.SubmissionStartTime,
		SubmissionEndTime:   in.SubmissionEndTime,
		Year:                in.SubmissionStartTime.Year(),
		Description:         in.Description,
		Volume:              in.Volume,
	}
	return global.MPS_DB.Model(&tables.JournalIssue{}).Where("id = ?", in.ID).Updates(issue).Error
}

// GetArticleCountByIssueID 根据IssueID获取文章数量
func GetArticleCountByIssueID(issueID uint) (count int64, err error) {
	// 查出Issue的开始时间和结束时间
	issue := new(tables.JournalIssue)
	if err = global.MPS_DB.Where("id = ?", issueID).First(issue).Error; err != nil {
		return
	}
	// 查出文章数量
	err = global.MPS_DB.Model(&tables.Paper{}).Where("created_at BETWEEN ? AND ?", issue.SubmissionStartTime, issue.SubmissionEndTime).Count(&count).Error
	return
}

// DeleteJournalIssue 删除期刊Issue
func DeleteJournalIssue(id uint) error {
	return global.MPS_DB.Where("id = ?", id).Delete(&tables.JournalIssue{}).Error
}

// GetLevelInJournal 获取用户在期刊的level
func GetLevelInJournal(journalID, userID uint) (level string, err error) {
	userJournalInfo := new(tables.UserJournal)
	if err = global.MPS_DB.Where("journal_id = ? and user_id = ?", journalID, userID).First(userJournalInfo).Error; err != nil {
		return
	}
	return userJournalInfo.Level, nil
}

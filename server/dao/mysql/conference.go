package mysql

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model/request"
	"server/model/tables"
)

// CreateConference 创建会议
func CreateConference(in *request.CreateConference, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Conference, err error) {
	conference := &tables.Conference{
		CommitteeId: in.CommitteeId,
		CreatorId:   in.CreatorId,
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
		Location:    in.Location,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 3. 插入conference表
	if err = tx.Create(conference).Error; err != nil {
		tx.Rollback()
		return
	}
	// 2. 将会议人员信息信息转换成UserConference结构体
	var presidents, vicePresidents, members []*tables.UserConference
	if presidents, err = userConference(in.Presidents, presidentsID, conference.ID); err != nil {
		return
	}
	if vicePresidents, err = userConference(in.VicePresidents, vicePresidentsID, conference.ID); err != nil {
		return
	}
	if members, err = userConference(in.Members, membersID, conference.ID); err != nil {
		return
	}
	// 3. 插入user_conference表
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
	// 从users表查出人员的权限 // 一共有4个权限，分别是101超级管理员，102可创建委员会、会议、会议者，103成员，104普通用户
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
	return conference, nil
}

// userConference 将主席/副主席/成员信息转换成userConference结构体
func userConference(members []*request.Info, IDs []uint, conferenceId uint) (memberInfos []*tables.UserConference, err error) {
	memberInfos = make([]*tables.UserConference, 0, len(members))
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
		memberInfos = append(memberInfos, &tables.UserConference{
			UserId:       IDs[i],
			ConferenceId: conferenceId,
			StartTime:    v.StartTime,
			EndTime:      v.EndTime,
			Position:     v.Position,
			Level:        v.Level,
		})
	}
	return
}

// UpdateConference 更新会议
func UpdateConference(in *request.UpdateConference, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Conference, err error) {
	conference := &tables.Conference{
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
		Location:    in.Location,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 2. 更新conference表
	if err = tx.Model(&tables.Conference{}).Where("id = ?", in.ID).Updates(conference).First(conference).Error; err != nil {
		tx.Rollback()
		return
	}
	// 3. 删除user_conference表中的会议人员的相关信息
	if err = tx.Where("conference_id = ?", in.ID).Delete(&tables.UserConference{}).Error; err != nil {
		tx.Rollback()
		return
	}
	// 4. 将会议人员信息信息转换成UserConference结构体
	var presidents, vicePresidents, members []*tables.UserConference
	if presidents, err = userConference(in.Presidents, presidentsID, conference.ID); err != nil {
		tx.Rollback()
		return
	}
	if vicePresidents, err = userConference(in.VicePresidents, vicePresidentsID, conference.ID); err != nil {
		tx.Rollback()
		return
	}
	if members, err = userConference(in.Members, membersID, conference.ID); err != nil {
		tx.Rollback()
		return
	}

	// 5. 插入user_conference表
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
	// 从users表查出人员的权限 // 一共有4个权限，分别是101超级管理员，102可创建委员会、会议、会议者，103成员，104普通用户
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
	return conference, nil
}

// GetConference 根据会议ID获取会议详情
func GetConference(conferenceID uint) (conference *tables.Conference, err error) {
	conference = new(tables.Conference)
	err = global.MPS_DB.Where("id = ?", conferenceID).First(conference).Error
	return
}

// GetUserConferenceInfo 获取用户会议信息
func GetUserConferenceInfo(id uint, position string) (members []*request.Info, err error) {
	memberInfos := make([]*tables.UserConference, 0)
	if err = global.MPS_DB.Where("conference_id = ? AND level = ?", id, position).Find(&memberInfos).Error; err != nil {
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

// GetAllConferences 会议列表
func GetAllConferences() (conferenceList []*tables.Conference, err error) {
	conferenceList = make([]*tables.Conference, 0)
	// 1. 查询会议
	err = global.MPS_DB.Model(&tables.Conference{}).Find(&conferenceList).Error
	return
}

// GetConferencesByCommitteeID 根据委员会ID获取会议列表
func GetConferencesByCommitteeID(committeeID uint) (conferenceList []*tables.Conference, err error) {
	conferenceList = make([]*tables.Conference, 0)
	err = global.MPS_DB.Where("committee_id = ?", committeeID).Find(&conferenceList).Error
	return
}

// DeleteConference 删除会议
func DeleteConference(id uint) (err error) {
	tx := global.MPS_DB.Begin()
	if err = tx.Where("id = ?", id).Delete(&tables.Conference{}).Error; err != nil {
		return
	}
	if err = tx.Where("conference_id = ?", id).Delete(&tables.UserConference{}).Error; err != nil {
		return
	}
	return tx.Commit().Error
}

// GetUserConferenceByIdAndUserId 根据会议ID和用户ID获取用户会议信息
func GetUserConferenceByIdAndUserId(Id, userId uint) (userConference *tables.UserConference, err error) {
	userConference = new(tables.UserConference)
	err = global.MPS_DB.Where("conference_id = ? AND user_id = ?", Id, userId).First(userConference).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// GetConferenceIDsByUserID 根据userID获取会议ID
func GetConferenceIDsByUserID(id uint) (conferenceIDs []uint, err error) {
	userConferences := make([]*tables.UserConference, 0)
	if err = global.MPS_DB.Where("user_id = ?", id).Find(&userConferences).Error; err != nil {
		return
	}
	for _, userConference := range userConferences {
		conferenceIDs = append(conferenceIDs, userConference.ConferenceId)
	}
	return
}

// GetConferencesByConferenceIDs 根据conferenceIDs获取会议
func GetConferencesByConferenceIDs(conferenceIDs []uint) (conferences []*tables.Conference, err error) {
	if err = global.MPS_DB.Where("id IN ?", conferenceIDs).Find(&conferences).Error; err != nil {
		return
	}
	return
}

// CreateConferenceIssue 创建Issue
func CreateConferenceIssue(year int, in *request.CreateConferenceIssue) error {
	issue := &tables.ConferenceIssue{
		ConferenceId:        in.ConferenceId,
		Name:                in.Name,
		SubmissionStartTime: in.SubmissionStartTime,
		SubmissionEndTime:   in.SubmissionEndTime,
		Description:         in.Description,
		Year:                year,
		Volume:              in.Volume,
	}
	return global.MPS_DB.Create(issue).Error
}

// GetAllConferenceIssues  查询会议Issue列表
func GetAllConferenceIssues(conferenceId uint) (issues []*tables.ConferenceIssue, err error) {
	issues = make([]*tables.ConferenceIssue, 0)
	err = global.MPS_DB.Where("conference_id = ?", conferenceId).Find(&issues).Error
	return
}

// UpdateConferenceIssue 更新会议Issue
func UpdateConferenceIssue(in *request.UpdateConferenceIssue) error {
	issue := &tables.ConferenceIssue{
		Name:                in.Name,
		SubmissionStartTime: in.SubmissionStartTime,
		SubmissionEndTime:   in.SubmissionEndTime,
		Year:                in.SubmissionStartTime.Year(),
		Description:         in.Description,
		Volume:              in.Volume,
	}
	return global.MPS_DB.Model(&tables.ConferenceIssue{}).Where("id = ?", in.ID).Updates(issue).Error
}

// GetArticleCountByIssueIDFromConference 根据IssueID获取文章数量
func GetArticleCountByIssueIDFromConference(issueID uint) (count int64, err error) {
	// 查出Issue的开始时间和结束时间
	issue := new(tables.ConferenceIssue)
	if err = global.MPS_DB.Where("id = ?", issueID).First(issue).Error; err != nil {
		return
	}
	// 查出文章数量
	err = global.MPS_DB.Model(&tables.Paper{}).Where("created_at BETWEEN ? AND ?", issue.SubmissionStartTime, issue.SubmissionEndTime).Count(&count).Error
	return
}

// DeleteConferenceIssue 删除会议Issue
func DeleteConferenceIssue(id uint) error {
	return global.MPS_DB.Where("id = ?", id).Delete(&tables.ConferenceIssue{}).Error
}

// GetLevelInConference 获取用户在会议的level
func GetLevelInConference(conferenceID, userID uint) (level string, err error) {
	userConferenceInfo := new(tables.UserConference)
	if err = global.MPS_DB.Where("conference_id = ? and user_id = ?", conferenceID, userID).First(userConferenceInfo).Error; err != nil {
		return
	}
	return userConferenceInfo.Level, nil
}

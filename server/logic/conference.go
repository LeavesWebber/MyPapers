package logic

import (
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
)

// CreateConference 创建会议
func CreateConference(in *request.CreateConference) (out *tables.Conference, err error) {
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
	// 4. 创建会议
	out, err = mysql.CreateConference(in, presidentsIDs, vicePresidentsIDs, membersIDs)
	return
}

// UpdateConference 更新会议
func UpdateConference(in *request.UpdateConference) (out *tables.Conference, err error) {
	out = new(tables.Conference)
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
	// 4. 更新会议
	out, err = mysql.UpdateConference(in, presidentsID, vicePresidentsID, membersID)
	return
}

// GetConference 获取会议详情
func GetConference(id uint) (out *response.ConferenceDetail, err error) {
	// 1. 根据会议ID获取会议信息
	conferenceInfo, err := mysql.GetConference(id)
	if err != nil {
		return
	}
	// 2. 根据会议ID获取主席信息
	presidents, err := mysql.GetUserConferenceInfo(id, "president")
	if err != nil {
		return nil, err
	}
	// 3. 根据副主席ID获取副主席信息
	vicePresident, err := mysql.GetUserConferenceInfo(id, "vice_president")
	if err != nil {
		return nil, err
	}
	// 4. 根据成员ID获取成员信息
	members, err := mysql.GetUserConferenceInfo(id, "member")
	if err != nil {
		return nil, err
	}
	// 5. 将会议信息转换成ConferenceDetail结构体
	out = &response.ConferenceDetail{
		CreatorId:      conferenceInfo.CreatorId,
		Name:           conferenceInfo.Name,
		Description:    conferenceInfo.Description,
		Category:       conferenceInfo.Category,
		Location:       conferenceInfo.Location,
		StartTime:      conferenceInfo.StartTime,
		EndTime:        conferenceInfo.EndTime,
		Presidents:     presidents,
		VicePresidents: vicePresident,
		Members:        members,
	}
	return
}

// GetAllConferences 获取会议列表
func GetAllConferences() (conferenceList []*tables.Conference, err error) {
	return mysql.GetAllConferences()
}

// GetAllConferencesByCommittee 根据委员会查询会议列表
func GetAllConferencesByCommittee(committeeID uint) (conferenceList []*tables.Conference, err error) {
	// 1. 根据委员会ID获取会议
	return mysql.GetConferencesByCommitteeID(committeeID)
}

// DeleteConference 删除会议
func DeleteConference(conferenceID uint) (err error) {
	// 1. 删除会议
	return mysql.DeleteConference(conferenceID)
}

// GetConferencesByUser 查询自己所在的会议列表
func GetConferencesByUser(userID uint) (out []*tables.Conference, err error) {
	// 1. 根据用户ID获取用户所在的会议ID
	conferenceIDs, err := mysql.GetConferenceIDsByUserID(userID)
	if err != nil {
		return
	}
	// 2. 根据会议ID获取会议信息
	out, err = mysql.GetConferencesByConferenceIDs(conferenceIDs)
	return
}

// CreateConferenceIssue 创建Issue
func CreateConferenceIssue(in *request.CreateConferenceIssue) error {
	// 1. 从投稿开始时间来填充Issue的年份字段
	year := in.SubmissionStartTime.Year()
	// 2. 创建Issue
	return mysql.CreateConferenceIssue(year, in)

}

// GetAllConferenceIssues  查询会议Issue列表
func GetAllConferenceIssues(conferenceID uint) (out []*tables.ConferenceIssue, err error) {
	// 1. 根据会议ID获取会议Issue列表
	out, err = mysql.GetAllConferenceIssues(conferenceID)
	return
}

// UpdateConferenceIssue 更新会议Issue
func UpdateConferenceIssue(in *request.UpdateConferenceIssue) error {
	// 2. 更新会议Issue
	return mysql.UpdateConferenceIssue(in)
}

// DeleteConferenceIssue 删除会议Issue
func DeleteConferenceIssue(issueID uint) error {
	// 1. 查看该会议Issue是否有文章
	articleCount, err := mysql.GetArticleCountByIssueID(issueID)
	if err != nil {
		return err
	}
	if articleCount > 0 {
		return global.ErrIssueHasPaper
	}
	// 2. 删除会议Issue
	return mysql.DeleteConferenceIssue(issueID)
}

// GetLevelInConference 获取用户在会议的level
func GetLevelInConference(conferenceID, userID uint) (level string, err error) {
	return mysql.GetLevelInConference(conferenceID, userID)
}

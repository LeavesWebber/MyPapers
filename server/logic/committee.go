package logic

import (
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
)

// CreateCommittee 创建委员会
func CreateCommittee(in *request.CreateCommittee) (out *tables.Committee, err error) {
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
	// 如果主席名字有误，名字和ID的长度就不一样
	if len(presidentsName) != len(presidentsIDs) {
		return nil, global.ErrorUserNotExist
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
	if len(vicePresidentsName) != len(vicePresidentsIDs) {
		return nil, global.ErrorUserNotExist
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
	if len(membersName) != len(membersIDs) {
		return nil, global.ErrorUserNotExist
	}
	// 如果id都没有获取到，说明名字有误
	if len(presidentsIDs) == 0 && len(vicePresidentsIDs) == 0 && len(membersIDs) == 0 {
		return nil, global.ErrorUserNotExist
	}
	// 4. 创建委员会
	out, err = mysql.CreateCommittee(in, presidentsIDs, vicePresidentsIDs, membersIDs)
	return
}

// UpdateCommittee 更新委员会
func UpdateCommittee(in *request.UpdateCommittee) (out *tables.Committee, err error) {
	out = new(tables.Committee)
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
	// 4. 更新委员会
	out, err = mysql.UpdateCommittee(in, presidentsID, vicePresidentsID, membersID)
	return
}

// GetCommitteesByUser 查询自己所在的委员会列表
func GetCommitteesByUser(userID uint) (out []*tables.Committee, err error) {
	// 1. 根据用户ID获取用户所在的期刊ID
	committeeIDs, err := mysql.GetCommitteeIDsByUserID(userID)
	if err != nil {
		return
	}
	// 2. 根据期刊ID获取期刊信息
	out, err = mysql.GetCommitteesByCommitteeIDs(committeeIDs)
	return
}

// GetCommittee 获取委员会详情
func GetCommittee(id uint) (out *response.CommitteeDetail, err error) {
	// 1. 根据委员会ID获取委员会信息
	committeeInfo, err := mysql.GetCommittee(id)
	if err != nil {
		return
	}
	// 2. 根据委员会ID获取主席信息
	presidents, err := mysql.GetUserCommitteeInfoByUserID(id, "president")
	if err != nil {
		return nil, err
	}
	// 3. 根据副主席ID获取副主席信息
	vicePresident, err := mysql.GetUserCommitteeInfoByUserID(id, "vice_president")
	if err != nil {
		return nil, err
	}
	// 4. 根据成员ID获取成员信息
	members, err := mysql.GetUserCommitteeInfoByUserID(id, "member")
	if err != nil {
		return nil, err
	}
	// 5. 将委员会信息转换成CommitteeDetail结构体
	out = &response.CommitteeDetail{
		CreatorId:      committeeInfo.CreatorId,
		Name:           committeeInfo.Name,
		Description:    committeeInfo.Description,
		Presidents:     presidents,
		VicePresidents: vicePresident,
		Members:        members,
	}
	return
}

// GetAllCommittees 获取委员会列表
func GetAllCommittees() (committeeList []*tables.Committee, err error) {
	return mysql.GetAllCommittees()
}

// DeleteCommittee 删除委员会
func DeleteCommittee(committeeID uint) (err error) {
	// 1. 删除委员会
	return mysql.DeleteCommittee(committeeID)
}

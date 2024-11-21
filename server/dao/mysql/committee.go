package mysql

import (
	"server/global"
	"server/model/request"
	"server/model/tables"
)

// CreateCommittee 创建委员会
func CreateCommittee(in *request.CreateCommittee, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Committee, err error) {
	committee := &tables.Committee{
		CreatorId:   in.CreatorId,
		Name:        in.Name,
		Description: in.Description,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 2. 插入committee表
	if err = tx.Create(committee).Error; err != nil {
		tx.Rollback()
		return
	}
	// 3. 将委员会人员信息信息转换成UserCommittee结构体
	var presidents, vicePresidents, members []*tables.UserCommittee
	if presidents, err = userCommittee(in.Presidents, presidentsID, committee.ID); err != nil {
		tx.Rollback()
		return
	}
	if vicePresidents, err = userCommittee(in.VicePresidents, vicePresidentsID, committee.ID); err != nil {
		tx.Rollback()
		return
	}
	if members, err = userCommittee(in.Members, membersID, committee.ID); err != nil {
		tx.Rollback()
		return
	}

	// 4. 插入user_committee表
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
	return committee, nil
}

// userCommittee 将主席/副主席/成员信息转换成userCommittee结构体
func userCommittee(members []*request.Info, IDs []uint, committeeId uint) (memberInfos []*tables.UserCommittee, err error) {
	memberInfos = make([]*tables.UserCommittee, 0, len(members))
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
		memberInfos = append(memberInfos, &tables.UserCommittee{
			UserId:      IDs[i],
			CommitteeId: committeeId,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			Position:    v.Position,
			Level:       v.Level,
		})
	}
	return
}

// UpdateCommittee 更新委员会
func UpdateCommittee(in *request.UpdateCommittee, presidentsID, vicePresidentsID, membersID []uint) (out *tables.Committee, err error) {
	committee := &tables.Committee{
		Name:        in.Name,
		Description: in.Description,
	}
	// 1. 开启事务
	tx := global.MPS_DB.Begin()
	// 2. 更新committee表
	if err = tx.Model(&tables.Committee{}).Where("id = ?", in.ID).Updates(committee).First(committee).Error; err != nil {
		tx.Rollback()
		return
	}
	// 3. 删除user_committee表中的委员会人员的相关信息
	if err = tx.Where("committee_id = ?", in.ID).Delete(&tables.UserCommittee{}).Error; err != nil {
		tx.Rollback()
		return
	}
	// 4. 将委员会人员信息信息转换成UserCommittee结构体
	var presidents, vicePresidents, members []*tables.UserCommittee
	if presidents, err = userCommittee(in.Presidents, presidentsID, committee.ID); err != nil {
		tx.Rollback()
		return
	}
	if vicePresidents, err = userCommittee(in.VicePresidents, vicePresidentsID, committee.ID); err != nil {
		tx.Rollback()
		return
	}
	if members, err = userCommittee(in.Members, membersID, committee.ID); err != nil {
		tx.Rollback()
		return
	}

	// 5. 插入user_committee表
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
	return committee, nil
}

// GetCommitteeIDsByUserID 根据userID获取委员会ID
func GetCommitteeIDsByUserID(id uint) (committeeIDs []uint, err error) {
	userCommittees := make([]*tables.UserCommittee, 0)
	if err = global.MPS_DB.Where("user_id = ?", id).Find(&userCommittees).Error; err != nil {
		return
	}
	for _, userCommittee := range userCommittees {
		committeeIDs = append(committeeIDs, userCommittee.CommitteeId)
	}
	return
}

// GetCommitteesByCommitteeIDs 根据committeeIDs委员会
func GetCommitteesByCommitteeIDs(committeeIDs []uint) (committees []*tables.Committee, err error) {
	if err = global.MPS_DB.Where("id IN ?", committeeIDs).Find(&committees).Error; err != nil {
		return
	}
	return
}

// GetCommittee 根据委员会ID获取委员会详情
func GetCommittee(committeeID uint) (committee *tables.Committee, err error) {
	committee = new(tables.Committee)
	err = global.MPS_DB.Where("id = ?", committeeID).First(committee).Error
	return
}

// GetUserCommitteeInfoByUserID 根据userID获取成员在委员会的信息
func GetUserCommitteeInfoByUserID(id uint, level string) (members []*request.Info, err error) {
	memberInfos := make([]*tables.UserCommittee, 0)
	if err = global.MPS_DB.Where("committee_id = ? AND level = ?", id, level).Find(&memberInfos).Error; err != nil {
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

// GetAllCommittees 委员会列表
func GetAllCommittees() (committeeList []*tables.Committee, err error) {
	committeeList = make([]*tables.Committee, 0)
	// 1. 查询委员会
	err = global.MPS_DB.Model(&tables.Committee{}).Find(&committeeList).Error
	return
}

// DeleteCommittee 删除委员会
func DeleteCommittee(id uint) (err error) {
	// 删除委员会的同时,删除委员会人员信息
	tx := global.MPS_DB.Begin()
	if err = tx.Where("id = ?", id).Delete(&tables.Committee{}).Error; err != nil {
		return
	}
	if err = tx.Where("committee_id = ?", id).Delete(&tables.UserCommittee{}).Error; err != nil {
		return
	}
	return tx.Commit().Error
}

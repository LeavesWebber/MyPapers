package logic

import (
	"server/dao/mysql"
	"server/model/request"
	"server/model/tables"
	"strconv"
)

// CreateAuthority 创建用户
func CreateAuthority(authority *request.CreateAuthority) (auth tables.Authority, err error) {
	// 存入数据库
	if auth, err = mysql.CreateAuthority(authority); err != nil {
		return
	} else {
		if err = AddMenuAuthority(request.DefaultMenu(), authority.AuthorityId); err != nil {
			return
		}
		if err = UpdateCasbin(authority.AuthorityId, request.DefaultCasbin()); err != nil {
			return
		}
	}
	return
}

// UpdateAuthority 更新角色信息
func UpdateAuthority(updateInfo *request.CreateAuthority) (authority tables.Authority, err error) {
	return mysql.UpdateAuthority(updateInfo)
}

// GetAuthorityList 获取角色列表
func GetAuthorityList(info *request.GetAuthorityList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	return mysql.GetAuthorityList(limit, offset)
}

// DeleteAuthority 删除角色
func DeleteAuthority(aid uint) (err error) {
	if err = mysql.DeleteAuthority(aid); err != nil {
		return
	} else {
		authorityId := strconv.Itoa(int(aid))
		ClearCasbin(0, authorityId)
	}
	return
}

// ChangeAuthority 切换角色
func ChangeAuthority(userId uint, in *request.ChangeAuthority) (err error) {
	return mysql.ChangeAuthority(userId, in)
}

package mysql

import (
	"errors"
	"server/global"
	"server/model/request"
	"server/model/tables"

	"gorm.io/gorm"
)

// CreateAuthority 创建角色
func CreateAuthority(authority *request.CreateAuthority) (auth tables.Authority, err error) {
	auth = tables.Authority{
		AuthorityId:   authority.AuthorityId,
		AuthorityName: authority.AuthorityName,
		ParentId:      authority.ParentId,
	}
	if !errors.Is(global.MPS_DB.Where("authority_id = ?", authority.AuthorityId).First(&auth).Error, gorm.ErrRecordNotFound) {
		return auth, global.ErrRoleExistence
	}
	err = global.MPS_DB.Create(&auth).Error
	return auth, err
}

// UpdateAuthority 更新角色信息
func UpdateAuthority(updateInfo *request.CreateAuthority) (authority tables.Authority, err error) {
	authority.AuthorityId = updateInfo.AuthorityId
	authority.AuthorityName = updateInfo.AuthorityName
	authority.ParentId = updateInfo.ParentId
	err = global.MPS_DB.Where("authority_id = ?", updateInfo.AuthorityId).First(&tables.Authority{}).Updates(&authority).Error
	return authority, err
}

// GetAuthorityList 获取角色列表
func GetAuthorityList(limit, offset int) (list interface{}, total int64, err error) {
	db := global.MPS_DB.Model(&tables.Authority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []tables.Authority
	// many2many:sys_data_authority_id多对多，其它表中的DataAuthorityId对应关系也会查出来映射其中
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

// DeleteAuthority 删除角色
func DeleteAuthority(aid uint) (err error) {
	var auth tables.Authority
	if errors.Is(global.MPS_DB.Debug().Preload("Users").Where("authority_id = ?", aid).First(&auth).Error, gorm.ErrRecordNotFound) {
		return global.ErrorAuthorityNoExist
	}
	if len(auth.Users) != 0 {
		return global.ErrorAuthorityUsing
	}
	if !errors.Is(global.MPS_DB.Where("authority_id = ?", aid).First(&tables.User{}).Error, gorm.ErrRecordNotFound) {
		return global.ErrorAuthorityUsing
	}
	if !errors.Is(global.MPS_DB.Where("parent_id = ?", aid).First(&tables.Authority{}).Error, gorm.ErrRecordNotFound) {
		return global.ErrorAuthorityHasChild
	}
	db := global.MPS_DB.Preload("BaseMenus").Where("authority_id = ?", auth.AuthorityId).First(&auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.BaseMenus) > 0 {
		err = global.MPS_DB.Model(&auth).Association("BaseMenus").Delete(auth.BaseMenus)
		if err != nil {
			return
		}
		//err = db.Association("BaseMenus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = global.MPS_DB.Delete(&[]tables.UserAuthority{}, "authority_authority_id = ?", auth.AuthorityId).Error
	return err
}

// findChildrenAuthority 递归获取角色的子角色
func findChildrenAuthority(authority *tables.Authority) (err error) {
	err = global.MPS_DB.Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

// ChangeAuthority 切换角色
func ChangeAuthority(userId uint, changeAuthority *request.ChangeAuthority) (err error) {
	return global.MPS_DB.Model(tables.User{}).Where("id = ?", userId).Update("authority_id", changeAuthority.AuthorityId).Error
}

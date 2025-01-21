package mysql

import (
	"errors"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"

	"gorm.io/gorm"
)

// UserNameExist 用户名是否存在
func UserNameExist(username string) (id uint, err error) {
	var user tables.User
	if !errors.Is(global.MPS_DB.Where("username = ?", username).First(&user).Error, gorm.ErrRecordNotFound) {
		return user.ID, global.ErrorUserExist
	}
	return 0, global.ErrorUserNotExist
}

// UserEmailExist 用户邮箱是否存在
func UserEmailExist(email string) (id uint, err error) {
	var user tables.User
	if !errors.Is(global.MPS_DB.Where("email = ?", email).First(&user).Error, gorm.ErrRecordNotFound) {
		return user.ID, global.ErrorUserEmailExist
	}
	return 0, global.ErrorUserEmailNotExist
}

// Login 用户登录
func Login(in *request.Login) (out *response.Login, err error) {
	out = new(response.Login)
	if err = global.MPS_DB.Where("username = ?", in.Username).
		Preload("Authorities").Preload("Authority").First(&out.UserInfo).Error; err == nil {
		//if ok := utils.BcryptCheck(u.Password, user.UserInfo.Password); !ok {
		//	return nil, global.ErrorInvalidPassword
		//}
		if in.Password != out.UserInfo.Password {
			return nil, global.ErrorInvalidPassword
		}
	}
	return out, err
}

// Register 创建用户
func Register(in *tables.User) (err error) {
	in.AuthorityId = 104                                                        // 默认注册用户的权限
	in.Authorities = append(in.Authorities, tables.Authority{AuthorityId: 104}) // 创建用户的时候中间表就会自动与角色对应关系
	//user.Password = utils.BcryptHash(user.Password) // 密码加密
	return global.MPS_DB.Create(in).Error
}

// GetSelfInfo 获取自身信息
func GetSelfInfo(uuid int64) (userInfo tables.User, err error) {
	err = global.MPS_DB.Where("uuid = ?", uuid).
		Preload("Authorities").Preload("Authority").First(&userInfo).Error
	return
}

// GetUserTreeMap 获取user的父子对应关系
//func GetUserTreeMap() (treeMap map[uint][]tables.User, err error) {
//	var allUsers []tables.User
//	treeMap = make(map[uint][]tables.User)
//	err = global.MPS_DB.Preload("Authorities").Find(&allUsers).Error
//	if err != nil {
//		return
//	}
//	for _, user := range allUsers {
//		treeMap[user.ParentId] = append(treeMap[user.ParentId], user)
//	}
//	return treeMap, err
//}

// GetAllUser 获取所有用户
func GetAllUser() (AllUser []response.GetAllUser, err error) {
	AllUser = make([]response.GetAllUser, 0)
	var users []tables.User
	if err = global.MPS_DB.Find(&users).Error; err != nil {
		return
	}
	for _, v := range users {
		AllUser = append(AllUser, response.GetAllUser{
			ID:          v.ID,
			AuthorityId: v.AuthorityId,
			Username:    v.Username,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Email:       v.Email,
			Department:  v.Department,
			Phone:       v.Phone,
			Address:     v.Address,
		})

	}
	return
}

// ChangePassword 修改密码
func ChangePassword(in *request.ChangePassword, UUID int64) (err error) {
	var user tables.User
	if err = global.MPS_DB.Where("uuid = ?", UUID).First(&user).Error; err != nil {
		return err
	}
	//if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
	//	return errors.New("原密码错误")
	//}
	//user.Password = utils.BcryptHash(u.NewPassword)
	if in.Password != user.Password {
		return errors.New("原密码错误")
	}
	user.Password = in.NewPassword
	err = global.MPS_DB.Save(&user).Error
	return err
}

// ChangeHeaderImg 修改头像
func ChangeHeaderImg(UUID int64, headerImg string) (err error) {
	var user tables.User
	if err = global.MPS_DB.Where("uuid = ?", UUID).First(&user).Error; err != nil {
		return err
	}
	user.HeaderImg = headerImg
	err = global.MPS_DB.Save(&user).Error
	return err
}

// ResetPassword 重置密码
func ResetPassword(UUID int64) error {
	return global.MPS_DB.Model(&tables.User{}).Where("uuid = ?", UUID).Update("password", utils.BcryptHash("123456")).Error
}

// DeleteUser 删除用户
func DeleteUser(id uint) (err error) {
	err = global.MPS_DB.Where("parent_id = ?", id).First(&tables.User{}).Error
	if err != nil {
		var user tables.User
		err = global.MPS_DB.Where("id = ?", id).Delete(&user).Error
		if err != nil {
			return err
		}
		err = global.MPS_DB.Delete(&[]tables.UserAuthority{}, "user_id = ?", id).Error
	} else {
		return global.ErrUserHasChild
	}
	return err
}

// SetUserAuthorities 设置用户权限组
func SetUserAuthorities(id uint, authorityIds []uint) error {
	return global.MPS_DB.Transaction(func(tx *gorm.DB) error { // 开启事务（要修改好几张表时可以开启事务）
		TxErr := tx.Delete(&[]tables.UserAuthority{}, "user_id = ?", id).Error // 先把原来对应的角色删掉
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []tables.UserAuthority
		for _, v := range authorityIds { // 把前端传来的角色新赋予用户
			useAuthority = append(useAuthority, tables.UserAuthority{
				UserId:               id,
				AuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&tables.User{}).Update("authority_id", authorityIds[0]).Error // 把用户的默认角色（刚登陆进去的角色）改为拥有第一个角色
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

// SetUserInfo 设置用户信息
func SetUserInfo(userInfo *tables.User) error {
	return global.MPS_DB.Updates(userInfo).Error
}

// SetUserAuthority 切换角色
func SetUserAuthority(id, authorityId uint) (err error) {
	assignErr := global.MPS_DB.Where("user_id = ? AND authority_authority_id = ?", id, authorityId).First(&tables.UserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return global.ErrUserNoAuthority
	}
	err = global.MPS_DB.Where("id = ?", id).First(&tables.User{}).Update("authority_id", authorityId).Error
	return err
}

// GetUsersIdByNames 根据用户名字获取用户ID
func GetUsersIdByNames(usersName []string) (usersId []uint, err error) {
	// 根据用户名字获取用户ID
	users := make([]*tables.User, 0)
	err = global.MPS_DB.Select("id").Where("username IN ?", usersName).Find(&users).Error
	for _, user := range users {
		usersId = append(usersId, user.MPS_MODEL.ID)
	}
	return
}

// GetUserInfoByID 根据用户id获取用户信息
func GetUserInfoByID(userID uint) (user *tables.User, err error) {
	user = new(tables.User)
	err = global.MPS_DB.Where("id = ?", userID).First(user).Error
	return
}

// GetReviewIdsByName 根据用户名字获取用户ID
func GetReviewIdsByName(usersName []string) (usersId []uint, err error) {
	// 根据用户名字获取用户ID
	usersId = make([]uint, 0)
	err = global.MPS_DB.Model(&tables.User{}).Select("id").Where("username IN ?", usersName).Find(&usersId).Error
	return
}

// GetUserIDsByPaperID 根据论文ID获取用户ID
func GetUserIDsByPaperID(paperID uint) (userIDs []uint, err error) {
	var userPaper []*tables.UserPaper
	err = global.MPS_DB.Where("paper_id = ?", paperID).Find(&userPaper).Error
	if err != nil {
		return
	}
	for _, userID := range userPaper {
		userIDs = append(userIDs, userID.UserId)
	}
	return
}

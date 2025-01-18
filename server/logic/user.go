package logic

import (
	"crypto/tls"
	"fmt"
	"log"
	"path/filepath"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gomail "gopkg.in/gomail.v2"
)

// Login 用户登录
func Login(in *request.Login) (out *response.Login, err error) {
	// 查询数据库
	if out, err = mysql.Login(in); err != nil {
		return nil, err
	}
	// 签发token
	token, err := utils.GenToken(out.UserInfo)
	if err != nil {
		return
	}
	out.Token = token
	// 返回信息
	return out, err
}

// Register 创建用户
func Register(in *request.Register) (err error) {
	// 判断用户名是已经否存在
	if _, err = mysql.UserNameExist(in.Username); err == global.ErrorUserExist {
		return
	}
	//判断邮箱是否存在
	if _, err = mysql.UserEmailExist(in.Email); err == global.ErrorUserEmailExist {
		return
	}
	// 创建用户
	user := &tables.User{
		Sex:               in.Sex,
		Username:          in.Username,
		Password:          in.Password,
		FirstName:         in.FirstName,
		LastName:          in.LastName,
		HeaderImg:         in.HeaderImg,
		Email:             in.Email,
		Department:        in.Department,
		Phone:             in.Phone,
		Address:           in.Address,
		Education:         in.Education,
		Title:             in.Title,
		Research:          in.Research,
		BlockChainAddress: in.BlockChainAddress,
		Affiliation:       in.Affiliation,
		AffiliationType:   in.AffiliationType,
	}
	// 生成uuid
	user.UUID = utils.GenID()
	// 存入数据库
	return mysql.Register(user)
}

func SendMail(in *request.SendMail) (err error) {

	host := "smtp.qq.com"
	port := 25
	userName := "2904976636@qq.com"
	password := "gmlvjlnjgbbpdcec" // qq邮箱填授权码
	log.Println(in.Verification + "111111111111111111111111111111111111111111111111")
	m := gomail.NewMessage()
	m.SetHeader("From", userName)
	m.SetHeader("To", in.MailReceiver)      // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Subject", "Varification")  // 邮件主题
	m.SetBody("text/html", in.Verification) //验证码
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {

		panic(err)
	}
	return nil
}

// // GetUserTree 获取用户树
//
//	func GetUserTree() (users []tables.User, err error) {
//		userTree, err := mysql.GetUserTreeMap() // 获取user的父子对应关系（此时还是乱序的）
//		users = userTree[0]
//		for i := 0; i < len(users); i++ {
//			err = getUserChildrenList(&users[i], userTree) // 从根节点开始遍历整理user树（从根节点开始有序）
//		}
//		return users, err
//	}

// GetAllUser 获取所有用户
func GetAllUser() (users []response.GetAllUser, err error) {
	return mysql.GetAllUser() // 获取user的父子对应关系（此时还是乱序的）

}

// // getChildrenList 生成一颗关系树
//
//	func getUserChildrenList(user *tables.User, treeMap map[uint][]tables.User) (err error) {
//		user.Children = treeMap[user.ID]
//		for i := 0; i < len(user.Children); i++ {
//			err = getUserChildrenList(&user.Children[i], treeMap)
//		}
//		return err
//	}
//
// GetSelfInfo 获取本用户信息
func GetSelfInfo(uuid int64) (userInfo tables.User, err error) {
	return mysql.GetSelfInfo(uuid)
}

// ChangePassword 修改密码
func ChangePassword(in *request.ChangePassword, uuid int64) error {
	return mysql.ChangePassword(in, uuid)
}

// ChangeHeaderImg 修改头像
func ChangeHeaderImg(c *gin.Context, uuid int64, in *request.ChangeHeaderImg) (url string, err error) {
	// 2. 生成文件名和保存路径
	filename := filepath.Base(in.Data.Filename)
	finalName := fmt.Sprintf("%d_%s", uuid, filename)
	saveFile := filepath.Join("./image/", finalName)
	// 保存文件
	if err = c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		return
	}
	// 3. 保存路径到数据库
	url = fmt.Sprintf("%s%s", global.MPS_CONFIG.ImagePath, finalName)
	if err = mysql.ChangeHeaderImg(uuid, url); err != nil {
		return
	}
	return url, nil
}

//
//// ResetPassword 重置密码
//func ResetPassword(uuid int64) error {
//	return mysql.ResetPassword(uuid)
//}
//
//// DeleteUser 删除用户
//func DeleteUser(id uint) error {
//	return mysql.DeleteUser(id)
//}

// SetUserAuthorities 设置用户权限组
func SetUserAuthorities(id uint, authorityIds []uint) error {
	return mysql.SetUserAuthorities(id, authorityIds)
}

// SetUserInfo 设置用户信息
func SetUserInfo(in *request.SetUserInfo) (err error) {
	return mysql.SetUserInfo(&tables.User{
		MPS_MODEL: global.MPS_MODEL{
			ID: in.ID,
		},
		AuthorityId:       in.AuthorityId,
		Sex:               in.Sex,
		Username:          in.Username,
		FirstName:         in.FirstName,
		LastName:          in.LastName,
		HeaderImg:         in.HeaderImg,
		Email:             in.Email,
		Department:        in.Department,
		Phone:             in.Phone,
		Address:           in.Address,
		Education:         in.Education,
		Title:             in.Title,
		Research:          in.Research,
		BlockChainAddress: in.BlockChainAddress,
		Affiliation:       in.Affiliation,
		AffiliationType:   in.AffiliationType,
	})
}

//// SetUserAuthority 切换角色
//func SetUserAuthority(user utils.BaseClaims, authorityId uint) (token string, err error) {
//	if err = mysql.SetUserAuthority(user.ID, authorityId); err != nil {
//		return
//	}
//	// 签发token
//	if token, err = utils.GenToken(tables.User{
//		MPS_MODEL: global.MPS_MODEL{
//			ID: user.ID,
//		},
//		UUID:        user.UUID,
//		Username:    user.Username,
//		AuthorityId: authorityId,
//		Name:        user.Name,
//	}); err != nil {
//		return
//	}
//	return
//}

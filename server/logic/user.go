package logic

import (
	"context"
	"fmt"
	"path/filepath"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"

	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
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
func Register(in *request.Register) (err error, id int64) {
	// 判断用户名是已经否存在
	if _, err = mysql.UserNameExist(in.Username); err == global.ErrorUserExist {
		return
	}
	//判断邮箱是否存在
	// if _, err = mysql.UserEmailExist(in.Email); err == global.ErrorUserEmailExist {
	// 	return
	// }
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
	return mysql.Register(user), user.UUID
}

// SendMail 发送邮件验证码
func SendMail(in *request.SendMail) error {
	// 检查Redis中邮箱验证码的TTL，以防止重复发送
	get := global.MPS_REDIS.Do(context.Background(), "TTL", global.REDIS_SMTP_PREFIX+in.Email)
	ttlResult, err := get.Result()
	if err != nil {
		global.MPS_LOG.Error("failed to check TTL in Redis", zap.String("email", in.Email), zap.Error(err))
		return fmt.Errorf("failed to check TTL in Redis: %v", err)
	}
	// 处理 TTL 返回值
	ttl, ok := ttlResult.(int64)
	if ok && ttl > 0 { // 键不存在
		if int64(global.SMTP_EXPIRED_TIME.Seconds())-ttl <= int64(global.SMTP_RETRY_TIME.Seconds()) {
			return global.ErrorInvalidEmailReSend{}
		}
	}

	// 生成6位数字验证码
	code, err := utils.GenerateRandomNumericCode(6)
	if err != nil {
		global.MPS_LOG.Error("failed to generate random numeric code", zap.Error(err))
		return fmt.Errorf("failed to generate random numeric code: %v", err)
	}

	// 输出验证码到控制台（仅在开发环境使用）
	global.MPS_LOG.Info("Email verification code",
		zap.String("email", in.Email),
		zap.String("code", code))

	// 构建邮件内容
	m := gomail.NewMessage()
	m.SetHeader("From", global.MPS_CONFIG.Smtp.Username)
	m.SetHeader("To", in.Email)
	m.SetHeader("Subject", "验证码")
	msg := fmt.Sprintf("您的验证码为: %s", code)
	m.SetBody("text/html", msg)

	// 连接SMTP服务器
	d := gomail.NewDialer(global.MPS_CONFIG.Smtp.Host, global.MPS_CONFIG.Smtp.Port, global.MPS_CONFIG.Smtp.Username, global.MPS_CONFIG.Smtp.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		global.MPS_LOG.Error("SendMail failed", zap.Error(err))
		return err
	}

	// 将验证信息存入Redis
	verificationInfo := map[string]interface{}{
		"code":       code,
		"attempts":   0,
		"verified":   false,
		"created_at": time.Now().Unix(),
	}

	if err := global.MPS_REDIS.HSet(context.Background(), global.REDIS_SMTP_PREFIX+in.Email, verificationInfo).Err(); err != nil {
		global.MPS_LOG.Error("redis HSet failed", zap.Error(err))
		return err
	}

	// 设置过期时间
	if err := global.MPS_REDIS.Expire(context.Background(), global.REDIS_SMTP_PREFIX+in.Email, global.SMTP_EXPIRED_TIME).Err(); err != nil {
		global.MPS_LOG.Error("redis Expire failed", zap.Error(err))
		return err
	}

	return nil
}

func VerifyMail(in *request.VerifyMail) (*request.VerifyMailResponse, error) {
	// 获取验证信息
	key := global.REDIS_SMTP_PREFIX + in.Email
	info, err := global.MPS_REDIS.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if len(info) == 0 {
		return nil, global.ErrorInvalidEmailCode{}
	}

	// 检查尝试次数
	attempts, _ := strconv.Atoi(info["attempts"])
	if attempts >= 5 {
		return nil, errors.New("验证尝试次数过多")
	}

	// 检查验证码是否匹配
	if info["code"] != in.Code {
		// 增加尝试次数
		global.MPS_REDIS.HIncrBy(context.Background(), key, "attempts", 1)
		return nil, global.ErrorInvalidEmailCode{}
	}

	// 检查是否已验证
	if info["verified"] == "true" {
		return nil, errors.New("验证码已被使用")
	}

	// 标记为已验证
	global.MPS_REDIS.HSet(context.Background(), key, "verified", true)

	// 生成验证token
	expiresAt := time.Now().Add(time.Hour).Unix()

	// 创建一个临时用户对象用于生成token
	tempUser := tables.User{
		Email: in.Email,
		UUID:  utils.GenID(),
	}

	token, err := utils.GenToken(tempUser)
	if err != nil {
		return nil, err
	}

	// 存储token
	global.MPS_REDIS.Set(context.Background(),
		global.REDIS_SMTP_PREFIX+"token:"+in.Email,
		token,
		time.Hour)

	return &request.VerifyMailResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
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
	saveFile := filepath.Join("./public/images/avatar/", finalName)
	// 保存文件
	if err = c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		return
	}
	// 3. 保存路径到数据库
	url = fmt.Sprintf("/public/images/avatar/%s", finalName)
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

func RegisterSendMPS(address string) {

}

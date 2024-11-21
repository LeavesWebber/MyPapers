package request

import "mime/multipart"

// Login 登录
type Login struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	//Captcha   string `json:"captcha" binding:"required"`   // 验证码
	//CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

// Register 注册
type Register struct {
	Sex               int    // 性别 1男 2女
	Username          string `json:"username" gorm:"unique;comment:用户登录名"`
	Password          string `json:"password"`            // 用户登录密码
	FirstName         string `json:"first_name"`          // 姓氏
	LastName          string `json:"last_name"`           // 名字
	HeaderImg         string `json:"header_img"`          // 用户头像
	Email             string `json:"email"`               // 用户邮箱
	Department        string `json:"department"`          // 工作单位
	Phone             string `json:"phone"`               // 用户手机号
	Address           string `json:"address"`             // 详细地址
	Education         string `json:"education"`           // 学历
	Title             string `json:"title"`               // 职称
	Research          string `json:"research"`            // 研究方向
	BlockChainAddress string `json:"block_chain_address"` // 区块链账户地址
	Affiliation       string `json:"affiliation"`         // 所属机构
	AffiliationType   string `json:"affiliation_type"`    // 所属机构类型
}

// ChangePassword 修改密码
type ChangePassword struct {
	UUID        int64  `json:"-"`           // 从 JWT 中提取 uuid，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// ChangeHeaderImg 修改头像
type ChangeHeaderImg struct {
	FileName string                `form:"file_name" binding:"required"`
	Data     *multipart.FileHeader `form:"data" binding:"required"`
}

// ResetPassword 重置用户密码
type ResetPassword struct {
	UUID int64 `json:"uuid" binding:"required"`
}

// DeleteUser 重置用户密码
type DeleteUser struct {
	ID uint `json:"id" binding:"required"`
}

// SetUserInfo 设置用户信息
type SetUserInfo struct {
	ID                uint   `json:"id"`
	AuthorityId       uint   `json:"authority_id"`
	Sex               int    `json:"sex"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	HeaderImg         string `json:"headerImg"`
	Email             string `json:"email"`
	Department        string `json:"department"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	Education         string `json:"education"`
	Title             string `json:"title"`
	Research          string `json:"research"`
	BlockChainAddress string `json:"block_chain_address"`
	Affiliation       string `json:"affiliation"`
	AffiliationType   string `json:"affiliation_type"`
}

// SetUserAuthorities 设置用户权限组
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authority_ids"` // 角色ID
}

// SetUserAuthority 切换角色
type SetUserAuthority struct {
	AuthorityId uint `json:"authority_id"`
}

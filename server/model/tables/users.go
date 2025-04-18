package tables

import (
	"server/global"
)

type User struct {
	global.MPS_MODEL
	UUID              int64        `json:"uuid" gorm:"comment:用户uuid"`
	AuthorityId       uint         `json:"authorityId" gorm:"default:102;comment:用户角色ID"` // 默认102普通用户
	Sex               int          `json:"sex" gorm:"comment:性别 1男 2女"`
	Username          string       `json:"username" gorm:"unique;comment:用户登录名"`
	Password          string       `json:"-" gorm:"comment:用户登录密码"`
	FirstName         string       `json:"first_name" gorm:"comment:姓氏"`
	MiddleName        string       `json:"middle_name" gorm:"comment:中间名"`
	LastName          string       `json:"last_name" gorm:"comment:名字"`
	HeaderImg         string       `json:"headerImg" gorm:"default:https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png;comment:用户头像"`
	Email             string       `json:"email" gorm:"comment:用户邮箱"`
	Department        string       `json:"department" gorm:"comment:工作单位"`
	Phone             string       `json:"phone" gorm:"comment:用户手机号"`
	Address           string       `json:"address" gorm:"comment:详细地址"`
	Education         string       `json:"education" gorm:"comment:学历"`
	Title             string       `json:"title" gorm:"comment:职称"`
	Research          string       `json:"research" gorm:"comment:研究方向"`
	BlockChainAddress string       `json:"block_chain_address" gorm:"comment:区块链账户地址"`
	Affiliation       string       `json:"affiliation" gorm:"comment:所属机构"`
	AffiliationType   string       `json:"affiliation_type" gorm:"comment:所属机构类型"`
	Authority         Authority    `json:"authority,omitempty" gorm:"references:AuthorityId;comment:用户角色"` // 用户最高的一个角色
	Authorities       []Authority  `json:"authorities" gorm:"many2many:user_authority;"`                   // 用户所拥有的角色 /创建用户时gorm创建的中间表会生成用户和角色对应关系
	Journals          []Journal    `json:"-" gorm:"many2many:user_journal;"`                               // 用户所属的期刊
	Conferences       []Conference `json:"-" gorm:"many2many:user_conference;"`                            // 用户所属的会议
	Committee         []Committee  `json:"-" gorm:"many2many:user_committee;"`                             // 用户所属的委员会
	Papers            []Paper      `json:"-" gorm:"many2many:user_paper;"`                                 // 用户的文章
}

package tables

// UserAuthority 是 User 和 Authority 的连接表
type UserAuthority struct {
	UserId               uint `gorm:"column:user_id"`
	AuthorityAuthorityId uint `gorm:"column:authority_authority_id"`
}

func (s *UserAuthority) TableName() string {
	return "user_authority"
}

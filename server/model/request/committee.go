package request

import "time"

// CreateCommittee 创建委员会
type CreateCommittee struct {
	CreatorId      uint    `json:"creator_id" comment:"创建者ID"`
	Name           string  `json:"name" comment:"委员会名称" binding:"required"`
	Description    string  `json:"description" comment:"委员会简介" binding:"required"`
	Presidents     []*Info `json:"presidents" comment:"主席信息" binding:"required"`
	VicePresidents []*Info `json:"vice_presidents" comment:"副主席信息"`
	Members        []*Info `json:"members" comment:"成员信息"`
}

// UpdateCommittee 更新委员会
type UpdateCommittee struct {
	ID             uint    `json:"id" comment:"委员会ID" binding:"required"`
	Name           string  `json:"name" comment:"委员会名称"`
	Description    string  `json:"description" comment:"委员会简介"`
	Presidents     []*Info `json:"presidents" comment:"主席信息"`
	VicePresidents []*Info `json:"vice_presidents" comment:"副主席信息"`
	Members        []*Info `json:"members" comment:"成员信息"`
}

// Info 主席/副主席信息
type Info struct {
	Name      string    `json:"name" comment:"姓名"`
	FirstName string    `json:"first_name" comment:"名"`
	LastName  string    `json:"last_name" comment:"姓"`
	HeaderImg string    `json:"header_img" comment:"头像"`
	Position  string    `json:"position" comment:"职位"`
	StartTime time.Time `json:"start_time" comment:"任职开始时间"`
	EndTime   time.Time `json:"end_time" comment:"任职结束时间"`
	Level     string    `json:"level" comment:"级别"`
}

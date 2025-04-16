package response

import "server/model/request"

// CommitteeDetail 委员会详情
type CommitteeDetail struct {
	CreatorId      uint            `json:"creator_id" comment:"创建人id"`
	Name           string          `json:"name" comment:"委员会名称"`
	Description    string          `json:"description" comment:"委员会简介"`
	Presidents     []*request.Info `json:"presidents" comment:"主席信息"`
	VicePresidents []*request.Info `json:"vice_presidents" comment:"副主席信息"`
	Members        []*request.Info `json:"members" comment:"成员信息"`
}

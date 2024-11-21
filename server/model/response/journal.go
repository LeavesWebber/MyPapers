package response

import (
	"server/model/request"
	"time"
)

// JournalDetail 期刊详情
type JournalDetail struct {
	CreatorId           uint            `json:"creator_id" comment:"创建人Id"`
	Name                string          `json:"name" comment:"期刊名称"`
	CreateAt            time.Time       `json:"create_at" comment:"创建时间"`
	Description         string          `json:"description" comment:"期刊简介"`
	Category            string          `json:"category" comment:"期刊类别"`
	IssuePeriod         string          `json:"issue_period" comment:"稿期"` // 季度刊、年刊 Quarterly Annual
	SubmissionStartTime time.Time       `json:"submission_start_time" comment:"投稿开始时间"`
	SubmissionDeadline  time.Time       `json:"submission_deadline" comment:"投稿截止时间"`
	CanSubmit           bool            `json:"can_submit" comment:"是否能够投稿"`
	Presidents          []*request.Info `json:"presidents" comment:"主席信息"`
	VicePresidents      []*request.Info `json:"vice_presidents" comment:"副主席信息"`
	Members             []*request.Info `json:"members" comment:"成员信息"`
}

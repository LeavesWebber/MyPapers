package request

type AllotReviewers struct {
	PaperId       uint     `json:"paper_id" form:"paper_id" binding:"required"`
	ReviewerNames []string `json:"reviewer_names" form:"reviewer_names" binding:"required"`
}

// SubmitReview 审核
type SubmitReview struct {
	PaperId uint   `json:"paper_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	//Hash          string `json:"hash" binding:"required"`
}

package response

// GetAuthorityList 获取角色列表
type GetAuthorityList struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

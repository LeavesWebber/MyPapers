package response

import "server/model/tables"

type Login struct {
	UserInfo tables.User `json:"userInfo"`
	Token    string      `json:"token"`
}

type GetUserTree struct {
	UserTree []tables.User `json:"user_tree"`
}
type GetAllUser struct {
	ID          uint   `json:"id"`
	AuthorityId uint   `json:"authority_id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Department  string `json:"department"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

package group

type UserGroupModal struct {
	UserId   string `json:"user_id"`
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type ReadUserGroups struct {
	List  []*UserContract `json:"list"`
	Total int             `json:"total"`
}

type UserContract struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AliasName string `json:"alias_name"`
	Avatar    string `json:"avatar"`
}

package user

type SaveUserGroupModel struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	GroupName string `json:"group_name"`
}

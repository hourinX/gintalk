package user

type UserRegisterModel struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Gender   int    `json:"gender"`
}

type UserLoginModel struct {
	UserCode string `json:"user_code" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha"`
}

type ReadUserLoginModel struct {
	Id           string `json:"id"`
	UserName     string `json:"user_name"`
	UserCode     string `json:"user_code"`
	Avatar       string `json:"avatar"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Gender       int    `json:"gender"`
	IsFrozen     int    `json:"is_frozen"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireTime   int64  `json:"expire_time"`
}

type SaveUserGroupModel struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	GroupName string `json:"group_name"`
}

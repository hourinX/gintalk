package user

type UserRegisterModel struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
}

type UserLoginModel struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha" validate:"required"`
}

package system

type SystemRefreshModel struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ReadSystemRefreshModel struct {
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"`
}

package log

type LoginLogModal struct {
	Type     int `json:"type"`
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

type ReadLoginLogListModel struct {
	List  []LoginLog `json:"list"`
	Count int        `json:"count"`
}

type LoginLog struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	LoginTime   string `json:"login_time"`
	IpAddress   string `json:"ip_address"`
	DeviceInfo  string `json:"device_info"`
	LoginStatus string `json:"login_status"`
	LoginMethod string `json:"login_method"`
	CreateTime  string `json:"create_time"`
}

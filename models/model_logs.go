package models

import "time"

type LoginLogs struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	LoginTime   time.Time `json:"login_time"`
	IpAddress   string    `json:"ip_address"`
	DeviceInfo  string    `json:"device_info"`
	LoginStatus string    `json:"login_status"`
	LoginMethod string    `json:"login_method"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

type LoginLogsWhere struct {
	Type int `json:"type"`
}

package models

import "time"

type User struct {
	Id             string    `json:"id"`
	UserCode       string    `json:"user_code"`
	UserName       string    `json:"user_name"`
	Password       string    `json:"password"`
	Phone          string    `json:"phone"`
	Avatar         string    `json:"avatar"`
	Gender         int       `json:"gender"`
	Email          string    `json:"email"`
	ImOnlineStatus int       `json:"im_online_status"`
	IsFrozen       int       `json:"is_frozen"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

type UserWhere struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	UserCode string `json:"user_code"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

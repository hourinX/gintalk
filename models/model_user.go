package models

import "time"

type User struct {
	Id        string    `json:"id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	NickName  string    `json:"nick_name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

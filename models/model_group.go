package models

import "time"

type UserContact struct {
	Id            string    `json:"id"`
	UserId        string    `json:"user_id"`
	ContactUserId string    `json:"contact_user_id"`
	Alias         string    `json:"alias"`
	Status        int       `json:"status"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

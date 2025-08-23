package daos

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"time"
)

func parseLoginLogWhere(w *models.LoginLogsWhere) (string, []interface{}) {
	where := "1=1"
	var args []interface{}

	if w.Type != 0 {
		switch w.Type {
		case 1:
			startOfWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Truncate(24 * time.Hour) // Start of this week
			where += " AND login_time >= ?"
			args = append(args, startOfWeek)

		case 2:
			startOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
			where += " AND login_time >= ?"
			args = append(args, startOfMonth)

		case 3:
			startOfYear := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local)
			where += " AND login_time >= ?"
			args = append(args, startOfYear)

		default:
			return "", nil
		}
	}
	if w.Type > 0 {
		curTime := time.Now()
		where += " AND login_time <= ?"
		args = append(args, utils.ParseDate(curTime, commons.YYYYMMDDHHMMSS))
	}

	return where, args
}

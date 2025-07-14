package daos

import "gin-online-chat-backend/models"

func parseUserWhere(w *models.UserWhere) (string, []interface{}) {
	where := "is_frozen = 1"
	var args []interface{}

	if w.Id != "" {
		where += " and id = ?"
		args = append(args, w.Id)
	}

	if w.UserName != "" {
		where += " and user_name = ?"
		args = append(args, w.UserName)
	}

	if w.UserCode != "" {
		where += " and user_code = ?"
		args = append(args, w.UserCode)
	}

	if w.Phone != "" {
		where += " and phone = ?"
		args = append(args, w.Phone)
	}

	if w.Email != "" {
		where += " and email = ?"
		args = append(args, w.Email)
	}

	return where, args

}

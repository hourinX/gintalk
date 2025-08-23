package log

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"
)

func GetLoginLogList(model *LoginLogModal) (*ReadLoginLogListModel, error) {
	w := &models.LoginLogsWhere{
		Type: model.Type,
	}
	field := "id, user_id, user_name, login_time, ip_address, device_info, login_status, login_method, create_time, update_time"
	list, count := daos.GetLoginLogList(w, field, "login_time desc", model.PageNo, model.PageSize)

	dataMap := make([]LoginLog, 0, len(list))
	for _, item := range list {
		dataMap = append(dataMap, LoginLog{
			Id:          item.Id,
			UserId:      item.UserId,
			UserName:    item.UserName,
			LoginTime:   utils.ParseDate(item.LoginTime, commons.YYYYMMDDHHMMSS),
			IpAddress:   item.IpAddress,
			DeviceInfo:  item.DeviceInfo,
			LoginStatus: item.LoginStatus,
			LoginMethod: item.LoginMethod,
			CreateTime:  utils.ParseDate(item.CreateTime, commons.YYYYMMDDHHMMSS),
		})
	}

	return &ReadLoginLogListModel{
		List:  dataMap,
		Count: count,
	}, nil
}

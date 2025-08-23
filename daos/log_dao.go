package daos

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
	"time"

	"gorm.io/gorm"
)

func GetLoginLogList(w *models.LoginLogsWhere, field, order string, pageNo, pageSize int) ([]*models.LoginLogs, int) {
	var (
		logs  []*models.LoginLogs
		count int64
	)
	limit := pageSize
	offset := (pageNo - 1) * pageSize
	where, args := parseLoginLogWhere(w)
	if limit > 0 {
		systems.DB.Table(commons.TableImUserLogs).Select(field).Order(order).Where(where, args...).Offset(offset).Limit(limit).Scan(&logs)
	} else {
		systems.DB.Table(commons.TableImUserLogs).Select(field).Order(order).Where(where, args...).Scan(&logs)
	}
	systems.DB.Table(commons.TableImUserLogs).Where(where, args...).Count(&count)
	return logs, int(count)
}

func InsertLoginLog(tx *gorm.DB, log *models.LoginLogs) error {
	if tx == nil {
		tx = systems.DB
	}

	id, err := utils.GenerateNumericID(commons.MaxIDLength)
	if err != nil {
		return err
	}
	curTime := time.Now()

	log.Id = id
	log.CreateTime = curTime
	log.UpdateTime = curTime
	return tx.Table(commons.TableImUserLogs).Create(log).Error
}

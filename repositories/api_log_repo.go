package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"strconv"
	"time"
)

type apiLogRepository struct {
}

func (lg apiLogRepository) Save(log model.ApiLog) error {

	db := infrastructure.GetDB()
	err := db.Create(&log).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func (lg apiLogRepository) FilterLog(departmentCode string, userId int, url, method string, data string, startTime *time.Time, endTime *time.Time, page int, pageSize int) ([]*model.ApiLog, error) {
	db := infrastructure.GetDB()

	var whereQuery string
	if departmentCode == "" {
		whereQuery += " logs.department_code ILIKE '%'"
	} else {
		whereQuery += " logs.department_code ILIKE '%" + departmentCode + "%'"
	}
	if data != "" {
		whereQuery += " AND logs.data ILIKE '%" + data + "%'"
	}
	if url != "" {
		whereQuery += " AND logs.url ILIKE '%" + url + "%'"
	}
	if method != "" {
		whereQuery += " AND logs.method ILIKE '%" + method + "%'"
	}
	if startTime != nil && endTime != nil {
		whereQuery += " AND logs.time >= " + "'" + startTime.Format("2006-01-02 00:00:00") + "'" + " AND logs.time <= " + "'" + endTime.Format("2006-01-02 00:00:00") + "'"
	}
	if userId > 0 {
		whereQuery += " AND logs.user_id = " + strconv.Itoa(userId)
	}
	result := []*model.ApiLog{}
	var total int64
	err := db.Table("logs").Select("*").Where(whereQuery).Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = db.Limit(pageSize).Offset((page - 1) * pageSize).
		Table("logs").Where(whereQuery).Order("time desc").Scan(&result).Error
	return result, nil
}

func NewApiLogRepository() model.ApiLogRepository {
	return apiLogRepository{}
}

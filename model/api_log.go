package model

import "time"

type ApiLog struct {
	ID             int64     `gorm:"column:id"`
	DepartmentCode string    `gorm:"column:department_code"`
	UserID         int       `gorm:"column:user_id"`
	URL            string    `gorm:"column:url"`
	Method         string    `gorm:"column:method"`
	Data           string    `gorm:"column:data"`
	Param          string    `gorm:"column:param"`
	Time           time.Time `gorm:"column:time"`
}

type ApiLogRepository interface {
	Save(l ApiLog) error
	FilterLog(departmentCode string, userId int, url, method string, data string, startTime *time.Time, endTime *time.Time, page int, pageSize int) ([]*ApiLog, error)
}

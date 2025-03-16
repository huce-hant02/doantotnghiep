package model

import "time"

type Log struct {
	ID         uint       `json:"id"`
	OutlineID  uint       `json:"outlineId"`
	EmployeeID uint       `json:"employeeId"`
	Type       string     `json:"type"`
	Content    string     `json:"content"`
	CreatedAt  *time.Time `json:"createdAt"`
	Status     string     `json:"status"`
	AuthRole   string     `json:"authRole"`

	Employee Employee `swaggerignore:"true" json:"employee" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Outline  Outline  `swaggerignore:"true" json:"-" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LogRepository interface {
	IBasicDBQueryRepo
	ISearchRepo
}

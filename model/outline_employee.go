package model

import "gorm.io/gorm"

type OutlineEmployee struct {
	ID         uint           `json:"id" gorm:"unique;autoIncrement"`
	OutlineID  uint           `json:"outlineId" gorm:"primaryKey"`
	EmployeeID uint           `json:"employeeId" gorm:"primaryKey"`
	Note       string         `json:"note"`
	Assignee   *bool          `json:"assignee"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
}

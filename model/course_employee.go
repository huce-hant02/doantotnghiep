package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseEmployee struct {
	ID         uint   `json:"id" gorm:"unique;autoIncrement"`
	EmployeeID uint   `json:"employeeId" gorm:"uniqueIndex:course_employee_uniq"`
	CourseID   uint   `json:"courseId" gorm:"uniqueIndex:course_employee_uniq"`
	Note       string `json:"note"`
	Assignee   *bool  `json:"assignee"`

	Employee *Employee `swaggerignore:"true" json:"employee" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Course   *Course   `swaggerignore:"true" json:"course" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

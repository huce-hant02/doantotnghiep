package model

import (
	"time"

	"gorm.io/gorm"
)

type EducationLog struct {
	ID          uint   `json:"id"`
	EducationID uint   `json:"educationId"`
	EmployeeID  uint   `json:"employeeId"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	AuthRole    string `json:"authRole"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	Employee  Employee         `swaggerignore:"true" json:"employee" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Education EducationProgram `swaggerignore:"true" json:"-" gorm:"foreignKey:EducationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EducationLogRepository interface {
	IBasicDBQueryRepo
	ISearchRepo
}

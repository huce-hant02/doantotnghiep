package model

import (
	"time"
)

// Giảng viên được phân công phụ trách Chương trình đào tạo CDIO
type CdioProgramEmployee struct {
	ID                 uint  `json:"id"`
	EducationProgramId uint  `json:"educationProgramId" gorm:"index:cdio_program_employee_unique,unique"`
	EmployeeID         uint  `json:"employeeId" gorm:"index:cdio_program_employee_unique,unique"`
	Disabled           *bool `json:"disabled"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	// FK
	Employee Employee `swaggerignore:"true" json:"employee" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

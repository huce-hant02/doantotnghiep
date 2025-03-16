package model

import "time"

// Sinh viên
type LoeStudent struct {
	ID         uint   `json:"id"`
	ClassCode  string `json:"classCode"`
	Code       string `json:"code" gorm:"index:student_uniq_code,unique"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	SchoolYear int    `json:"schoolYear"` // Niên khoá

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	OtherFilter LoeStudentOtherFilter `json:"otherFilter" gorm:"-"`
}

type LoeStudentOtherFilter struct {
	ListId   []int    `json:"listId"`   // IDs
	ListCode []string `json:"listCode"` // Codes
}

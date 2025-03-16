package model

import (
	"time"
)

// Thời gian đào tạo chính khoá
type TrainingTime struct {
	ID uint `json:"id"`

	Code      string   `json:"code" gorm:"uniqueIndex:training_time_uniq"` // Mã
	Title     string   `json:"title"`                                      // Tiêu đề
	NumOfYear *float64 `json:"numOfYear" gorm:"default:0"`                 // Số năm

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

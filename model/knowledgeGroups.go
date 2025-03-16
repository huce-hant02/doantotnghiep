package model

import (
	"time"

	"gorm.io/gorm"
)

// Khối kiến thức
type KnowledgeGroup struct {
	ID uint `json:"id"`

	Code      string `json:"code" gorm:"unique"`
	Title     string `json:"title"`
	FrameCode string `json:"frameCode"` // ex: A
	// NumberOfCredit float64 `json:"numberOfCredit"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

// var EDUCATION_KNOWLEDGE_GROUP = {
//   GDDC: 'Giáo dục đại cương',
//   CSN: 'Cơ sở ngành',
//   BT: 'Bổ trợ',
//   CN: 'Chuyên ngành',
//   DAKLTN: 'Thực tập, Đồ án/Khóa luận tốt nghiệp'
// };

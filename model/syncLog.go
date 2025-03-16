package model

import (
	"time"
)

// Lịch sử đồng bộ dữ liệu
type SyncLog struct {
	ID uint `json:"id"`

	Model    *string `json:"model"` // Đối tượng đồng bộ
	ApiID    uint    `json:"apiId"` // API đồng bộ
	Url      string  `json:"url"`
	Method   string  `json:"method"`
	Params   string  `json:"params"`
	Payload  string  `json:"payload"`
	Status   string  `json:"status"` // Kết quả đồng bộ
	Response string  `json:"response"`
	Error    string  `json:"error"`
	Logs     *string `json:"logs"` // Log error | data

	Api *API `json:"api" gorm:"foreignKey:ApiID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"` // API đồng bộ

	CreatedAt  time.Time  `json:"createdAt" swaggerignore:"true"` // Thời điểm bắt đầu đồng bộ
	UpdatedAt  time.Time  `json:"updatedAt" swaggerignore:"true"` // Thời điểm kết thúc đồng bộ
	FinishedAt *time.Time `json:"finishedAt"`                     // Thời điểm kết thúc đồng bộ
}

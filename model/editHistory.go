// Bảng đại diện cho Lịch sử /Yêu cầu chỉnh sửa hồ sơ
package model

import (
	"time"

	"gorm.io/gorm"
)

// Lịch sử /Yêu cầu chỉnh sửa hồ sơ
type EditHistory struct {
	ID         uint    `json:"id"`
	ModelType  *string `json:"modelType"`  // Loại hồ sơ chinh sửa (camelCase)
	ModelId    uint    `json:"modelId"`    // Id của hồ sơ bị chinh sửa
	Data       *string `json:"data"`       // Dữ liệu của hồ sơ bị chinh sửa
	Active     *bool   `json:"active"`     // Trạng thái kích hoạt của lần chỉnh sửa này
	Status     *string `json:"status"`     // Trạng thái của lần chỉnh sửa này (pending, approved, rejected)
	ModifierId *uint   `json:"modifierId"` // Id của employee thực hiện chỉnh sửa
	Note       *string `json:"note"`       // Ghi chú của lần chỉnh sửa này

	// Foreign key: Thông tin (model.Student) của Sinh viên thực hiện chỉnh sửa
	Modifier *Employee `json:"modifier" gorm:"foreignKey:ModifierId;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

type EditHistoryRepo struct {
}

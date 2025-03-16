package model

import (
	"time"

	"gorm.io/gorm"
)

// RoleFieldPermission Quyền truy cập vào từng field trong mỗi model
type RoleFieldPermission struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	RoleModelTypeId uint `json:"roleModelTypeId" gorm:"uniqueIndex:model_type_field_role_uniq"`
	FieldId         uint `json:"fieldId" gorm:"uniqueIndex:model_type_field_role_uniq"`

	ReadPermission     *bool `json:"readPermission" gorm:"default:false"`
	CreatePermission   *bool `json:"createPermission" gorm:"default:false"`
	UpdatePermission   *bool `json:"updatePermission" gorm:"default:false"`
	DeletePermission   *bool `json:"deletePermission" gorm:"default:false"`
	DownloadPermission *bool `json:"downloadPermission" gorm:"default:false"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`

	OtherFilter OtherRFPFilter `json:"otherFilter" gorm:"-"`

	// FK
	RoleModelType *RoleModelTypePermission `swaggerignore:"true" json:"-" gorm:"foreignKey:RoleModelTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Field         *ModelTypeField          `swaggerignore:"true" json:"-" gorm:"foreignKey:FieldId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type OtherRFPFilter struct {
	ListRoleModelTypeID []int `json:"listRoleModelTypeId"`
}

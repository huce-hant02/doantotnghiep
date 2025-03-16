package model

import (
	"time"

	"gorm.io/gorm"
)

type RoleModelTypePermission struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	RoleId      uint                   `json:"roleId" gorm:"uniqueIndex:model_type_role_uniq"`
	ModelTypeId uint                   `json:"modelTypeId" gorm:"uniqueIndex:model_type_role_uniq"`
	Conditions  ListRoleModelCondition `json:"conditions" gorm:"type:text"`

	ReadPermission     *bool `json:"readPermission" gorm:"default:false"`
	CreatePermission   *bool `json:"createPermission" gorm:"default:false"`
	UpdatePermission   *bool `json:"updatePermission" gorm:"default:false"`
	DeletePermission   *bool `json:"deletePermission" gorm:"default:false"`
	DownloadPermission *bool `json:"downloadPermission" gorm:"default:false"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`

	OtherFilter OtherDRWFilter `json:"otherFilter" gorm:"-"`

	// FK
	Role      *Role                 `swaggerignore:"true" json:"-" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ModelType *ModelType            `swaggerignore:"true" json:"-" gorm:"foreignKey:ModelTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Fields    []RoleFieldPermission `json:"fields" gorm:"foreignKey:RoleModelTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type OtherDRWFilter struct {
	ListRoleId    []int    `json:"listRoleId"`
	ListModelType []string `json:"listModelType"`
}

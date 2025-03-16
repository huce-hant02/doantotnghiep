// Package model Lưu lại Cấu hình Authorizer: URL - Method - Roles
package model

import "time"

type RoleAPI struct {
	ID         uint                 `json:"id"`
	RoleId     uint                 `json:"roleId" gorm:"uniqueIndex:role_api_uniq"`
	ApiId      uint                 `json:"apiId" gorm:"uniqueIndex:role_api_uniq"`
	Active     *bool                `json:"active" gorm:"default:true"`
	Conditions ListRoleAPICondition `json:"conditions" gorm:"type:text"`

	API  *API  `json:"api" gorm:"foreignKey:ApiId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role *Role `json:"role" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

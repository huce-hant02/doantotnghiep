// Package model Lưu lại Cấu hình Authorizer: URL - Method - Roles
package model

import "time"

type RoleRoute struct {
	ID      uint  `json:"id"`
	RoleId  uint  `json:"roleId" gorm:"uniqueIndex:role_route_uniq"`
	RouteId uint  `json:"routeId" gorm:"uniqueIndex:role_route_uniq"`
	Active  *bool `json:"active" gorm:"default:true"`

	Route *Route `json:"route" gorm:"foreignKey:RouteId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role  *Role  `json:"role" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

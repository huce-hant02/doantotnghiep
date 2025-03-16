package model

import "time"

type UserRole struct {
	ID uint `json:"id" gorm:"autoIncrement"`

	UserId uint `json:"userId" gorm:"uniqueIndex:user_role_uniq"`
	RoleId uint `json:"roleId" gorm:"uniqueIndex:user_role_uniq"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	// you can have those role but they can be turned on/off based on your credential
	// for ex: login as ecg-candidate => student role_id is turned "off" (false)
	Active *bool `json:"active" gorm:"default:true"`
	// FK
	Role *Role `swaggerignore:"true" json:"-" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User *User `swaggerignore:"true" json:"-" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

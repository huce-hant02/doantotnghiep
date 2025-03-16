package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Type     string `json:"type" gorm:"uniqueIndex:user_uniq"` //  NHANSU | PHONGBAN
	Username string `json:"username" gorm:"uniqueIndex:user_uniq"`
	Password string `json:"password" gorm:"-; type:varchar(255)"`
	Role     string `json:"role"`
	PuCode   string `json:"puCode"`
	Name     string `json:"name"`
	Active   *bool  `json:"active" gorm:"default:true"`
	KeepRole *bool  `json:"keepRole" gorm:"default:false"`

	UserRoles []UserRole `json:"userRoles" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
}

type UserRepository interface {
	Insert(*gorm.DB, []User) ([]User, error)
	Update(*gorm.DB, uint, *User) (*User, error)
	Delete(*gorm.DB, uint) error
	ISearchRepo
}

type UserRecord struct {
	ID       uint   `json:"id"`
	Type     string `json:"type"` //  NHANSU | PHONGBAN
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	PuCode   string `json:"puCode"`
	Name     string `json:"name"`
	Active   *bool  `json:"active" gorm:"default:true"`

	UserRoles []UserRole `json:"userRoles" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`
}

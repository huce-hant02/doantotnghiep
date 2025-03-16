package model

import (
	"time"

	"gorm.io/gorm"
)

type System struct {
	ID                 uint       `json:"id"`
	Active             *bool      `json:"active" gorm:"default:true"` // Treo thông báo gián đoạn
	NotificationActive *string    `json:"notificationActive"`         // Nội dung thông báo
	NextTimeActive     *time.Time `json:"nextTimeActive"`             // Thời gian dự kiến trở lại

	Name               string `json:"name"`
	InstitutionName    string `json:"institutionName"`    // Tên tổ chức
	InstitutionWebsite string `json:"institutionWebsite"` // URL
	InstitutionPicture string `json:"institutionPicture"` // URL

	Favicon   string `json:"favicon"`
	DarkIcon  string `json:"darkIcon"`
	LightIcon string `json:"lightIcon"`
	DarkLogo  string `json:"darkLogo"`
	LightLogo string `json:"lightLogo"`

	Copyright string `json:"copyright"`

	SidebarColor               string `json:"sidebarColor"`
	SidebarTextColor           string `json:"sidebarTextColor"`
	SidebarActiveMenuColor     string `json:"sidebarActiveMenuColor"`
	SidebarActiveMenuTextColor string `json:"sidebarActiveMenuTextColor"`
	PrimaryColor               string `json:"primaryColor"`
	SecodaryColor              string `json:"secodaryColor"`

	ForgotPassword *string `json:"forgotPassword"` // Link Quên mật khẩu
	ChangePassword *string `json:"changePassword"` // Link Đổi mật khẩu
	Guide          *string `json:"guide"`          // Link HDSD

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
}

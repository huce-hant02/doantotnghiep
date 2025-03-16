package model

import "time"

type SystemEmail struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Server   string `json:"server"`

	Active    *bool `json:"active"`
	IsDefault *bool `json:"isDefault"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

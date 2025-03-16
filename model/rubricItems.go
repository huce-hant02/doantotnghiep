package model

import (
	"time"
)

type RubricItem struct {
	ID       uint   `json:"id"`
	RubricID uint   `json:"rubricId"`
	UUID     string `json:"uuid"`

	Criteria   string  `json:"criteria"`
	Point0039  string  `json:"point0039"`
	Point4054  string  `json:"point4054"`
	Point5569  string  `json:"point5569"`
	Point7084  string  `json:"point7084"`
	Point85100 string  `json:"point85100"`
	Percent    float32 `json:"percent"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

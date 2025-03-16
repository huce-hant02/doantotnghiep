package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TemplateRubricItem struct {
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

func (tmp TemplateRubricItem) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp TemplateRubricItem) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *TemplateRubricItem) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)

	var result TemplateRubricItem
	err := json.Unmarshal([]byte(str), &result)

	*tmp = TemplateRubricItem(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp TemplateRubricItem) Value() (driver.Value, error) {
	return TemplateRubricItem(tmp).GormDataType(), nil
}

// ==================================================================================
type ListTemplateRubricItem []TemplateRubricItem

func (tmp ListTemplateRubricItem) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListTemplateRubricItem) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListTemplateRubricItem) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []TemplateRubricItem
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListTemplateRubricItem(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListTemplateRubricItem) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListTemplateRubricItem(tmp).GormDataType(), nil
}

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

type TemplateRubric struct {
	UUID string `json:"uuid"`

	Code             string `json:"code"` //  gorm:"uniqueIndex:rubric_uniq"
	Title            string `json:"title"`
	IsClinicalCourse *bool  `json:"isClinicalCourse"`

	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	RubricItems ListTemplateRubricItem `json:"rubricItems" gorm:"type:text"`
}

func (tmp TemplateRubric) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp TemplateRubric) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *TemplateRubric) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)

	var result TemplateRubric
	err := json.Unmarshal([]byte(str), &result)

	*tmp = TemplateRubric(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp TemplateRubric) Value() (driver.Value, error) {
	return TemplateRubric(tmp).GormDataType(), nil
}

// ==================================================================================
type ListTemplateRubric []TemplateRubric

func (tmp ListTemplateRubric) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListTemplateRubric) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListTemplateRubric) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []TemplateRubric
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListTemplateRubric(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListTemplateRubric) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListTemplateRubric(tmp).GormDataType(), nil
}

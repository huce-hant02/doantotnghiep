package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EducationConfig struct {
	Course       *bool `json:"course"`       // Khung chương trình
	Target       *bool `json:"target"`       // Mục tiêu
	Output       *bool `json:"output"`       // Chuẩn đầu ra
	ProgressTree *bool `json:"progressTree"` // Sơ đồ cây tiến trình
	Itu          *bool `json:"itu"`          // Ma trận ITU
	Other        *bool `json:"other"`        // Thông tin khác
}

func (tmp EducationConfig) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp EducationConfig) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *EducationConfig) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result EducationConfig
	err := json.Unmarshal([]byte(str), &result)
	*tmp = EducationConfig(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp EducationConfig) Value() (driver.Value, error) {
	return EducationConfig(tmp).GormDataType(), nil
}

// ==================================================================================
type ListEducationConfig []EducationConfig

func (tmp ListEducationConfig) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListEducationConfig) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListEducationConfig) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []EducationConfig
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListEducationConfig(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListEducationConfig) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListEducationConfig(tmp).GormDataType(), nil
}

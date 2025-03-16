package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProgramTerm struct {
	Code string `json:"code"` // Mã

	Semester    int     `json:"semester"`    // Học kỳ
	Year        int     `json:"year"`        // Năm học (thứ mấy)
	NumOfCredit float64 `json:"numOfCredit"` // Số tín chỉ
	Note        string  `json:"note"`        // Ghi chú
}

type ListProgramTerm []ProgramTerm

func (tmp ProgramTerm) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ProgramTerm) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ProgramTerm) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result ProgramTerm
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ProgramTerm(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ProgramTerm) Value() (driver.Value, error) {
	return ProgramTerm(tmp).GormDataType(), nil
}

func (tmp ListProgramTerm) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListProgramTerm) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListProgramTerm) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []ProgramTerm
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListProgramTerm(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListProgramTerm) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListProgramTerm(tmp).GormDataType(), nil
}

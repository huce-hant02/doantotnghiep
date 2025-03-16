package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListEmployeeDepartment []XProfileDepartment

func (tmp XProfileDepartment) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp XProfileDepartment) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *XProfileDepartment) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result XProfileDepartment
	err := json.Unmarshal([]byte(str), &result)
	*tmp = XProfileDepartment(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp XProfileDepartment) Value() (driver.Value, error) {
	return XProfileDepartment(tmp).GormDataType(), nil
}

func (tmp ListEmployeeDepartment) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListEmployeeDepartment) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListEmployeeDepartment) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []XProfileDepartment
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListEmployeeDepartment(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListEmployeeDepartment) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListEmployeeDepartment(tmp).GormDataType(), nil
}

package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleModelCondition struct {
	Column   string `json:"column"`   // ex: Code
	ClaimKey string `json:"claimKey"` // ex: code
}

func (tmp RoleModelCondition) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(byteString)
}

func (tmp RoleModelCondition) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *RoleModelCondition) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result RoleModelCondition
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		log.Println(err)
	}
	*tmp = RoleModelCondition(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp RoleModelCondition) Value() (driver.Value, error) {
	return RoleModelCondition(tmp).GormDataType(), nil
}

type ListRoleModelCondition []RoleModelCondition

func (tmp ListRoleModelCondition) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListRoleModelCondition) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListRoleModelCondition) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []RoleModelCondition
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListRoleModelCondition(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListRoleModelCondition) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListRoleModelCondition(tmp).GormDataType(), nil
}

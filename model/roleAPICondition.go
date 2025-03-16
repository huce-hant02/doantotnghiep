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

type RoleAPICondition struct {
	Column   *string `json:"column"`   // ex: Code
	Param    *string `json:"param"`    // ex: studentCode
	ClaimKey string  `json:"claimKey"` // ex: code
}

func (tmp RoleAPICondition) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(byteString)
}

func (tmp RoleAPICondition) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *RoleAPICondition) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result RoleAPICondition
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		log.Println(err)
	}
	*tmp = RoleAPICondition(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp RoleAPICondition) Value() (driver.Value, error) {
	return RoleAPICondition(tmp).GormDataType(), nil
}

type ListRoleAPICondition []RoleAPICondition

func (tmp ListRoleAPICondition) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListRoleAPICondition) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListRoleAPICondition) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []RoleAPICondition
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListRoleAPICondition(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListRoleAPICondition) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListRoleAPICondition(tmp).GormDataType(), nil
}

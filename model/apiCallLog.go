// Package model Log lại API (cần sử dụng Authorizer)
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

type APICallLog struct {
	Id              uint              `json:"id"`
	URL             *string           `json:"url"`
	UserId          *uint             `json:"userId"`
	Role            *string           `json:"role"`
	Method          *string           `json:"method"`
	Code            *string           `json:"code"`
	QueryParameters ListParamKeyValue `json:"queryParameters" gorm:"type:text"`
	Body            string            `json:"body"`
	CallTime        *time.Time        `json:"callTime"`

	OtherFilter OtherFilterAPICall `json:"otherFilter" gorm:"-:all"`
	Authorized  bool               `json:"authorized"`
	Error       string             `json:"error"`
	CreatedAt   *time.Time         `json:"createdAt"`
}

type OtherFilterAPICall struct {
	TimeRange  *TimePayload `json:"timeRange"`
	ListURL    []string     `json:"listURL"`
	ListUserId []uint       `json:"listUserId"`
	ListRole   []string     `json:"listRole"`
	ListMethod []string     `json:"listMethod"`
	ListCode   []string     `json:"listCode"`
	BodyPart   []string     `json:"bodyPart"`
	ErrorPart  []string     `json:"errorPart"`
}
type ParamKeyValue struct {
	ParamKey   string      `json:"paramKey"`
	ParamValue interface{} `json:"paramValue" gorm:"type:text"`
}

type ListParamKeyValue []ParamKeyValue

func (tmp ParamKeyValue) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ParamKeyValue) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ParamKeyValue) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result ParamKeyValue
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ParamKeyValue(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ParamKeyValue) Value() (driver.Value, error) {
	return ParamKeyValue(tmp).GormDataType(), nil
}

func (tmp ListParamKeyValue) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListParamKeyValue) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListParamKeyValue) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []ParamKeyValue
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListParamKeyValue(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListParamKeyValue) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListParamKeyValue(tmp).GormDataType(), nil
}

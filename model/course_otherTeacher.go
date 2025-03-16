package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourseOtherTeacher struct {
	ID            uint   `json:"id"`
	Code          string `json:"code"`
	AcademicTitle string `json:"academicTitle"`
	Degree        string `json:"degree"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Note          string `json:"note"`
	Role          string `json:"role"`
	Assignee      *bool  `json:"assignee"`
}

type ListCourseOtherTeacher []CourseOtherTeacher

func (tmp ListCourseOtherTeacher) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListCourseOtherTeacher) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListCourseOtherTeacher) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result ListCourseOtherTeacher
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListCourseOtherTeacher(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListCourseOtherTeacher) Value() (driver.Value, error) {
	return ListCourseOtherTeacher(tmp).GormDataType(), nil
}

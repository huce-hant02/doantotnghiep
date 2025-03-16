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

// Sinh viên trong lớp học phần
// Điểm thi của từng sinh viên
type LoeCourseClassStudent struct {
	ID uint `json:"id"`

	CourseClassID uint `json:"courseClassId" gorm:"index:course_class_student_uniq,unique"`
	StudentID     uint `json:"studentId" gorm:"index:course_class_student_uniq,unique"`

	Points    ListLoePoint           `json:"points" gorm:"type:text"`
	CloPoints ListLoeCloStudentPoint `json:"cloPoints" gorm:"type:text"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	// FK
	CourseClass *LoeCourseClass `json:"courseClass" gorm:"foreignKey:CourseClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Student     *LoeStudent     `json:"student" gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LoePoint struct {
	Question int     `json:"question"` // Câu hỏi số mấy
	Point    float64 `json:"point"`    // Điểm cho câu hỏi này
}

type ListLoePoint []LoePoint

func (tmp LoePoint) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp LoePoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *LoePoint) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result LoePoint
	err := json.Unmarshal([]byte(str), &result)
	*tmp = LoePoint(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp LoePoint) Value() (driver.Value, error) {
	return LoePoint(tmp).GormDataType(), nil
}

func (tmp ListLoePoint) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListLoePoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListLoePoint) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []LoePoint
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListLoePoint(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListLoePoint) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListLoePoint(tmp).GormDataType(), nil
}

/* STUDENT_CLO_POINT */

type LoeCloStudentPoint struct {
	Clo   string  `json:"clo"`   // CĐR
	Point float64 `json:"point"` // Điểm
	Rank  string  `json:"rank"`  // Xếp loại  ĐẠT-KHÔNG ĐẠT
}

type ListLoeCloStudentPoint []LoeCloStudentPoint

func (tmp LoeCloStudentPoint) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp LoeCloStudentPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *LoeCloStudentPoint) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result LoeCloStudentPoint
	err := json.Unmarshal([]byte(str), &result)
	*tmp = LoeCloStudentPoint(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp LoeCloStudentPoint) Value() (driver.Value, error) {
	return LoeCloStudentPoint(tmp).GormDataType(), nil
}

func (tmp ListLoeCloStudentPoint) GormDataType() string {
	byteString, err := json.Marshal(tmp)
	if err != nil {
		return ""
	}
	return string(byteString)
}

func (tmp ListLoeCloStudentPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{fmt.Sprintf("%v", tmp.GormDataType())},
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (tmp *ListLoeCloStudentPoint) Scan(value interface{}) error {
	str := fmt.Sprintf("%v", value)
	var result []LoeCloStudentPoint
	err := json.Unmarshal([]byte(str), &result)
	*tmp = ListLoeCloStudentPoint(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (tmp ListLoeCloStudentPoint) Value() (driver.Value, error) {
	if len(tmp) == 0 {
		return nil, nil
	}
	return ListLoeCloStudentPoint(tmp).GormDataType(), nil
}

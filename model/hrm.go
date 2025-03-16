package model

import (
	"time"

	"github.com/lib/pq"
)

type XProfileDepartment struct {
	Department            string `json:"department"`
	PositionCode          string `json:"positionCode"`
	PositionDetailID      int    `json:"positionDetailId"`
	PositionDetailCode    string `json:"positionDetailCode"`
	PositionDetailAcronym string `json:"positionDetailAcronym"`
	IsTeachingDepartment  bool   `json:"isTeachingDepartment"`
	IsMainDepartment      bool   `json:"isMainDepartment"`

	// native
	DepartmentName string         `json:"departmentName"`
	SubjectID      uint           `json:"subjectId"`
	Position       string         `json:"position"`
	PositionDetail string         `json:"positionDetail"`
	Roles          pq.StringArray `json:"roles" gorm:"type:text"`
}

type Profile struct {
	Code                  string     `json:"code" gorm:"unique_index"`
	Academic              string     `json:"academic"`
	Degree                string     `json:"degree"`
	Name                  string     `json:"name"`
	Gender                string     `json:"gender"`
	DateOfBirth           *time.Time `json:"dateOfBirth"`
	Phone                 string     `json:"phone"`
	Email                 string     `json:"email" gorm:"unique"`
	Department            string     `json:"department"`            // pud007
	MainDepartment        string     `json:"mainDepartment"`        // MainDepartment
	SpecializedDepartment string     `json:"specializedDepartment"` // Khoa chuyên môn (giảng viên)
	PositionCode          string     `json:"positionCode"`
	PositionDetailCode    string     `json:"positionDetailCode"`
	PositionDetailAcronym string     `json:"positionDetailAcronym"` // PositionDetailAcronym
	Specialize            string     `json:"specialize"`
	Status                string     `json:"status"`

	ListProfileDepartment []XProfileDepartment `json:"listProfileDepartment"`
}

type Department struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Email *string `json:"email"`
}

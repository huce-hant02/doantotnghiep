package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Employee struct {
	ID            uint           `json:"id"`
	UserID        uint           `json:"userId" gorm:"default:0"`
	Code          string         `json:"code" gorm:"unique"`
	Name          string         `json:"name"`
	Gender        string         `json:"gender"`
	AcademicTitle string         `json:"academicTitle"`
	Degree        string         `json:"degree"`
	PhoneNumber   pq.StringArray `json:"phoneNumber" gorm:"type:text"`
	EmailAddress  string         `json:"emailAddress"` //  gorm:"unique"
	// FacultyCode   string         `json:"facultyCode"`
	// FacultyID             uint                   `json:"facultyId"`
	Signature string `json:"signature"`
	// SubjectID             uint                   `json:"subjectId"`
	Avatar                string                 `json:"avatar"`
	Position              string                 `json:"position"`
	PositionCode          string                 `json:"positionCode"`
	PositionDetailCode    string                 `json:"positionDetailCode"`
	PositionDetailAcronym string                 `json:"positionDetailAcronym"` // PositionDetailAcronym
	Department            string                 `json:"department"`            // pud007
	MainDepartment        string                 `json:"mainDepartment"`        // MainDepartment
	SpecializedDepartment string                 `json:"specializedDepartment"` // khoa chuyên môn
	Specialize            string                 `json:"specialize"`
	Status                string                 `json:"status"`
	SyncedAt              *time.Time             `json:"syncedAt"`
	EmployeeDepartments   ListEmployeeDepartment `json:"employeeDepartments" gorm:"type:text"`
	IsTeacher             *bool                  `json:"isTeacher" gorm:"default:false"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`

	// foreignKey
	User *User `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Specialties []Specialty `json:"specialties" gorm:"many2many:specialty_employees;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Subject *Subject `json:"subject" gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Faculty *Faculty `json:"faculty" gorm:"foreignKey:DepartmentCode;references:PuCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// field for contain data for relation table with outline
	OtherFilter EmployeeOtherFilter `json:"otherFilter" gorm:"-"`
}

type EmployeeRepository interface {
	Create(db *gorm.DB, newEmployees []Employee) ([]Employee, error)
	Update(db *gorm.DB, ID uint, employee *Employee) (*Employee, error)
	Delete(db *gorm.DB, ID uint) error
	IAssociationOpRepo
	ISearchRepo
}

type EmployeeOtherFilter struct {
	ListId             []uint   `json:"listId"`
	ListCode           []string `json:"listCode"`
	ListDepartmentCode []string `json:"listDepartmentCode"`
}

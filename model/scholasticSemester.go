package model

import (
	"time"

	"gorm.io/gorm"
)

// hoc ky trong nam hoc (HK1, HK2)
type ScholasticSemester struct {
	ID          uint   `json:"id"`
	Code        string `json:"code" gorm:"uniqueIndex:pakage_uniq"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartAt     string `json:"startAt"`
	Status      string `json:"status"`
	Active      *bool  `json:"active" gorm:"default:true"` // Active / Deactive : Ẩn/hiện

	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`

	OutlineDeadline          *time.Time         `json:"outlineDeadline"`
	EducationProgramDeadline *time.Time         `json:"educationProgramDeadline"`
	MaxTeacherInCourse       *int               `json:"maxTeacherInCourse"`               // Số giảng viên tối đa cho mỗi học phần
	MinTeacherInCourse       *int               `json:"minTeacherInCourse"`               // Số giảng viên tối thiểu cho mỗi học phần
	MaxCriteriaInClinical    *int               `json:"maxCriteriaInClinical"`            // Số tiêu chí đánh giá tối đa cho mỗi bảng đánh giá lâm sàng
	MinCriteriaInClinical    *int               `json:"minCriteriaInClinical"`            // Số tiêu chí đánh giá tối thiểu cho mỗi bảng đánh giá lâm sàng
	MaxCdioOutcomes          *int               `json:"maxCdioOutcomes"`                  // CDIO - Số CĐR học phần tối đa
	MinCdioOutcomes          *int               `json:"minCdioOutcomes"`                  // CDIO - Số CĐR học phần tối thiểu
	MaxAbetOutcomes          *int               `json:"maxAbetOutcomes"`                  // ABET - Số CĐR học phần tối đa
	MinAbetOutcomes          *int               `json:"minAbetOutcomes"`                  // ABET - Số CĐR học phần tối thiểu
	TemplateRubrics          ListTemplateRubric `json:"templateRubrics" gorm:"type:text"` // Rubric mẫu

	ScholasticID uint `json:"scholasticId"`

	ExamTypes   []ExamType   `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BloomWords  []BloomWord  `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MajorGroups []MajorGroup `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Majors      []Major      `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseTypes []CourseType `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Courses     []Course     `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	EducationPrograms []EducationProgram `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Outlines          []Outline          `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//AbetPrograms      []AbetProgram      `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//AbetOutlines      []AbetOutline      `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DefaultStandardOutputs []DefaultStandardOutput `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DefaultITUs            []DefaultITU            `swaggerignore:"true" json:"-" gorm:"foreignKey:ScholasticSemesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ScholasticSemesterRepository interface {
	ISearchRepo
}

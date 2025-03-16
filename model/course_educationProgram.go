package model

import (
	"time"

	"github.com/lib/pq"
)

type CourseEducationProgram struct {
	ID                 uint `json:"id"`
	CourseID           uint `json:"courseId"`
	EducationProgramID uint `json:"educationProgramId"`

	KnowledgeGroup string  `json:"knowledgeGroup"`
	KnowledgeType  string  `json:"knowledgeType"`
	ElectiveGroup  *string `json:"electiveGroup"`
	Year           int     `json:"year"`     // Thuộc năm học thứ mấy (1-2-3-4...)
	Semester       *int    `json:"semester"` // Thuộc học kỳ số mấy trong năm học (1-2-3...)
	// OldSemester            *int          `json:"oldSemester"` // backup
	Index                  int           `json:"index"`
	TreeIndex              *int          `json:"treeIndex"`
	OtherPrequisiteCourses pq.Int32Array `json:"otherPrequisiteCourses" gorm:"type:integer[];default:'{}'"`
	OtherParallelCourses   pq.Int32Array `json:"otherParallelCourses" gorm:"type:integer[];default:'{}'"`
	OtherFirstlyCourses    pq.Int32Array `json:"otherFirstlyCourses" gorm:"type:integer[];default:'{}'"`
	CreatedAt              time.Time     `json:"createdAt" swaggerignore:"true"`
	UpdatedAt              time.Time     `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt              gorm.DeletedAt `json:"-" swaggerignore:"true"`

	OtherFilter OtherFilterCourseEducationProgram `json:"otherFilter" gorm:"-"`

	OrientationID   uint   `json:"orientationId"`
	OrientationType string `json:"orientationType"`

	// FK
	Orientation      *Orientation      `json:"orientation" gorm:"foreignKey:OrientationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Course           *Course           `json:"course" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationProgram *EducationProgram `json:"educationProgram" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OtherFilterCourseEducationProgram struct {
	ListCourseId           []uint `json:"listCourseId"`
	ListEducationProgramId []uint `json:"listEducationProgramId"`
}

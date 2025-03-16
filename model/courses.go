package model

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Course struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"column:scholastic_semester_id;uniqueIndex:course_unique"`
	Code                 string `json:"code" gorm:"column:code;uniqueIndex:course_unique"` //  gorm:"uniqueIndex:course_code_scholastic_semester_unique, where deleted_at is null"

	Title               string   `json:"title"`
	TitleEng            string   `json:"titleEng"`
	Type                string   `json:"type"`
	CourseTypeID        *int     `json:"courseTypeId"`
	NumOfCredit         *float32 `json:"numOfCredit" gorm:"default:0"`
	NumOfCreditLt       *float32 `json:"numOfCreditLt" gorm:"default:0"`
	NumOfCreditTh       *float32 `json:"numOfCreditTh" gorm:"default:0"`
	NumOfTest           *float32 `json:"numOfTest" gorm:"default:0"`
	NumOfTestLt         *float32 `json:"numOfTestLt" gorm:"default:0"`
	NumOfTestTh         *float32 `json:"numOfTestTh" gorm:"default:0"`
	SelfTaughtPeriod    *float32 `json:"selfTaughtPeriod" gorm:"default:0"`
	NumOfCreditClinical *float32 `json:"numOfCreditClinical" gorm:"default:0"` // STC lâm sàng
	Sector              string   `json:"sector"`
	// FacultyID         uint                   `json:"facultyId"`
	DepartmentCode    string                 `json:"departmentCode" gorm:"default:''"`
	GroupType         string                 `json:"groupType"`
	Description       string                 `json:"description"`
	RelatedCourses    []RelatedCourse        `json:"relatedCourses" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PrequisiteCourses pq.Int32Array          `json:"prequisiteCourses" gorm:"type:integer[];default:'{}'"`
	FirstlyCourses    pq.Int32Array          `json:"firstlyCourses" gorm:"type:integer[];default:'{}'"`
	ParallelCourses   pq.Int32Array          `json:"parallelCourses" gorm:"type:integer[];default:'{}'"`
	IsGeneralCourse   *bool                  `json:"isGeneralCourse"`
	IsClinicalCourse  *bool                  `json:"isClinicalCourse"`
	SubjectID         uint                   `json:"subjectId"`
	OtherTeachers     ListCourseOtherTeacher `json:"otherTeachers" gorm:"type:text; default:'[]'"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`

	// for hold data in joint table
	KnowledgeGroup  string  `json:"knowledgeGroup" gorm:"-"`
	KnowledgeType   string  `json:"knowledgeType" gorm:"-"`
	ElectiveGroup   *string `json:"electiveGroup" gorm:"-"`
	TreeIndex       *int    `json:"treeIndex" gorm:"-"`
	EducationIndex  int     `json:"educationIndex" gorm:"-"`
	Semester        int     `json:"semester" gorm:"-"`
	OrientationID   uint    `json:"orientationId" gorm:"-"`
	OrientationType string  `json:"orientationType" gorm:"-"`
	// OtherPrequisiteCourses pq.Int32Array `json:"otherPrequisiteCourses" gorm:"type:integer[];default:'{}'"`
	// OtherParallelCourses   pq.Int32Array `json:"otherParallelCourses" gorm:"type:integer[];default:'{}'"`
	// OtherFirstlyCourses    pq.Int32Array `json:"otherFirstlyCourses" gorm:"type:integer[];default:'{}'"`

	// foreignKey
	Faculty           *Faculty                 `json:"faculty" gorm:"foreignKey:DepartmentCode;references:PuCode;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	CourseType        *CourseType              `json:"courseType" gorm:"foreignKey:CourseTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EducationPrograms []CourseEducationProgram `json:"educationPrograms" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// list teacher that write outline for this course
	Teachers []CourseEmployee `json:"teachers" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Outlines []Outline        `json:"-" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Specialties []Specialty      `json:"specialties" gorm:"foreignKey:course_specialty;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Subject   *Subject   `json:"subject" gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ITUTables []ITUTable `json:"-" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// OtherFilter search
	OtherFilter OtherFilterCourse `json:"otherFilter" gorm:"-"`
}

type OtherFilterCourse struct {
	ListCode []string `json:"listCode"`
	ListId   []uint   `json:"listId"`
}

// === Hooks ===================================
func (u *Course) AfterDelete(tx *gorm.DB) (err error) {
	if err := tx.Delete(&CourseEducationProgram{CourseID: u.ID}).Error; err != nil {
		fmt.Println("err remove course-educations after delete course")
	}
	if err := tx.Delete(&Outline{CourseID: u.ID}).Error; err != nil {
		fmt.Println("err remove outlines after delete course")
	}
	return
}

package model

import (
	"time"

	"gorm.io/gorm"
)

// version and majorID are unique
type EducationProgram struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:edu_program_code_scholastic_semester_unique, where deleted_at is null"`
	Code                 string `json:"code" gorm:"uniqueIndex:edu_program_code_scholastic_semester_unique, where deleted_at is null"` // Mã chương trình

	Title            string  `json:"title"`
	TitleEng         string  `json:"titleEng"`
	PromulgateNumber string  `json:"promulgateNumber"` // Số QĐ ban hành
	MajorID          uint    `json:"majorId"`          // Ngành đào tạo - Khoa quản lý
	Version          string  `json:"version"`
	TrainingLevel    string  `json:"trainingLevel"`  // Trình độ đào tạo
	TrainingType     string  `json:"trainingType"`   // Hình thức đào tạo
	TrainingTime     float32 `json:"trainingTime"`   // Thời gian đào tạo chính khóa
	MaxTermPerYear   int     `json:"maxTermPerYear"` // Số học kỳ tối đa trong năm học

	Status      string     `json:"status"`
	StartTime   *time.Time `json:"startTime"`
	Deadline    *time.Time `json:"deadline"`
	Draft       *bool      `json:"draft"`
	SumOfCredit int        `json:"sumOfCredit"` // Tổng số tín chỉ

	EnrollmentPlan                         string `json:"enrollmentPlan"`
	GradeScale                             string `json:"gradeScale"`
	EnrollmentObject                       string `json:"enrollmentObject"`
	EnrollmentSizeAndAdmissionRequirements string `json:"enrollmentSizeAndAdmissionRequirements"`
	TrainingMethod                         string `json:"trainingMethod"`
	GraduationConditions                   string `json:"graduationConditions"`
	GraduationMakingConditions             string `json:"graduationMakingConditions"`
	AfterGraduatePosition                  string `json:"afterGraduatePosition"`
	TeachingAndEvaluatingMethod            string `json:"teachingAndEvaluatingMethod"`
	SimilarPrograms                        string `json:"similarPrograms"`
	NumOfCredit                            string `json:"numOfCredit" gorm:"default:'{}'"`
	// OldNumOfCredit                         string  `json:"oldNumOfCredit"`
	ProgressTreeImage *string `json:"progressTreeImage"`

	EducationConfig *EducationConfig `json:"educationConfig" gorm:"type:text"`
	Logs            []EducationLog   `json:"logs" gorm:"foreignKey:EducationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	// foreignKey
	Major                    *Major                       `json:"major" gorm:"foreignKey:MajorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ITUTables                []ITUTable                   `json:"ituTables" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Courses                  []CourseEducationProgram     `json:"courses" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationStandardOutputs []EducationStandardOutput    `json:"educationStandardOutputs" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EducationTargets         []EducationTarget            `json:"educationTargets" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SchoolYears              []SchoolYearEducationProgram `json:"schoolYears" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Orientations             []Orientation                `json:"orientations" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Employees []CdioProgramEmployee `json:"employees" gorm:"foreignKey:EducationProgramId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	OtherFilter OtherFilterEducationProgram `json:"otherFilter" gorm:"-"`
}

type OtherFilterEducationProgram struct {
	ListId      []uint `json:"listId"`
	ListMajorId []uint `json:"listMajorId"`
	IsAssigned  *bool  `json:"isAssigned"`
}

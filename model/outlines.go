package model

import (
	"time"

	"gorm.io/gorm"
)

// constraint version and courseID are unique
type Outline struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:scholastic_outline_uniq, where deleted_at is null"`
	Code                 string `json:"code" gorm:"uniqueIndex:scholastic_outline_uniq, where deleted_at is null"`
	// educationProgramCode string unique
	CourseCode       string    `json:"courseCode"`
	CourseID         uint      `json:"courseId"`
	Status           string    `json:"status"`
	Deadline         time.Time `json:"deadline"`
	Draft            *bool     `json:"draft" gorm:"default:false"`
	Assigner         uint      `json:"assigner"`
	EducationIndex   string    `json:"educationIndex"`
	PromulgateNumber string    `json:"promulgateNumber"`

	// Assignee                       uint       `json:"assignee"`
	Version                        string `json:"version"`
	Regulation                     string `json:"regulation"`
	Responsibility                 string `json:"responsibility"`
	RequirementsForFacilities      string `json:"requirementsForFacilities"`
	RequirementsForExtracurricular string `json:"requirementsForExtracurricular"`
	// EducationProgramID             uint       `json:"educationProgramId"`
	UpdateTimes int `json:"updateTimes"`
	// ScholasticSemesterID uint       `json:"scholasticSemesterId" gorm:"index:outline_code_scholastic_semester_update_times,unique"`
	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`

	// assignee info
	// AssigneeInfo Employee `json:"assigneeInfo" gorm:"foreignKey:Assignee;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// signature
	TeacherApproved      *bool      `json:"teacherApproved"`
	TeacherApprovedDate  *time.Time `json:"teacherApprovedDate"`
	TeacherID            uint       `json:"teacherId"`
	DeanApproved         *bool      `json:"deanApproved"`
	DeanApprovedDate     *time.Time `json:"deanApprovedDate"`
	DeanID               uint       `json:"deanId"`
	TrainingApproved     *bool      `json:"trainingApproved"`
	TrainingApprovedDate *time.Time `json:"trainingApprovedDate"`
	TrainingID           uint       `json:"trainingId"`
	SectorApproved       *bool      `json:"sectorApproved"`
	SectorApprovedDate   *time.Time `json:"sectorApprovedDate"`
	SectorID             uint       `json:"sectorId"`
	AdminApproved        *bool      `json:"adminApproved"`
	AdminApprovedDate    *time.Time `json:"adminApprovedDate"`
	AdminID              uint       `json:"adminId"`

	// OtherFilter search
	OtherFilter OtherFilter `json:"otherFilter" gorm:"-"`

	// foreign key
	// EducationProgram      *EducationProgram      `json:"educationProgram" gorm:"foreignKey:EducationProgramID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Course                *Course                `json:"course" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logs                  []Log                  `json:"logs" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseTargets         []CourseTarget         `json:"courseTargets" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rubrics               []Rubric               `json:"rubrics" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeachingPlans         []TeachingPlan         `json:"teachingPlans" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseStandardOutputs []CourseStandardOutput `json:"courseStandardOutputs" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ResultEvaluates       []ResultEvaluate       `json:"resultEvaluates" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Teachers               []Employee             `json:"teachers" gorm:"many2many:outline_employees;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Documents              []OutlineDocument      `json:"documents" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OutlineUpdateProcesses []OutlineUpdateProcess `json:"outlineUpdateProcesses" gorm:"foreignKey:OutlineID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OtherFilter struct {
	ListCourseCode []string `json:"listCourseCode"`
	ListOutlineId  []uint   `json:"listOutlineId"`
	ListCourseId   []uint   `json:"listCourseId"`
	Course         *Course  `json:"course"`
	ListEduId      []uint   `json:"listEduId"`
	IsAssigned     *bool    `json:"isAssigned"`
}
type OutlineRepository interface {
	UpdateDocumentIndexing(db *gorm.DB, outlineDoc *OutlineDocument) error
}

// hook
// func (o *Outline) AfterFind(tx *gorm.DB) (err error) {
// 	listDocumentID := []uint{}
// 	for _, doc := range o.Documents {
// 		listDocumentID = append(listDocumentID, doc.ID)
// 	}

// 	listOutlineDoc := []OutlineDocument{}
// 	if errDoc := tx.Find(&listOutlineDoc, "outline_id = ? AND document_id IN ?", o.ID, listDocumentID).Error; errDoc != nil {
// 		return
// 	}

// 	for index, doc := range o.Documents {
// 		for _, outlineDoc := range listOutlineDoc {
// 			if outlineDoc.DocumentID == doc.ID {
// 				o.Documents[index].Type = outlineDoc.Type
// 				o.Documents[index].Indexing = outlineDoc.Indexing
// 				o.Documents[index].CreatedAt = outlineDoc.CreatedAt
// 				o.Documents[index].UpdatedAt = outlineDoc.UpdatedAt
// 			}
// 		}
// 	}
// 	return nil
// }

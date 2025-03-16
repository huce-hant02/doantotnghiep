package model

import (
	"time"

	"gorm.io/gorm"
)

type Major struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"uniqueIndex:major_unique"`
	Code                 string `json:"code" gorm:"uniqueIndex:major_unique"`
	DepartmentCode       string `json:"departmentCode"`
	MajorGroupID         uint   `json:"majorGroupId"`

	Title         string         `json:"title"`
	TitleEng      string         `json:"titleEng"`
	EducationType string         `json:"educationType"`
	DiplomaName   string         `json:"diplomaName"`
	Group         string         `json:"group"`
	OpenNumber    string         `json:"openNumber"`
	OpenSigner    string         `json:"openSigner"`
	CreatedAt     time.Time      `json:"createdAt" swaggerignore:"true"`
	DeletedAt     gorm.DeletedAt `json:"-" swaggerignore:"true"`
	UpdatedAt     time.Time      `json:"updatedAt" swaggerignore:"true"`

	//foreignKey
	MajorGroup        MajorGroup         `json:"majorGroup" gorm:"foreignKey:MajorGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Faculty           *Faculty           `json:"faculty" gorm:"foreignKey:DepartmentCode;references:PuCode;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT;"`
	EducationPrograms []EducationProgram `json:"-" gorm:"foreignKey:MajorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type MajorRepository interface {
	Create(db *gorm.DB, Majors []Major) ([]Major, error)
	Update(db *gorm.DB, ID uint, major *Major) (*Major, error)
	Delete(db *gorm.DB, ID uint) error
	ISearchRepo
}

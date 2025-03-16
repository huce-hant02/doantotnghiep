package model

import (
	"time"
)

type TeachingPlanDocument struct {
	ID uint `json:"id" gorm:"unique;autoIncrement"`
	// DocumentID        uint   `json:"documentId"`
	TeachingPlanID    uint   `json:"teachingPlanId" gorm:"uniqueIndex:teaching_plan_document_uniq"`
	OutlineDocumentID uint   `json:"outlineDocumentId" gorm:"uniqueIndex:teaching_plan_document_uniq"`
	Note              string `json:"note"`
	Indexing          int    `json:"indexing"`

	// Document *Document `swaggerignore:"true" json:"document" gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OutlineDocument *OutlineDocument `swaggerignore:"true" json:"outlineDocument" gorm:"foreignKey:OutlineDocumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeachingPlan    *TeachingPlan    `swaggerignore:"true" json:"teachingPlan" gorm:"foreignKey:TeachingPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	// DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

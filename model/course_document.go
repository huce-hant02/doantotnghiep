package model

type CourseDocument struct {
	ID         uint `json:"id" gorm:"unique;autoIncrement"`
	CourseID   uint `json:"courseId" gorm:"primaryKey"`
	DocumentID uint `json:"documentId" gorm:"primaryKey"`
}

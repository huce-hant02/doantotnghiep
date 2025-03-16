package model

import (
	"os"
	"path"
	"time"

	"gorm.io/gorm"
)

type OutlineVersion struct {
	ID         uint           `json:"id"`
	CourseID   uint           `json:"courseId"`
	CourseCode string         `json:"courseCode"`
	FileURL    string         `json:"fileURL"`
	Title      string         `json:"title"`
	CreatedAt  time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt  time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt  gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (o *OutlineVersion) BeforeDelete(tx *gorm.DB) (err error) {
	var outlineVersion *OutlineVersion
	if err = tx.First(&outlineVersion, o.ID).Error; err != nil {
		return nil
	}
	root, _ := os.Getwd()
	fileURL := path.Join(path.Clean(root), path.Clean(outlineVersion.FileURL))
	_ = os.Remove(fileURL)

	return nil
}

package model

import (
	"time"
)

// Bảng nối học phần bắt buộc - song hành - học trước với 01 học phần
type RelatedCourse struct {
	ID              uint   `json:"id"`
	CourseID        uint   `json:"courseId" gorm:"uniqueIndex:related_course_uniq"`
	RelatedCourseID uint   `json:"relatedCourseId" gorm:"uniqueIndex:related_course_uniq"`
	Type            string `json:"type" gorm:"uniqueIndex:related_course_uniq"` // prequisite | parallel | firstly

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`

	Course        *Course `json:"course" gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RelatedCourse *Course `json:"relatedCourse" gorm:"foreignKey:RelatedCourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

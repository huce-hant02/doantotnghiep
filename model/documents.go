package model

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID           uint       `json:"id"`
	Code         string     `json:"code" gorm:"unique"`
	Title        string     `json:"title"`
	Author       string     `json:"author"`
	Year         string     `json:"year"`
	Publisher    string     `json:"publisher"`
	Edition      string     `json:"edition"`
	DocumentType string     `json:"documentType"` // giao trinh hay bai bao
	JournalName  string     `json:"journalName"`
	DOI          string     `json:"doi"`
	ISBN         string     `json:"isbn"`
	ISSN         string     `json:"issn"`
	Link         string     `json:"link"`
	InLibrary    *bool      `json:"inLibrary"`
	Source       string     `json:"source"`
	NumOfItems   int        `json:"numOfItems"`
	DocumentNote *string    `json:"documentNote"`
	RelatedCode  *string    `json:"relatedCode"` // Tài liệu gốc ở thư viện// để map só bản ghi chính xác cho các cuốn sách mà GV khai báo thủ công
	SyncedAt     *time.Time `json:"syncedAt"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

type DocumentRepository interface {
	Create(db *gorm.DB, documents []Document) ([]Document, error)
}

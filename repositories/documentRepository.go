package repository

import (
	"doantotnghiep/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type documentRepo struct {
}

func (d *documentRepo) Create(db *gorm.DB, documents []model.Document) ([]model.Document, error) {
	if err := db.Limit(1).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "code"}},
			DoUpdates: clause.AssignmentColumns([]string{"title", "author", "year", "publisher", "edition", "document_type", "journal_name", "doi", "isbn", "issn", "link", "source", "synced_at"}),
		}).Omit("id").CreateInBatches(&documents, 1000).Error; err != nil {
		return nil, err
	}
	// fmt.Println("sync-success")
	return documents, nil
}

func NewDocumentRepo() model.DocumentRepository {
	return &documentRepo{}
}

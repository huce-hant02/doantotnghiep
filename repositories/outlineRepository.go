package repository

import (
	"doantotnghiep/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type outlineRepo struct {
}

func (o *outlineRepo) UpdateDocumentIndexing(db *gorm.DB, outlineDoc *model.OutlineDocument) error {
	if err := db.Model(&outlineDoc).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "document_id"}, {Name: "outline_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"indexing": outlineDoc.Indexing, "type": outlineDoc.Type}),
	}).Create(outlineDoc).Error; err != nil {
		return err
	}

	return nil
}

func NewOutLineRepo() model.OutlineRepository {
	return &outlineRepo{}
}

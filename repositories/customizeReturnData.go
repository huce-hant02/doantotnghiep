package repository

import (
	"gorm.io/gorm"
)

func (b *basicQueryRepository) UpdateCustomize(db *gorm.DB, filterModel interface{}, tempDB *gorm.DB, claims map[string]interface{}) (*gorm.DB, error) {
	return tempDB, nil
}

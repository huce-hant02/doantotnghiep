package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"doantotnghiep/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type logRepository struct{}

func (l *logRepository) Add(db *gorm.DB, datas interface{}) (newDatas interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (l *logRepository) Update(db *gorm.DB, data interface{}) (updatedData interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (l *logRepository) Delete(db *gorm.DB, ID uint) (err error) {
	panic("not implemented") // TODO: Implement
}

func (l *logRepository) AdvanceFilter(filterModel interface{}, pagination model.Pagination, ignoreAssociation, includeSoftDelete bool, sort []model.SortItem) (_ interface{}, _ int, err error) {
	db := infrastructure.GetDB()
	sortQuery := ""
	for index, item := range sort {
		if index == len(sort)-1 {
			sortQuery += item.Key + " " + item.Value
		} else {
			sortQuery += item.Key + " " + item.Value + ","
		}
	}
	whereQuery := utils.GetFilterQuery(filterModel)
	if includeSoftDelete {
		db = db.Unscoped()
	}
	total := int64(0)
	db.Model(&model.Log{}).Where(whereQuery).Count(&total)

	var ListData []model.Log

	tempDB := db.Model(&model.Log{}).
		Limit(pagination.PageSize).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Where(whereQuery)

	if !ignoreAssociation {
		tempDB = tempDB.Preload(clause.Associations)
	}

	tempDB = tempDB.Order(sortQuery)
	if err = tempDB.Find(&ListData).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return ListData, int(total), nil
}

func NewLogRepository() model.LogRepository {
	return &logRepository{}
}

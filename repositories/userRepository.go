package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"doantotnghiep/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct{}

func (u *userRepository) Insert(db *gorm.DB, users []model.User) ([]model.User, error) {
	if err := db.Omit("id").Create(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Update(db *gorm.DB, userID uint, user *model.User) (*model.User, error) {
	if err := db.Model(&model.User{ID: userID}).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Delete(db *gorm.DB, userID uint) error {
	if err := db.Delete(&model.User{}, userID).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) AdvanceFilter(filterModel interface{}, pagination model.Pagination, ignoreAssociation, includeSoftDelete bool, sort []model.SortItem) (_ interface{}, _ int, err error) {
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
	db.Model(&model.User{}).Where(whereQuery).Count(&total)

	var ListData []model.User

	tempDB := db.Model(&model.User{}).
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

func NewUserRepository() model.UserRepository {
	return &userRepository{}
}

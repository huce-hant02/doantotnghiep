package repository

import (
	"doantotnghiep/model"
	"reflect"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BasicQueryRepository interface {
	Add(db *gorm.DB, ObjectType string, datas interface{}, ignoreAssociation bool, listOmitField []string) (interface{}, error)
	Update(db *gorm.DB, ObjectType string, data interface{}) (interface{}, error)
	Delete(db *gorm.DB, ObjectType string, ID uint) error
}

type basicQueryRepository struct {
}

func (b *basicQueryRepository) Add(db *gorm.DB, ObjectType string, datas interface{}, ignoreAssociation bool, listOmitField []string) (interface{}, error) {
	if ignoreAssociation {
		listOmitField = append(listOmitField, clause.Associations)
	}
	var tempData = reflect.ValueOf(datas)
	if reflect.ValueOf(datas).Kind() == reflect.Ptr {
		tempData = reflect.ValueOf(datas).Elem()
	}
	if (tempData.Kind() == reflect.Array || tempData.Kind() == reflect.Slice) && tempData.Len() <= 0 {
		return datas, nil
	}

	if err := db.Unscoped().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Omit(listOmitField...).Table(strcase.ToSnake(ObjectType)).CreateInBatches(reflect.ValueOf(datas).Interface(), 1000).Error; err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *basicQueryRepository) Update(db *gorm.DB, ObjectType string, data interface{}) (interface{}, error) {
	if err := db.Unscoped().Model(reflect.ValueOf(data).Interface()).Updates(reflect.ValueOf(data).Interface()).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (b *basicQueryRepository) Delete(db *gorm.DB, ObjectType string, ID uint) error {
	dataModel := model.MODEL_MAP[ObjectType]
	tempModel := reflect.New(reflect.TypeOf(dataModel))
	id := tempModel.Elem().FieldByName("ID")
	if id.CanSet() {
		id.SetUint(uint64(ID))
	}
	if err := db.Delete(tempModel.Interface(), ID).Error; err != nil {
		return err
	}

	return nil
}

func NewBasicQueryRepo() BasicQueryRepository {
	return &basicQueryRepository{}
}

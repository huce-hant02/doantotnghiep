package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"doantotnghiep/utils"
	"errors"
	"reflect"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BasicQueryV2Repository interface {
	Add(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, datas interface{}, ignoreAssociation bool, listOmitField []string) (interface{}, error)
	Update(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, data interface{}, isList bool) (interface{}, error)
	Delete(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, data interface{}) error
}

type basicQueryV2Repository struct {
}

func (b *basicQueryV2Repository) Add(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, datas interface{}, ignoreAssociation bool, listOmitField []string) (interface{}, error) {
	mapQuery, err := GetMapForBasicQuery(db, ObjectType, role, claims, ignoreAssociation, "Create")
	if err != nil {
		return nil, err
	}
	listOmitField = append(listOmitField, mapQuery.ListOmitField...)
	validateColumns := mapQuery.MapRoleTableValidateColumn[ObjectType]

	if ignoreAssociation {
		listOmitField = append(listOmitField, clause.Associations)
	}

	isAllPassed := true

	if reflect.TypeOf(datas).Elem().Kind() == reflect.Struct {
		if len(validateColumns) > 0 {
			hasPer := PassValidateColumns(validateColumns, claims, datas)
			if !hasPer {
				isAllPassed = false
			}
		}
	}
	if reflect.TypeOf(datas).Elem().Kind() == reflect.Slice {
		dataSlice := reflect.ValueOf(datas).Elem()
		if len(validateColumns) > 0 {
			for i := 0; i < dataSlice.Len(); i++ {
				hasPer := PassValidateColumns(validateColumns, claims, dataSlice.Index(i).Interface())
				if !hasPer {
					isAllPassed = false
					break
				}
			}
		}
	}

	if !isAllPassed {
		return nil, errors.New("No permission")
	}

	if err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Omit(listOmitField...).Table(strcase.ToSnake(ObjectType)).CreateInBatches(reflect.ValueOf(datas).Interface(), 1000).Error; err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *basicQueryV2Repository) Update(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, data interface{}, isList bool) (interface{}, error) {
	mapQuery, err := GetMapForBasicQuery(db, ObjectType, role, claims, false, "Update")
	if err != nil {
		return nil, err
	}
	if len(mapQuery.ListOmitField) > 0 {
		db.Omit(mapQuery.ListOmitField...)
	}
	validateColumns := mapQuery.MapRoleTableValidateColumn[ObjectType]

	if !isList {
		if len(validateColumns) > 0 {
			hasPer := PassValidateColumns(validateColumns, claims, data)
			if !hasPer {
				return nil, errors.New("You don't have permission to do this action")
			}
		}
		if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(reflect.ValueOf(data).Interface()).Updates(reflect.ValueOf(data).Interface()).Error; err != nil {
			infrastructure.InfoLog.Println(err)
			return nil, err
		}
		return data, nil
	} else {
		if reflect.TypeOf(data).Kind() == reflect.Slice {
			dataSlice := reflect.ValueOf(data)
			err := db.Transaction(func(tx *gorm.DB) error {
				for i := 0; i < dataSlice.Len(); i++ {
					if len(validateColumns) > 0 {
						hasPer := PassValidateColumns(validateColumns, claims, dataSlice.Index(i).Interface())
						if hasPer {
							if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Model(dataSlice.Index(i).Interface()).Updates(dataSlice.Index(i).Interface()).Error; err != nil {
								return err
							}
						}
					} else {
						if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Model(dataSlice.Index(i).Interface()).Updates(dataSlice.Index(i).Interface()).Error; err != nil {
							return err
						}
					}
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
		}

		if reflect.TypeOf(data).Elem().Kind() == reflect.Slice {
			dataSlice := reflect.ValueOf(data).Elem()
			err := db.Transaction(func(tx *gorm.DB) error {
				for i := 0; i < dataSlice.Len(); i++ {
					if len(validateColumns) > 0 {
						hasPer := PassValidateColumns(validateColumns, claims, dataSlice.Index(i).Interface())
						if hasPer {
							if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Model(dataSlice.Index(i).Interface()).Updates(dataSlice.Index(i).Interface()).Error; err != nil {
								return err //
							}
						}
					} else {
						if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Model(dataSlice.Index(i).Interface()).Updates(dataSlice.Index(i).Interface()).Error; err != nil {
							return err //
						}
					}
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
	}
	return data, nil
}

func (b *basicQueryV2Repository) Delete(db *gorm.DB, role string, claims map[string]interface{}, ObjectType string, data interface{}) error {
	objectInterface := model.MODEL_MAP[ObjectType]
	objectData := reflect.New(reflect.TypeOf(objectInterface))

	tmp := data.([]int)

	mapQuery, err := GetMapForBasicQuery(db, ObjectType, role, claims, true, "Delete")
	if err != nil {
		return err
	}
	shouldDeleteRecords := []interface{}{}
	shouldDeleteIDs := tmp

	validateColumns := mapQuery.MapRoleTableValidateColumn[ObjectType]

	if len(validateColumns) > 0 {
		if err := db.
			Model(&objectInterface).Where("id IN (" + utils.ArrayToString(tmp, ",") + ")").
			Find(&shouldDeleteRecords).Error; err != nil {
			return err
		}
		shouldDeleteIDs = []int{}
		for _, item := range shouldDeleteRecords {
			hasPer := PassValidateColumns(validateColumns, claims, data)
			if !hasPer {
				continue
			} else {
				shouldDeleteIDs = append(shouldDeleteIDs, utils.GetFieldValue(item, "id").(int))
			}
		}
	}

	if len(shouldDeleteIDs) == 0 {
		return errors.New("No records to delete")
	}
	// reflect.ValueOf(objectInterface).
	if err := db.Where("id IN (" + utils.ArrayToString(shouldDeleteIDs, ",") + ")").Delete(objectData.Interface()).Error; err != nil {
		return err
	}
	mapDelete := model.MAP_FOREIGN_KEY_DELETE[ObjectType]
	for key, element := range mapDelete {
		objectInterface := model.MODEL_MAP[key]
		objectData := reflect.New(reflect.TypeOf(objectInterface))
		if err := db.Where(element + " IN (" + utils.ArrayToString(shouldDeleteIDs, ",") + ")").Delete(objectData.Interface()).Error; err != nil {
			return err
		}
	}
	if ObjectType == "employees" {
		var listUserId []int
		query := `select user_id from employees where id in (` + utils.ArrayToString(shouldDeleteIDs, ",") + `)`
		err := db.Raw(query).Scan(&listUserId).Error
		if err != nil {
			return err
		}
		if err := db.Select(clause.Associations).Where("id IN (" + utils.ArrayToString(listUserId, ",") + ")").Delete(&model.User{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewBasicQueryV2Repo() BasicQueryV2Repository {
	return &basicQueryV2Repository{}
}

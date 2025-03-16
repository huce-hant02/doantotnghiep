package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"doantotnghiep/utils"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

type AdvanceFilterV2Repo interface {
	AdvanceFilter(objectType string, search string, filterModel interface{}, pagination model.Pagination, ignoreAssociation bool, isExactSearch bool, isUnscoped bool, joinColumn, selectColumn []string, selectPreload []string, sortOpts []model.Sort, timeFilterOpts []model.TimeFilter, claims map[string]interface{}) (_ interface{}, total int, err error, validationError []model.ValidationError)
	CountRowReturn(objectType string, filterModel interface{}) (_ int, err error)
}

type advanceFilterV2Repo struct {
	db *gorm.DB
}

func (a *advanceFilterV2Repo) switchObjectTypeAddFilter(db *gorm.DB, filterModel interface{}, tempDB *gorm.DB, objectType string, claims map[string]interface{}) (*gorm.DB, error) {
	switch objectType {
	case "outlines":
		tempDB, err := a.addFilterOutline(db, tempDB, filterModel, claims)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "educationPrograms":
		tempDB, err := a.addFilterEducationProgram(db, tempDB, filterModel, claims)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "courseEducationPrograms":
		tempDB, err := a.addFilterCourseEducationProgram(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "employees":
		tempDB, err := a.addFilterEmployee(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "courses":
		tempDB, err := a.addFilterCourse(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "abetPrograms":
		tempDB, err := a.addFilterAbetProgram(db, tempDB, filterModel, claims)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "abetOutlines":
		tempDB, err := a.addFilterAbetOutline(db, tempDB, filterModel, claims)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "abetRelationshipCourseAndPrograms":
		tempDB, err := a.addFilterAbetProgramCourse(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "abetExamMatrices":
		tempDB, err := a.addFilterAbetExamMatrix(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "loeStudents":
		tempDB, err := a.addFilterLoeStudent(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	case "loeCourseClasses":
		tempDB, err := a.addFilterLoeCourseClass(db, tempDB, filterModel)
		if err != nil {
			return nil, err
		}
		return tempDB, nil
	}
	return tempDB, nil
}

func (a *advanceFilterV2Repo) CountRowReturn(objectType string, filterModel interface{}) (_ int, err error) {
	db := infrastructure.GetDB()

	tmpWhereQuery := utils.GetFilterQueryV2(filterModel, false)

	total := int64(0)
	if err := db.Table(strcase.ToSnake(objectType)).Where(tmpWhereQuery).Count(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: non-spacing marks
}

func (a *advanceFilterV2Repo) AdvanceFilter(objectType string, search string, filterModel interface{}, pagination model.Pagination, ignoreAssociation bool, isExactSearch bool, isUnscoped bool, joinColumn, selectColumn []string, selectPreload []string, sortOpts []model.Sort, timeFilterOpts []model.TimeFilter, claims map[string]interface{}) (_ interface{}, totol int, err error, validationErr []model.ValidationError) {
	// check column per role
	role, err := getRoleFromClaims(claims)
	if err != nil {
		return nil, 0, errors.New("[ERR01] permission denied. " + err.Error()), nil
	}

	// Omit Fields
	queryUtils, err := GetMapForAdvanceFilter(a.db, objectType, role, claims, ignoreAssociation, selectPreload)
	if err != nil {
		return nil, 0, err, nil
	}

	listOmitField := queryUtils.ListOmitField
	MapRoleTableOmitField := queryUtils.MapRoleTableOmitField
	MapRoleTableValidateColumn := queryUtils.MapRoleTableValidateColumn
	MapModelTypeFieldPreload := queryUtils.MapModelTypeFieldPreload

	// VALIDATE COLS
	validateColumns := MapRoleTableValidateColumn[objectType]
	isValid := false
	if len(validateColumns) == 0 {
		isValid = true
	}

	for _column, _claimsKey := range validateColumns {
		column := _column
		claimsKey := _claimsKey
		rm := reflect.ValueOf(filterModel)
		value := reflect.Indirect(rm).FieldByName(column)
		// fmt.Println(column, fmt.Sprintf("%v", &value))
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueStr := strings.TrimSpace(fmt.Sprintf("%v", value))

		if claimsKey != "" {
			claimsValue, ok := claims[claimsKey]
			if !ok {
				return nil, 0, errors.New("[ERR02] permission denied"), nil
			}
			claimsValueStr := fmt.Sprintf("%v", claimsValue)
			// fmt.Println("value", valueStr)
			if !utils.IsListStringContains(model.ListDefaultValue, valueStr) && valueStr == claimsValueStr {
				isValid = true
			}
		}
	}
	if !isValid {
		return nil, 0, errors.New("[ERR03] permission denied"), nil
	}

	// next

	var sortOrder string
	for i := range sortOpts {
		if sortOpts[i].Key != "" && i < len(sortOpts)-1 {
			sortOrder += utils.ToSnakeCase(sortOpts[i].Key) + " " + sortOpts[i].Value + ","
		}
		if sortOpts[i].Key != "" && i == len(sortOpts)-1 {
			sortOrder += utils.ToSnakeCase(sortOpts[i].Key) + " " + sortOpts[i].Value
		}
	}
	if len(sortOpts) == 0 {
		sortOrder = "id asc"
	}

	db := infrastructure.GetDB()
	whereQuery := utils.GetFilterQueryV2(filterModel, isExactSearch)
	if len(joinColumn) > 0 {
		// convert to snack case
		listJoinStringSnackCase := []string{}
		for _, item := range joinColumn {
			listJoinStringSnackCase = append(listJoinStringSnackCase, utils.ToSnakeCase(item))
		}
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		search = strings.Replace(search, "đ", "d", -1)
		search = strings.Replace(search, "Đ", "d", -1)
		result, _, _ := transform.String(t, search)

		ilikeQuery := ""
		if isExactSearch {
			ilikeQuery = " = '" + result + "'"
		} else {
			ilikeQuery = " ILIKE '%" + result + "%'"
		}

		stringJoinColumn := ""
		for index, column := range listJoinStringSnackCase {
			if index == 0 {
				stringJoinColumn += column
			} else {
				stringJoinColumn += ",' '," + column
			}
		}
		whereQuery += " AND (select unaccent(TRIM(CONCAT(" + stringJoinColumn + ")))" + ilikeQuery + ")"
	}

	// get the actual element
	if reflect.ValueOf(filterModel).Kind() == reflect.Ptr {
		filterModel = reflect.ValueOf(filterModel).Elem().Interface()
	}

	filterModelType := reflect.TypeOf(filterModel)
	filterModelSlice := reflect.MakeSlice(reflect.SliceOf(filterModelType), 0, 0)
	datas := reflect.New(filterModelSlice.Type())
	datas.Elem().Set(filterModelSlice)
	// fmt.Println("START")

	listSelectStringSnackCase := []string{}
	for _, item := range selectColumn {
		listSelectStringSnackCase = append(listSelectStringSnackCase, utils.ToSnakeCase(item))
	}

	tmpDB := db.Table(utils.ToSnakeCase(strings.Split(objectType, "_")[0])).
		Select(strings.Join(listSelectStringSnackCase, ","))

	tmpDBTotal := db.Table(utils.ToSnakeCase(strings.Split(objectType, "_")[0])).
		Select(strings.Join(listSelectStringSnackCase, ","))

	tmpDB, err = a.switchObjectTypeAddFilter(db, filterModel, tmpDB, objectType, claims)
	if err != nil {
		fmt.Println("Death 1", err)
		return nil, 0, err, nil
	}

	tmpDBTotal, err = a.switchObjectTypeAddFilter(db, filterModel, tmpDBTotal, objectType, claims)
	if err != nil {
		fmt.Println("Death 2", err)
		return nil, 0, err, nil
	}

	tempDB := tmpDB.Limit(pagination.PageSize).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Where(whereQuery).Order(sortOrder)

	if objectType == "systems" {
		tempDB = tempDB.Omit("email_password")
	}
	tempDBTotal := tmpDBTotal.Where(whereQuery).Order(sortOrder)

	if len(listOmitField) > 0 {
		tempDB = tempDB.Omit(listOmitField...)
	}
	if !ignoreAssociation {
		listAssociationPreload := getAssociationV2(objectType, selectPreload)

		for key, value := range listAssociationPreload {
			if !utils.Contains(listOmitField, key) {
				var asModelType = MapModelTypeFieldPreload[objectType][key]

				if value != "" {
					tempDB = tempDB.Preload(key, value)
				} else {
					tempDB = tempDB.Preload(key, func(db *gorm.DB) *gorm.DB {
						return db.Omit(MapRoleTableOmitField[asModelType]...)
					})
				}
			}
		}
	}

	if isUnscoped {
		tempDB = tempDB.Unscoped()
		tempDBTotal = tempDBTotal.Unscoped()
	}

	table_name := utils.ToSnakeCase(strings.Split(objectType, "_")[0])
	for _, tFilter := range timeFilterOpts {
		var counterCol int64
		findColumnQuery := "SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE table_schema = 'public' AND table_name = '" + table_name + "' AND column_name = '" + tFilter.ColumnName + "'"
		if err := tempDBTotal.Raw(findColumnQuery).Scan(&counterCol).Error; err != nil {
			return nil, 0, err, nil
		}
		if counterCol == 0 {
			return nil, 0, errors.New("column " + tFilter.ColumnName + " not found on table " + table_name), nil
		} else {
			timeWhereQuery := tFilter.ColumnName + " >= " + tFilter.StartAt + " AND " + tFilter.ColumnName + " <= " + tFilter.EndAt
			tempDB = tempDB.Where(timeWhereQuery)
		}

	}

	if err := tempDB.Find(datas.Interface()).Error; err != nil {
		return nil, 0, err, nil
	}

	total := int64(0)
	var countColDeletedAt int64

	findColumnDeletedAtQuery := "SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE table_schema = 'public' AND table_name = '" + table_name + "' AND column_name = 'deleted_at'"
	if err := tempDBTotal.Raw(findColumnDeletedAtQuery).Scan(&countColDeletedAt).Error; err != nil {
		return nil, 0, err, nil
	}

	if countColDeletedAt != 0 {
		if err := tempDBTotal.Where("deleted_at is null").Count(&total).Error; err != nil {
			return nil, 0, err, nil
		}
	} else if countColDeletedAt == 0 {
		if err := tempDBTotal.Count(&total).Error; err != nil {
			return nil, 0, err, nil
		}
	}

	if objectType == "outlines" && !ignoreAssociation {
		// can't do validation without information from related tables
		// or we need to query again to get data
		validationErr, err := a.ValidateOutline(datas.Elem().Interface())
		if err != nil {
			return nil, 0, err, nil
		}
		return datas.Elem().Interface(), int(total), nil, validationErr

	}

	return datas.Elem().Interface(), int(total), nil, nil
}

func NewAdvanceFilterV2Repo() AdvanceFilterV2Repo {
	return &advanceFilterV2Repo{
		db: infrastructure.GetDB(),
	}
}

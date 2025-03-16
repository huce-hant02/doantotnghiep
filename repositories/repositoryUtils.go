package repository

import (
	"context"
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"doantotnghiep/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

func getAssociationV2(modelName string, selectPreload []string) (_ map[string]interface{}) {
	source := model.MODEL_LIST_ASSOCIATION_MAP_V2[modelName]
	if len(selectPreload) == 0 || (len(selectPreload) == 1 && selectPreload[0] == "*") {
		return source
	}

	slices := map[string]interface{}{}

	for k, v := range source {
		if ext, _ := utils.InArray(k, selectPreload); ext {
			slices[k] = v
		}
	}
	return slices
}

func getRoleFromClaims(accessClaims map[string]interface{}) (string, error) {
	role, ok := accessClaims["role"].(string)
	if !ok {
		infrastructure.ErrLog.Println("unauthorized")
		return "", errors.New("unauthorized")
	}

	return role, nil
}

func GetAndDecodeToken(token string) (map[string]interface{}, error) {
	if token == "" {
		return nil, nil
	}
	decodedToken, err := infrastructure.GetDecodeAuth().Decode(strings.Fields(token)[1])
	if err != nil {
		return nil, err
	}
	claims, err := decodedToken.AsMap(context.Background())
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetRole(r *http.Request) (string, map[string]interface{}, error) {
	accessCookie, errAccessCookie := r.Cookie("AccessToken")
	if accessCookie == nil {
		fmt.Println("accessCookie is nil")
		return "guest", nil, nil
	}
	if errAccessCookie != nil {
		fmt.Println("errAccessCookie", errAccessCookie)
		return "", nil, errors.New("unauthorized")
	}
	accessToken := accessCookie.Value
	accessClaims, err := GetAndDecodeToken("Bearer " + accessToken)
	if err != nil {
		return "", nil, errors.New("unauthorized")
	}
	role, ok := accessClaims["role"].(string)
	// fmt.Println("role", role)
	if !ok {
		return "", accessClaims, errors.New("unauthorized")
	}
	return role, accessClaims, nil
}

// FILTER
func QueryMapForAdvanceFilter(db *gorm.DB, objectType string, role string, claims map[string]interface{}, ignoreAssociation bool, selectPreload []string) (model.RepoQueryUtils, error) {
	var res = model.RepoQueryUtils{
		ListOmitField:              []string{},
		MapRoleTableOmitField:      make(map[string][]string, 0),
		MapRoleTableValidateColumn: make(map[string]map[string]string, 0),
		MapModelTypeFieldPreload:   make(map[string]map[string]string, 0),
	}

	modelTypes := []string{}
	// for objectType
	roleObjType := model.RoleModelTypePermission{}
	if err := db.Model(&model.RoleModelTypePermission{}).
		Preload("Fields.Field").
		Where("role_id = (SELECT id FROM roles WHERE code = ?)", role).
		Where("model_type_id = (SELECT id FROM model_types WHERE code = ?) AND read_permission = true", objectType).
		First(&roleObjType).Error; err != nil {
		//
		return res, err
	}
	res.MapModelTypeFieldPreload[objectType] = make(map[string]string, 0)
	res.MapRoleTableValidateColumn[objectType] = make(map[string]string, 0)

	for _, c := range roleObjType.Conditions {
		res.MapRoleTableValidateColumn[objectType][c.Column] = c.ClaimKey
	}
	for _, col := range roleObjType.Fields {
		if col.ReadPermission == nil {
			res.MapRoleTableOmitField[objectType] = append(res.MapRoleTableOmitField[objectType], col.Field.Name)
			res.MapModelTypeFieldPreload[objectType][col.Field.Name] = col.Field.Code
			res.ListOmitField = append(res.ListOmitField, col.Field.Name)
		}
	}

	// for associations
	if !ignoreAssociation {
		listAssociationPreload := getAssociationV2(objectType, selectPreload)
		// fmt.Println("listAssociationPreload", listAssociationPreload)
		for key := range listAssociationPreload { // Class => classes
			var asModelType = res.MapModelTypeFieldPreload[objectType][key]
			modelTypes = append(modelTypes, asModelType)
		}
	}

	roleModels := []model.RoleModelTypePermission{}
	if err := db.Model(&model.RoleModelTypePermission{}).
		Preload("Fields.Field").
		Where("role_id = (SELECT id FROM roles WHERE code = ?)", role).
		Where("model_type_id IN (SELECT id FROM model_types WHERE code IN (?)) AND read_permission = true", modelTypes).
		Find(&roleModels).Error; err != nil {
		//
		return res, err
	}

	for _, model := range roleModels {
		res.MapRoleTableOmitField[model.ModelType.Code] = make([]string, 0)
		res.MapRoleTableValidateColumn[model.ModelType.Code] = make(map[string]string, 0)

		for _, col := range model.Fields {
			if col.ReadPermission == nil {
				res.MapRoleTableOmitField[model.ModelType.Code] = append(res.MapRoleTableOmitField[model.ModelType.Code], col.Field.Name)
			}
		}
	}

	return res, nil
}

func GetMapForAdvanceFilter(db *gorm.DB, objectType string, role string, claims map[string]interface{}, ignoreAssociation bool, selectPreload []string) (*model.RepoQueryUtils, error) {
	key := infrastructure.GetDBName() + "::" + "map_filter::" + role + "::" + objectType
	var res *model.RepoQueryUtils = nil
	client := infrastructure.GetRedisClient()
	resStr, err := client.Get(key).Result()
	if err != nil {
		goto QUERY
	}
	err = json.Unmarshal([]byte(resStr), &res)
	if err != nil {
		goto QUERY
	}
	return res, nil
QUERY:
	dat, err := QueryMapForAdvanceFilter(db, objectType, role, claims, ignoreAssociation, selectPreload)
	if err != nil {
		return nil, err
	}
	jsonString, err := json.Marshal(dat)
	if err != nil {
		return nil, err
	}
	if errAccess := client.
		Set(key, jsonString, 72*time.Hour).
		Err(); errAccess != nil {
		return nil, errAccess
	}
	return &dat, nil
}

// BASIC QUERY
func QueryMapForBasicQuery(db *gorm.DB, objectType string, role string, claims map[string]interface{}, ignoreAssociation bool, key string) (model.RepoQueryUtils, error) {
	// key = Create | Update | Delete
	fieldPerName := key + "Permission"
	modelPerCol := utils.ToSnakeCase(fieldPerName)

	var res = model.RepoQueryUtils{
		ListOmitField:              []string{},
		MapRoleTableOmitField:      make(map[string][]string, 0),
		MapRoleTableValidateColumn: make(map[string]map[string]string, 0),
		MapModelTypeFieldPreload:   make(map[string]map[string]string, 0),
	}

	// modelTypes := []string{}
	// for objectType
	roleObjType := model.RoleModelTypePermission{}
	if err := db.Model(&model.RoleModelTypePermission{}).
		Preload("Fields.Field").
		Where("role_id = (SELECT id FROM roles WHERE code = ?)", role).
		Where("model_type_id = (SELECT id FROM model_types WHERE code = ?) AND "+modelPerCol+" = true", objectType).
		First(&roleObjType).Error; err != nil {
		//
		fmt.Println("die", err)
		return res, err
	}
	res.MapModelTypeFieldPreload[objectType] = make(map[string]string, 0)
	res.MapRoleTableValidateColumn[objectType] = make(map[string]string, 0)
	for _, c := range roleObjType.Conditions {
		res.MapRoleTableValidateColumn[objectType][c.Column] = c.ClaimKey
	}
	for _, col := range roleObjType.Fields {
		if utils.GetFieldValue(col, fieldPerName) == nil {
			res.MapRoleTableOmitField[objectType] = append(res.MapRoleTableOmitField[objectType], col.Field.Name)
			res.MapModelTypeFieldPreload[objectType][col.Field.Name] = col.Field.Code
			res.ListOmitField = append(res.ListOmitField, col.Field.Name)
		}
	}
	return res, nil
}

func GetMapForBasicQuery(db *gorm.DB, objectType string, role string, claims map[string]interface{}, ignoreAssociation bool, key string) (*model.RepoQueryUtils, error) {
	redisKey := infrastructure.GetDBName() + "::" + "map_query::" + role + "::" + objectType
	var res *model.RepoQueryUtils = nil
	client := infrastructure.GetRedisClient()
	resStr, err := client.Get(redisKey).Result()
	if err != nil {
		goto QUERY
	}
	err = json.Unmarshal([]byte(resStr), &res)
	if err != nil {
		goto QUERY
	}
	return res, nil
QUERY:
	dat, err := QueryMapForBasicQuery(db, objectType, role, claims, ignoreAssociation, key)
	if err != nil {
		return nil, err
	}
	jsonString, err := json.Marshal(dat)
	if err != nil {
		return nil, err
	}
	if errAccess := client.
		Set(redisKey, jsonString, 72*time.Hour).
		Err(); errAccess != nil {
		return nil, errAccess
	}
	return &dat, nil
}

func PassValidateColumns(validateColumns map[string]string, claims map[string]interface{}, data interface{}) bool {
	isValid := false
	if len(validateColumns) == 0 {
		isValid = true
	}

	for _column, _claimsKey := range validateColumns {
		column := _column
		claimsKey := _claimsKey
		rm := reflect.ValueOf(data)
		value := reflect.Indirect(rm).FieldByName(column)
		// fmt.Println(column, fmt.Sprintf("%v", &value))
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueStr := strings.TrimSpace(fmt.Sprintf("%v", value))

		if claimsKey != "" {
			claimsValue, ok := claims[claimsKey]
			if !ok {
				return false
			}
			claimsValueStr := fmt.Sprintf("%v", claimsValue)
			// fmt.Println("value", valueStr)
			if !utils.IsListStringContains(model.ListDefaultValue, valueStr) && valueStr == claimsValueStr {
				isValid = true
			}
		}
	}
	return isValid
}

func (a *advanceFilterV2Repo) ValidateOutline(outlines interface{}) ([]model.ValidationError, error) {
	res := make([]model.ValidationError, 0)
	if reflect.TypeOf(outlines) != reflect.TypeOf([]model.Outline{}) {
		return nil, errors.New("type mismatch: need []model.Outline, got " + reflect.TypeOf(outlines).String())
	}
	ols := outlines.([]model.Outline)
	for _, o := range ols {
		validateErr := model.ValidationError{
			ID:    o.ID,
			Error: make(map[string]interface{}),
		}
		type1 := "Thông tin học phần"
		type2 := "Quy định học phần"
		type3 := "Chuẩn đầu ra học phần"
		//type4 := "Mục tiêu học phần"
		// type5 := "Bài đánh giá"
		// type6 := "Nội dung, kế hoạch giảng dạy"
		// type7 := "Quy định cho sinh viên"

		/*
			1. Thông tin học phần
		*/
		type1Errors := make([]string, 0)
		type2Errors := make([]string, 0)
		type3Errors := make([]string, 0)
		type4Errors := make([]string, 0)
		if o.Course != nil {
			if o.Course.CourseTypeID != nil {
				if *o.Course.CourseTypeID == 0 {
					type1Errors = append(type1Errors, "Chưa có phân loại học phần")
				}
			} else {
				type1Errors = append(type1Errors, "Chưa có phân loại học phần")
			}
			if len(o.Course.Teachers) == 0 {
				type1Errors = append(type1Errors, "Chưa phân công giảng viên cho học phần")
			}
			if o.Course.Description == "" {
				type1Errors = append(type1Errors, "Chưa có mô tả học phần")
			}
			if o.Course.CourseType != nil {
				if o.Course.CourseType.EvaluateForm == 1 {
					if o.Course.NumOfTest == nil {
						type1Errors = append(type1Errors, "Học phần chưa khai báo số bài kiểm tra")
					} else {
						if *o.Course.NumOfTest == 0 {
							type1Errors = append(type1Errors, "Học phần chưa khai báo số bài kiểm tra")
						}
					}
				} else {
					if o.Course.NumOfTest != nil {
						if *o.Course.NumOfTest > 0 {
							type1Errors = append(type1Errors, "Học phần không cần khai báo số bài kiểm tra. Hãy đặt số bài kiểm tra về 0")
						}
					}
				}
			}

			listTeacher := o.Course.Teachers
			if len(listTeacher) > 1 {
				checkPhuTrach := 0
				for _, t := range listTeacher {
					if t.Note == "Phụ trách" {
						checkPhuTrach++
					}
				}
				if checkPhuTrach == len(listTeacher) {
					type1Errors = append(type1Errors, "Học phần cần chỉ được phép có 1 giảng viên phụ trách biên soạn đề cương")
				}
			}
			if len(listTeacher) < 2 {
				type1Errors = append(type1Errors, "Học phần cần 1 giảng viên phụ trách và tối thiểu 1 giảng viên tham gia biên soạn đề cương")
			}

			for _, t := range listTeacher {
				if t.Employee != nil {
					if len(t.Employee.PhoneNumber) == 0 || t.Employee.EmailAddress == "" || t.Note == "" {
						type1Errors = append(type1Errors, "Thông tin về giảng viên cần được cung cấp đầy đủ (Học hàm/học vị - Họ tên - SĐT - Email - Vai trò")
					}
				}
			}

			// Nếu numOfTest != numOfTestLt + numOfTestTh
			if *o.Course.NumOfCreditLt+*o.Course.NumOfCreditTh != *o.Course.NumOfTest {
				type1Errors = append(type1Errors, "Tổng số bài kiểm tra phải bằng số bài lý thuyết + số bài thực hành")
			}

		}

		validateErr.Error[type1] = type1Errors

		/*
			Type 2: Quy định học phần
		*/

		if len(o.Documents) >= 2 {
			numPrimary, numSecondary := 0, 0
			for _, d := range o.Documents {
				if d.Type == "primary" {
					numPrimary++
				}
				if d.Type == "secondary" {
					numSecondary++
				}
			}
			if numPrimary == 0 || numSecondary == 0 {
				type2Errors = append(type2Errors, "Thiếu tài liệu tham khảo")
			}
		} else {
			type2Errors = append(type2Errors, "Thiếu tài liệu tham khảo")
		}

		if o.RequirementsForFacilities == "" {
			type2Errors = append(type2Errors, "Chưa có quy định về cơ sở vật chất, trang thiết bị phục vụ dạy học")
		}
		validateErr.Error[type2] = type2Errors

		/*
			Type 3: Chuẩn đầu ra học phần
		*/
		listOutput := o.CourseStandardOutputs
		if len(listOutput) == 0 {
			type3Errors = append(type3Errors, "Chưa có chuẩn đầu ra")
		} else {
			indexingCheck := utils.DuplicateFieldValue(listOutput, "Indexing")
			if indexingCheck {
				type3Errors = append(type3Errors, "Trùng Indexing!")
			}
			for _, o := range listOutput {
				if m, _ := regexp.MatchString(`/^\d\d*([.]\d\d*[A-Za-z]?)?$/`, o.Indexing); !m {
					type3Errors = append(type3Errors, "Regex không match")
				}
			}
			lstId := make([]uint, 0)
			for _, o := range listOutput {
				lstId = append(lstId, o.ID)
			}
		}
	Type3Bloom:
		for _, t := range o.CourseTargets {
			if t.KeywordID == 0 {
				type3Errors = append(type3Errors, "Chưa sử dụng động từ bloom")
				break Type3Bloom
			}
		}

		validateErr.Error[type3] = type3Errors

		/*
			Type 4: Mục tiêu học phần
		*/
		lstTarget := o.CourseTargets
		if len(lstTarget) == 0 {
			type4Errors = append(type4Errors, "Chưa có mục tiêu học phần")
		} else {
			titleCheck := utils.DuplicateFieldValue(lstTarget, "Title")
			if titleCheck {
				type4Errors = append(type4Errors, "Mục tiêu bị trùng title")
			}
		Type4Bloom:
			for _, t := range o.CourseTargets {
				if t.KeywordID == 0 {
					type4Errors = append(type4Errors, "Chưa sử dụng động từ bloom")
					break Type4Bloom
				}

			}

		}

		/*
			Type 5: Bài đánh giá
		*/

		res = append(res, validateErr)

	}

	return res, nil
}

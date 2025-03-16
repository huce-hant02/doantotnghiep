package utils

import (
	"backend/doantotnghiep/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/iancoleman/strcase"
	"github.com/lib/pq"
	"github.com/lithammer/shortuuid"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	LayoutISO     = "2006-01-02"
	MAXCONCURENCY = 1024
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func GenCode() string {
	id := shortuuid.New()
	return strings.ToUpper(id[0:10])
}

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()

		if tag != "" && tag != "-" {
			tag = strcase.ToSnake(tag)
			res[tag] = field
		}

	}
	return res
}

func StructToMapV2(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		var field interface{}
		if reflectValue.Field(i).Kind() == reflect.Ptr && !reflectValue.Field(i).IsNil() {
			field = reflectValue.Field(i).Elem()
		} else {
			field = reflectValue.Field(i).Interface()
		}
		if tag != "" && tag != "-" {
			tag = strcase.ToSnake(tag)
			res[tag] = field
		}

	}
	return res
}

func StructToMapType(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			tag = strcase.ToSnake(tag)

			res[tag] = reflect.TypeOf(field)
		}
	}
	return res
}

func InterfaceToFilterQueryV2(typeMap, valueMap map[string]interface{}, isExactSearch bool) string {
	query := "id > 0"
	for key, val := range typeMap {
		tmpVal := fmt.Sprintf("%v", val)
		if tmpVal == "int" || tmpVal == "float32" || tmpVal == "float64" || tmpVal == "uint" || tmpVal == "*int" || tmpVal == "*float32" || tmpVal == "*float64" || tmpVal == "*uint" {
			tmp := fmt.Sprintf("%v", valueMap[key])
			if tmp != "<nil>" && tmp != "0" && tmp != "99999999" {
				query += " AND \"" + key + "\" = " + tmp
			}
		}
		if tmpVal == "bool" || tmpVal == "*bool" {
			tmp := fmt.Sprintf("%v", valueMap[key])

			if tmp != "<nil>" {
				query += " AND \"" + key + "\" = " + tmp
			}
		}
		if tmpVal == "string" || tmpVal == "*string" {
			tmp := fmt.Sprintf("%v", valueMap[key])
			switch tmp {
			case "":
				break
			case "<nil>":
				break
			default:
				if isExactSearch {
					query += " AND \"" + key + "\" = '" + tmp + "'"
					break
				}
				tmp = strings.Replace(tmp, "đ", "d", -1)
				tmp = strings.Replace(tmp, "Đ", "d", -1)
				t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
				result, _, _ := transform.String(t, tmp)
				query += " AND unaccent(" + key + ") ILIKE '%" + result + "%'"
				break
			}
		}
		if tmpVal == "pq.StringArray" {
			tmp := []string{}
			for index := range valueMap[key].(pq.StringArray) {
				tmp = append(tmp, valueMap[key].(pq.StringArray)[index])
			}

			if len(tmp) > 0 {
				queryValue := strings.Join(tmp, ",")
				query += " AND \"" + key + "\" @> '{" + queryValue + "}'"
			}
		}

		if tmpVal == "pq.IntArray" {
			tmp := fmt.Sprintf("%v", valueMap[key])

			if tmp != "[]" {
				query += " AND \"" + key + "\" @> '" + strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmp, "[", "{"), "]", "}"), " ", ",") + "'"
			}
		}
		if strings.Contains(tmpVal, "time.Time") {
			tmp := fmt.Sprintf("%v", valueMap[key])
			if tmp != "<nil>" && StringTimeToString(tmp) != "0001-01-01 00:00:00" {
				query += " AND \"" + key + "\"::date = '" + StringTimeToString(tmp) + "'"
			}
		}
	}
	return query
}
func InterfaceToFilterQuery(typeMap, valueMap map[string]interface{}, modelName ...string) string {
	query := "id > 0"
	if len(modelName) > 0 {
		query = modelName[0] + "." + query
	}
	for key, val := range typeMap {
		tmpVal := fmt.Sprintf("%v", val)
		if tmpVal == "int" || tmpVal == "float32" || tmpVal == "float64" || tmpVal == "uint" {
			tmp := fmt.Sprintf("%v", valueMap[key])
			if tmp != "0" {
				query += " AND \"" + key + "\" = " + tmp
			}
		}
		if tmpVal == "bool" {
			tmp := fmt.Sprintf("%v", valueMap[key])
			if tmp != "<nil>" {
				query += " AND \"" + key + "\" = " + tmp
			}
		}
		if tmpVal == "*bool" {
			check := reflect.ValueOf(valueMap[key])
			tmp := fmt.Sprintf("%v", check.Elem())
			if !check.IsNil() {
				query += " AND \"" + key + "\" = " + tmp
			}
		}
		if tmpVal == "string" {
			tmp := fmt.Sprintf("%v", valueMap[key])
			switch tmp {
			case "":
				break
			case "<nil>":
				break
			default:
				if key == "title" {
					// query += " AND \"" + key + "\" % '" + tmp + "'"
					break
				}
				query += " AND \"" + key + "\" ILIKE '%" + tmp + "%'"
				break
			}
		}
		// if tmpVal == "utils.IntArray" {
		// 	tmp := fmt.Sprintf("%v", valueMap[key])
		// 	if tmp != "[]" {
		// 		query += " AND \"" + key + "\" @> '" + strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmp, "[", "{"), "]", "}"), " ", ",") + "'"
		// 	}
		// }
		// if tmpVal == "model.Point" {
		// 	tmp := fmt.Sprintf("%v", valueMap[key])
		// 	if tmp != "{0 0}" {
		// 		query += " AND \"" + key + "\" <-> '" + strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmp, "{", "("), "}", ")"), " ", ",") + "' < 0.01"
		// 	}
		// }
		if strings.Contains(tmpVal, "time.Time") {
			tmp := fmt.Sprintf("%v", valueMap[key])
			if tmp != "<nil>" && StringTimeToString(tmp) != "0001-01-01 00:00:00" {
				query += " AND \"" + key + "\"::date = '" + StringTimeToString(tmp) + "'"
			}
		}

	}
	return query
}

func GetFilterQuery(model interface{}, modelName ...string) string {
	typeMap := StructToMapType(model)
	valueMap := StructToMap(model)
	return InterfaceToFilterQuery(typeMap, valueMap, modelName...)
}

func GetFilterQueryV2(model interface{}, isExactSearch bool) string {
	typeMap := StructToMapType(model)
	valueMap := StructToMapV2(model)
	return InterfaceToFilterQueryV2(typeMap, valueMap, isExactSearch)
}

func StringTimeToString(s string) string {
	tmpTime, err := time.Parse("2006-01-02 15:05:05 +0000 UTC", s)
	if err != nil {
		fmt.Println(err)
	}
	return tmpTime.Format("2006-01-02 15:04:05")
}

func pgStringToArray(s string) pq.StringArray {
	s = strings.TrimLeft(s, "{")
	s = strings.TrimRight(s, "}")

	return strings.Split(s, ",")
}

func RowToObj(rows *sql.Rows, tableName string, output interface{}) error {
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {
		scanArgs := make([]interface{}, count)
		listStringArrayCol := []string{}
		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4", "INT8":
				scanArgs[i] = new(sql.NullInt64)
				break
			case "_TEXT":
				scanArgs[i] = new(sql.NullString)
				listStringArrayCol = append(listStringArrayCol, strcase.ToLowerCamel(v.Name()))
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return err
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[strcase.ToLowerCamel(v.Name())] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[strcase.ToLowerCamel(v.Name())] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[strcase.ToLowerCamel(v.Name())] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[strcase.ToLowerCamel(v.Name())] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[strcase.ToLowerCamel(v.Name())] = z.Int32
				continue
			}

			masterData[strcase.ToLowerCamel(v.Name())] = scanArgs[i]
		}

		dataModel := reflect.New(reflect.TypeOf(model.MODEL_MAP[tableName])).Interface()

		//convert string to string array in master data
		for _, column := range listStringArrayCol {
			masterData[column] = pgStringToArray(masterData[column].(string))
		}
		err = mapstructure.Decode(masterData, dataModel)
		if err != nil {
			return err
		}
		finalRows = append(finalRows, dataModel)
	}

	z, err := json.Marshal(finalRows)
	if err != nil {
		return err
	}

	err = json.Unmarshal(z, &output)
	if err != nil {
		return err
	}

	return nil
}

func InitializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			InitializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			InitializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// PatternGet using in get keys
func PatternGet(id uint) string {
	return strconv.Itoa(int(id)) + "-:--*"
}

// GetPattern get start pattern for redis saved data
func GetPattern(id uint) string {
	return strconv.Itoa(int(id)) + "-:--"
}

func FirstNonNil(datas ...interface{}) interface{} {
	for _, v := range datas {
		if v != nil {
			return v
		}
	}
	return nil
}

func Contains(list []string, obj string) bool {
	for index := range list {
		if list[index] == obj {
			return true
		}
	}
	return false
}

func ContainsInt(list []int, obj int) bool {
	for index := range list {
		if list[index] == obj {
			return true
		}
	}
	return false
}

func ContainsUint(list []uint, obj uint) bool {
	for index := range list {
		if list[index] == obj {
			return true
		}
	}
	return false
}

func GetAttr(obj interface{}, fieldName string) (reflect.Value, bool) {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct

	// Check if the passed interface is a pointer
	if pointToStruct.Type().Kind() == reflect.Ptr {
		curStruct = pointToStruct.Elem()
	}

	if curStruct.Kind() != reflect.Struct {
		return reflect.ValueOf(nil), false
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		return reflect.ValueOf(nil), false
	}
	return curField, true
}

func GetFieldValue(v interface{}, field string) any {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

// for limiting concurrency work (fetching document data from library koha)
func Do(concurrency int, fn func(int), listJob []int) {
	workQueue := make(chan int, concurrency)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for job := range workQueue {
				fn(job)
			}
			wg.Done()
		}()
	}
	for _, i := range listJob {
		workQueue <- i
	}
	close(workQueue)
	wg.Wait()
}

//func ConnectTLS() (*ldap.Conn, error) {
//	// You can also use IP instead of FQDN
//	l, err := ldap.DialURL("ldap://10.20.2.253:389")
//	if err != nil {
//		log.Println(err)
//		l, err = ldap.DialURL("ldap://10.20.2.222:51276")
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	if err := l.Bind(model.BindDN, model.LDAPPassword); err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	return l, nil
//}

//func BindAndSearch(l *ldap.Conn, username string) (*ldap.SearchResult, error) {
//	searchReq := ldap.NewSearchRequest(
//		model.BaseDN,
//		2, // you can also use ldap.ScopeWholeSubtree
//		ldap.NeverDerefAliases,
//		0,
//		0,
//		false,
//		model.LDAPUserFilter+"(userPrincipalName="+username+"))",
//		[]string{},
//		nil,
//	)
//	result, err := l.Search(searchReq)
//	if err != nil {
//		return nil, fmt.Errorf("Search Error: %s", err)
//	}
//	if len(result.Entries) > 0 {
//		return result, nil
//	} else {
//		return nil, fmt.Errorf("Couldn't fetch search entries")
//	}
//}

func Paginate(x []interface{}, pageNum int, pageSize int) []interface{} {
	if pageSize == -1 {
		return x
	}

	sliceLength := len(x)
	var start = 0
	if pageNum != 1 {
		start = (pageNum - 1) * pageSize
	}

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return x[start:end]
}
func CustomPaginateTotal(datas []interface{}, dbTotal int) int {
	if dbTotal != -1 {
		return dbTotal
	}

	countMap := make(map[uint64]int)
	for _, data := range datas {
		field, hasField := GetAttr(data, "ID")
		if hasField && field.Uint() != 0 {
			countMap[field.Uint()] = 1
		}
	}
	return len(countMap)
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
func IsListStringContains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func DuplicateFieldValue(values interface{}, fieldName string) bool {
	refVal := reflect.ValueOf(values)
	if refVal.Kind() != reflect.Slice {
		panic("not a slice")
	}

	mapVal := make(map[interface{}]int)
	for i := 0; i < refVal.Len(); i++ {
		fieldValue := refVal.Index(i).FieldByName(fieldName).Interface()
		if fieldValue != nil {
			mapVal[fieldValue]++
		} else {
			continue
		}
	}

	for _, v := range mapVal {
		if v >= 2 {
			return true
		} else {
			return false
		}
	}

	return true
}

// RemoveDuplicateInterface might be inconsistent
func RemoveDuplicateInterface(data interface{}, fieldName string) []interface{} {
	mapValue := make(map[interface{}]interface{})
	reflectData := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() != reflect.Slice {
		// log.Println(reflect.TypeOf(data))
		panic("value kind is not slice")
	}

	for i := 0; i < reflectData.Len(); i++ {
		d := reflectData.Index(i)
		fieldValue := d.FieldByName(fieldName).Interface()
		mapValue[fieldValue] = d.Interface()
	}
	res := make([]interface{}, 0)
	for _, v := range mapValue {
		res = append(res, v)
	}

	return res
}

// FillInterfaceArrayToArray fills data of interface array type into a concrete type array.
// return interface{} (underlying: []modelType)
func FillInterfaceArrayToArray(data []interface{}, modelType string) interface{} {
	slice := reflect.MakeSlice(reflect.TypeOf(model.LIST_MODEL_MAP[modelType]), 0, len(data))
	for _, d := range data {
		slice = reflect.Append(slice, reflect.ValueOf(d))
	}

	return slice.Interface()
}

func GetValueOf(obj any, field string) any {
	value := reflect.ValueOf(obj)
	val := reflect.Indirect(value).FieldByName(field)
	if val.IsValid() {
		return val
	}
	return nil
}

func ArrayIntToString(arr []int, _sort bool) []string {
	res := []string{}

	if _sort == true {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}

	for _, item := range arr {
		s := strconv.Itoa(item)
		res = append(res, s)
	}
	return res
}

func SortArrayString(arr []string) []string {
	tmp := arr
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	return tmp
}

func ParsePtr2Bool(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return bool(*ptr)
}

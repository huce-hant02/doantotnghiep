package repository

import (
	"doantotnghiep/infrastructure"
	"doantotnghiep/model"
	"strconv"
)

type authorRepository struct {
}

func (a authorRepository) FilterPolicy(authorApiID, authorID, ApiID int, Role, URL, Method, Value string) ([]model.Policy, error) {
	db := infrastructure.GetDB()

	var whereQuery string
	if Role == "" {
		whereQuery += " authors.role ILIKE '%'"
	} else {
		whereQuery += " authors.role ILIKE " + Role
	}

	if URL != "" {
		whereQuery += " AND apis.URL LIKE " + URL
	}
	if Method != "" {
		whereQuery += " AND apis.Method LIKE " + Method
	}
	if Value != "" {
		whereQuery += " AND author_apis.value LIKE" + Value
	}

	if authorApiID > 0 {
		whereQuery += " AND author_apis.id = " + strconv.Itoa(authorApiID)
	}
	if authorID > 0 {
		whereQuery += " AND authors.id = " + strconv.Itoa(authorID)
	}
	if ApiID > 0 {
		whereQuery += " AND apis.id = " + strconv.Itoa(ApiID)
	}

	result := []model.Policy{}
	err := db.Table("author_apis").Select("author_apis.author_id AS  author_id, author_apis.api_id AS api_id, " +
		"authors.role AS role, apis.url AS url, apis.method AS method, author_apis.value AS value").
		Joins(" JOIN authors ON authors.id = author_apis.author_id ").
		Joins(" JOIN apis ON apis.id = author_apis.api_id").
		Where(whereQuery).Scan(&result).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, err
	}
	return result, nil
}

func (a authorRepository) GetPolicyById(authorApiID int) (*model.Policy, error) {
	db := infrastructure.GetDB()
	result := model.Policy{}
	err := db.Table("author_apis").Select("author_apis.Author_id AS api_id, author_apis.Api_id AS api_id, "+
		"authors.Role AS role, apis.url AS url, apis.method AS method, author_apis.value AS value").
		Joins(" JOIN authors ON authors.id = author_apis.author_id ").
		Joins(" JOIN apis ON apis.id = author_apis.api_id").
		Where("author_apis.id = ?", authorApiID).Scan(&result).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, err
	}
	return &result, nil
}

func (a authorRepository) AddNewApi(URL, Method string) error {
	db := infrastructure.GetDB()
	err := db.Create(&model.Api{
		URL:    URL,
		Method: Method,
	}).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func (a authorRepository) GetAllApi(page int, pageSize int) (apis []model.Api, total int64, err error) {
	db := infrastructure.GetDB()
	var temp []model.Api
	err = db.Model(&model.Api{}).
		Count(&total).
		Scan(&temp).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, 0, err
	}

	err = db.Limit(pageSize).
		Offset((page - 1) * pageSize).
		Model(&model.Api{}).
		Select("*").Order("url, method asc ").
		Scan(&apis).Error

	if err != nil {
		return nil, 0, err
	}

	return apis, total, nil
}

func (a authorRepository) DeleteApi(idAPi int) error {
	db := infrastructure.GetDB()
	err := db.Where("id =  ?", idAPi).Delete(model.Api{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a authorRepository) UpdateApi(id int, newAPi model.Api) error {
	db := infrastructure.GetDB()

	err := db.Model(&model.Api{}).Where("id = ?", id).Updates(newAPi).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func (a authorRepository) AddNewAuthor(roleName, desc string) error {
	db := infrastructure.GetDB()
	err := db.Create(&model.Author{
		Role:        roleName,
		Description: desc,
	}).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func (a authorRepository) GetAllAuthor(page int, pageSize int) (authors []model.Author, total int64, err error) {
	db := infrastructure.GetDB()
	var temp []model.Author
	err = db.Model(&model.Author{}).
		Count(&total).
		Scan(&temp).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, 0, err
	}

	err = db.Limit(pageSize).
		Offset((page - 1) * pageSize).
		Model(&model.Author{}).
		Select("*").Order("role ASC ").
		Scan(&authors).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, 0, err
	}

	return authors, total, nil
}

func (a authorRepository) DeleteAuthor(id int) error {
	db := infrastructure.GetDB()
	err := db.Where("id =  ?", id).Delete(model.Author{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a authorRepository) UpdateAuthor(id int, newAuthor model.Author) error {
	db := infrastructure.GetDB()

	err := db.Model(&model.Author{}).Where("id = ?", id).Updates(newAuthor).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func (a authorRepository) AddAuthor_api(newPolicy model.Policy) error {
	db := infrastructure.GetDB()
	err := db.Create(&model.Author_api{
		AuthorId: newPolicy.AuthorId,
		ApiId:    newPolicy.ApiId,
		Value:    newPolicy.Value,
	}).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err

	}
	return nil
}

func (a authorRepository) GetAuthor_api(page int, pageSize int) (au_apis []model.Author_api, total int64, err error) {
	db := infrastructure.GetDB()
	var temp []model.Author
	err = db.Model(&model.Author_api{}).
		Count(&total).
		Scan(&temp).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, 0, err
	}

	err = db.Limit(pageSize).
		Offset((page - 1) * pageSize).
		Model(&model.Author_api{}).
		Select("*").
		Scan(&au_apis).Error

	if err != nil {
		infrastructure.ErrLog.Println(err)
		return nil, 0, err
	}

	return au_apis, total, nil
}

func (a authorRepository) DeleteAuthor_api(id int) error {
	db := infrastructure.GetDB()
	err := db.Where("id =  ?", id).Delete(model.Author_api{}).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err

	}
	return nil
}

func (a authorRepository) DeleteAuthor_api_byauthorandapi(authorId int, apiId int) error {
	db := infrastructure.GetDB()
	err := db.Where("author_id =  ? AND api_id = ?", authorId, apiId).Delete(model.Author_api{}).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err

	}
	return nil
}

func (a authorRepository) UpdateAuthor_api(id int, newAuthor_api model.Author_api) error {
	db := infrastructure.GetDB()

	err := db.Model(&model.Author_api{}).Where("id = ?", id).Updates(newAuthor_api).Error
	if err != nil {
		infrastructure.ErrLog.Println(err)
		return err
	}
	return nil
}

func NewAuthorRepository() model.AuthorRepository {

	return authorRepository{}
}

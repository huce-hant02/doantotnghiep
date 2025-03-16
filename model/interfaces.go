package model

import "gorm.io/gorm"

// repository responsible for quering in db, filterModel find in /repository/mapData.go
type ISearchRepo interface {
	AdvanceFilter(filterModel interface{}, pagination Pagination, ignoreAssociation, includeSoftDelete bool, sort []SortItem) (_ interface{}, count int, err error)
}

// repository responsible for manipulating association relation in database, associationName find in /repository/mapData.go
type IAssociationOpRepo interface {
	AddAssociation(db *gorm.DB, associationName string, model interface{}, association interface{}) (associationDatas interface{}, err error)
	RemoveAssociation(db *gorm.DB, associationName string, model interface{}, association interface{}) (associationDatas interface{}, err error)
	ClearAssociation(db *gorm.DB, associationName string, model interface{}) error
	ReplaceAssociation(db *gorm.DB, associationName string, mainModel interface{}, datas interface{}) (associationDatas interface{}, err error)
}

// repository responsible for standard db querying like. insert, update, delete
type IBasicDBQueryRepo interface {
	Add(db *gorm.DB, datas interface{}) (newDatas interface{}, err error)
	Update(db *gorm.DB, data interface{}) (updatedData interface{}, err error)
	Delete(db *gorm.DB, ID uint) (err error)
}

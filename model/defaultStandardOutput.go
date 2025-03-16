package model

import (
	"time"

	"gorm.io/gorm"
)

type DefaultStandardOutput struct {
	ID                   uint   `json:"id"`
	ScholasticSemesterID uint   `json:"scholasticSemesterId" gorm:"index:default_standard_output_indexing_scholastic_semester_unique,unique"`
	Indexing             string `json:"indexing" gorm:"index:default_standard_output_indexing_scholastic_semester_unique,unique"`

	OutputLevel int    `json:"outputLevel"` // education standard output level
	Description string `json:"description"`

	KeywordID uint `json:"keywordId"`
	LevelID   uint `json:"levelId"`
	AreaID    uint `json:"areaId"`

	Keyword   *Keyword   `json:"keyword" swaggerignore:"true" gorm:"foreignKey:KeywordID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`
	BloomWord *BloomWord `json:"area" swaggerignore:"true" gorm:"foreignKey:AreaID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`
	Level     *Level     `json:"level" swaggerignore:"true" gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`

	// foreignKey
	DefaultITUs []DefaultITU `swaggerignore:"true" json:"-" gorm:"foreignKey:DefaultStandardOutputID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updatedAt" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

type DefaultStandardOutputRepository interface {
	ISearchRepo
}

package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (t Category) TableName() string {
	return "category"
}

type Category struct {
	CategoryId int `gorm:"primary_key;AUTO_INCREMENT"`
	CategoryName string `gorm:"type:varchar(50);not null"`
	CategoryDesc string `gorm:"type:text;"`
	CreatedOn time.Time
	CreatedBy string `gorm:"type:int;"`
	AppId int32 `gorm:"type:varchar(100);"`
	Status int32 `gorm:"type:int;"`
}

func(Category) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Category{},).Error
	if err != nil {
		return err
	}

	return nil
}

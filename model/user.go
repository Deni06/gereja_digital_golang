package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (t User) TableName() string {
	return "user"
}

type User struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"type:varchar(50);not null"`
	CreatedOn time.Time
	CreatedBy string `gorm:"type:int;"`
	LastUpdate time.Time
	Password string `gorm:"type:varchar(50);not null"`
	AppId int32 `gorm:"type:int;"`
	Name string `gorm:"type:varchar(50);not null"`
	Phone string `gorm:"type:varchar(30);not null"`
	Email string `gorm:"type:varchar(50);not null"`
	Position string `gorm:"type:varchar(100);not null"`
	StatusRegistrasi int32 `gorm:"type:int;"`
}

func(User) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&User{},).Error
	if err != nil {
		return err
	}

	return nil
}

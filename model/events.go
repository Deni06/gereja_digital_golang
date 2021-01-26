package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (t Events) TableName() string {
	return "events"
}

type Events struct {
	EventId uint `gorm:"primary_key;AUTO_INCREMENT"`
	EventName string `gorm:"type:varchar(100);not null"`
	EventStart time.Time `gorm:"type:time;"`
	EventDetail string `gorm:"type:text;"`
	EventLocation string `gorm:"type:text;"`
	CreatedDate time.Time
	CreatedBy string `gorm:"type:varchar(50);"`
	EventDate time.Time `gorm:"type:date;"`
	EventEnd time.Time `gorm:"type:time;"`
	EventImg string `gorm:"type:varchar(255);"`
	AppId int32 `gorm:"type:int;"`
	Status int32 `gorm:"type:int;"`
}

func(Events) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Events{},).Error
	if err != nil {
		return err
	}

	return nil
}

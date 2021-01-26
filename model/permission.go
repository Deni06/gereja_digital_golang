package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (t Permission) TableName() string {
	return "permission"
}

type Permission struct {
	PermissionId uint `gorm:"primary_key;AUTO_INCREMENT"`
	PermissionName string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
}

func(Permission) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Permission{},).Error
	if err != nil {
		return err
	}
	return nil
}

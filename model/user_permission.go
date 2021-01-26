package model

import (
	"github.com/jinzhu/gorm"
)

func (t UserPermission) TableName() string {
	return "user_permission"
}

type UserPermission struct {
	PermissionId int `gorm:"primary_key;type:int;not null;default:0"`
	UserId int `gorm:"primary_key;type:int;not null;default:0"`
}

func(UserPermission) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&UserPermission{},).Error
	if err != nil {
		return err
	}
	exist := tx.Dialect().HasIndex("user_permission","userPermission")
	if !exist{
		err = tx.Model(&UserPermission{},).AddUniqueIndex("userPermission", "permission_id", "user_id").Error
		if err != nil {
			return err
		}
	}

	return nil
}

package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func (t App) TableName() string {
	return "app"
}

type App struct {
	AppId int `gorm:"primary_key;AUTO_INCREMENT"`
	AppName string `gorm:"type:varchar(50);"`
	PackageName string `gorm:"type:varchar(100);"`
	SubscriptionType string `gorm:"type:varchar(50);"`
	SubscriptionStart time.Time
	SubscriptionEnd time.Time
	CreatedBy string `gorm:"type:int;"`
	CreatedOn time.Time
	LastUpdate time.Time
	Icon string `gorm:"type:varchar(100);"`
	StyleLayout string `gorm:"type:varchar(50);"`
	ColorCodeTabBar string `gorm:"type:varchar(50);"`
	ColorCodeHeader string `gorm:"type:varchar(50);"`
	ColorCodeBackground string `gorm:"type:varchar(50);"`
	MetodePembayaran string `gorm:"type:varchar(50);"`
	IsTrial int32 `gorm:"type:int;"`
	TrialStartDate time.Time
	TrialEndDate time.Time
	Voucher string `gorm:"type:varchar(100);"`
	CurrentVersion string `gorm:"type:varchar(10);"`
	StatusDesign int32 `gorm:"type:int;"`
	LogoSplashscreen string `gorm:"type:varchar(100);"`
}

func(App) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&App{},).Error
	if err != nil {
		return err
	}

	return nil
}

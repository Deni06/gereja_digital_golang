package driver

import (
	"github.com/jinzhu/gorm"
)

type GetConnectionGeneric interface {
	newDBGorm() (*gorm.DB, error)
}

type CustomQueryInterface interface {
	GetInsertQuery(db *gorm.DB, query string, primary_key string, table_name string) *gorm.DB
	GetQuery(*string)
}

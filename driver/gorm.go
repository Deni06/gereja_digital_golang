package driver

import (
	"gintest/constant"
	"gintest/util"
	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbInstanceGorm = gorm.DB{}
var parameter = Parameter{}

type Parameter struct {
	UseCli bool
	Host string
	HostParam string
	User string
	Password string
	SslMode string
	DbName string
	Dialect string
	DbPath string
}

func InitGorm() (*gorm.DB, error){
	db, err := newDBGorm()
	if err != nil {
		return nil, err
	}
	return db,nil
}

func SetParamGorm(input Parameter){
	parameter = input
}

func newDBGorm() (*gorm.DB, error) {
	if parameter.Dialect == "" {
		parameter.Dialect = constant.POSTGRESQL_DIALECT
	}
	connectionHandler, err := GetConnectionHandler(parameter.Dialect)
	if err != nil {
		return nil, err
	}
	db, err := connectionHandler.newDBGorm()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBInstanceGorm()(gorm.DB,error){
	if (dbInstanceGorm .DB == nil){
		return dbInstanceGorm,util.UnhandledError{ErrorMessage:constant.UNHANDLED_ERROR}
	}else{
		return dbInstanceGorm, nil
	}
}
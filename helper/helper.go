package helper

import (
	"gintest/driver"
	"github.com/jinzhu/gorm"
	"strings"
)


var currentDriver driver.GetConnectionGeneric
var currentDialect string
var currentImplementation driver.CustomQueryInterface
var currentDBInstance gorm.DB

func InitDriver(dialect string) error{
	driver, err := driver.GetConnectionHandler(dialect)
	if err!=nil{
		return err
	}
	currentDriver = driver
	currentDialect = dialect
	currentImplementation, err = SetCurrentImplementation()
	if err!=nil{
		return  err
	}
	currentDBInstance, err = SetDBInstance()
	if err!=nil{
		return  err
	}
	return nil
}

func GetCurrentDriver()driver.GetConnectionGeneric{
	return currentDriver
}

func SetCurrentImplementation()(driver.CustomQueryInterface, error){
	currentImplementation, err := driver.GetCustomQuery(currentDialect)
	if err!=nil{
		return nil, err
	}
	return currentImplementation, nil
}

func GetCurrentImplementation()driver.CustomQueryInterface{
	return currentImplementation
}

func SetDBInstance()(gorm.DB, error){
	currentDBInstance, err :=  driver.GetDBInstanceGorm()
	if err !=nil{
		return currentDBInstance, err
	}
	return currentDBInstance, nil
}

func GetDBInstance()gorm.DB{
	return currentDBInstance
}

func CheckSingleQuoteLoyalty(in *string) string{
	title := ""
	if strings.Contains(*in, "'"){
		for _, data := range strings.Split(*in, ""){
			if data =="'"{
				title += "'"
				title += data
			}else{
				title += data
			}
		}
		*in = title
	}
	return *in
}
package driver

import (
	"fmt"
	"gintest/util"
	"github.com/jinzhu/gorm"
	"strings"
)

func (p PostgreeImplementation)GetInsertQuery(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	query = query+fmt.Sprintf(" Returning %v ",primaryKey)
	newQuery := db.Raw(query)
	return newQuery
}

func (m MysqlImplementation)GetInsertQuery(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	newQuery := db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Exec(query)
	}, func(db *gorm.DB) *gorm.DB {
		return db.Select("LAST_INSERT_ID()").Table(tableName)
	})
	return newQuery
}

func (s SqliteImplementation)GetInsertQuery(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	newQuery := db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Exec(query)
	}, func(db *gorm.DB) *gorm.DB {
		return db.Select("last_insert_rowid()").Table(tableName)
	})
	return newQuery
}

func (ss SqlServerImplementation)GetInsertQuery(db *gorm.DB, query string, primaryKey string, tableName string) *gorm.DB{
	query = query+" select ID = convert(bigint, SCOPE_IDENTITY()) "
	newQuery := db.Raw(query)
	return newQuery
}

func (p PostgreeImplementation)GetQuery(in *string){

}

func (m MysqlImplementation)GetQuery(in *string){
}

func (s SqliteImplementation)GetQuery(in *string){
	val := strings.Replace(*in,"EXTRACT(DAY from", "strftime('%d', ",-1)
	val = strings.Replace(*in,"EXTRACT(MONTH from", "strftime('%m', ",-1)
	val = strings.Replace(*in,"EXTRACT(WEEK from", "strftime('%W', ",-1)
	val = strings.Replace(*in,"EXTRACT(YEAR from", "strftime('%Y', ",-1)
	val = strings.Replace(val, "now()","date('now')",-1)
	*in = val
}

func (ss SqlServerImplementation)GetQuery(in *string){
	val := strings.Replace(*in,"false", "0",-1)
	val = strings.Replace(val,"true", "1",-1)
	val = strings.Replace(val, "now()","GETDATE()",-1)
	val = strings.Replace(val, "EXTRACT(DAY from","DATEPART(DAY, ",-1)
	val = strings.Replace(val, "EXTRACT(WEEK from","DATEPART(WEEK, ",-1)
	val = strings.Replace(val, "EXTRACT(MONTH from","DATEPART(MONTH, ",-1)
	val = strings.Replace(val, "EXTRACT(YEAR from","DATEPART(YEAR, ",-1)
	*in = val
}

func (p PostgreeImplementation)HandlingError(in error) error{
	if strings.Contains(in.Error(), "unique constraint \\\"order_number_unique\\\""){
		in = util.UnhandledError{ErrorMessage:"order number already exist"}
	}
	return in
}

func (m MysqlImplementation)HandlingError(in error) error{
	return in
}

func (s SqliteImplementation)HandlingError(in error) error{
	return in
}

func (ss SqlServerImplementation)HandlingError(in error) error{
	if strings.Contains(in.Error(), "unique index 'order_number_unique'"){
		in = util.UnhandledError{ErrorMessage:"order number already exist"}
	}
	return in
}
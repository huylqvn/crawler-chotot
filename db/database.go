package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupDB(dbHost, dbUser, dbPassword, dbName, dbPort string) *gorm.DB {
	//connect to db
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	db.SingularTable(true)
	fmt.Println("Database Connected", dbHost, dbUser)
	return db
}

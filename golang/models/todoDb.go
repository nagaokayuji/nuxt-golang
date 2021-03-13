package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	HOST := "db"
	PORT := "3306"
	DBNAME := "golang-nuxt-app"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	fmt.Println("connecting URL = ", URL)

	db, err := gorm.Open(DBMS, URL)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	fmt.Println("db connected: ", &db)
	return db

}

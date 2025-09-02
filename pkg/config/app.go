package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db_credentials := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	d, err := gorm.Open("mysql", db_credentials)
	if err != nil {
		panic(err)
	}
	log.Println("Connection Opened to Database")
	db = d
}

func GetDB() *gorm.DB {
	return db
}

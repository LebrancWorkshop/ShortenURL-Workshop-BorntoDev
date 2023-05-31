package db

import (
	"log"

	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/forms"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *gorm.DB

func OpenDB(dbUsername string, dbPassword string, dbName string, dbHost string, dbPort string) (*gorm.DB) {
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&forms.ShortlyURL{})

	return db 
}

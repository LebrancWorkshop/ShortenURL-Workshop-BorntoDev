package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/controllers"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/db"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/forms"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db := db.OpenDB(dbUsername, dbPassword, dbName, dbHost, dbPort)

	model := models.NewModel(db)
	controllers.NewController(model)

	router.Run(":8101")
}


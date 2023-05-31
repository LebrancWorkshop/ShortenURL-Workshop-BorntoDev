package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/db"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/forms"
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

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/api/v1/url", func(c *gin.Context) {
		type RequestData struct {
			URL string `json:"url" binding: "required"`
		}
		var data RequestData

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var newLink forms.ShortlyURL
		result := db.Where("original_url = ?", data.URL).First(&newLink)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				shortURL := generateShortURL()
				newLink = forms.ShortlyURL{
					OriginalURL: data.URL,
					ShortURL: shortURL,
				}
				result = db.Create(&newLink)
				if result.Error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": result.Error.Error(),
					})
					return
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"short_url": newLink.ShortURL,
		})
	})

	router.GET("/api/v1/url", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		var link forms.ShortlyURL
		result := db.Where("short_url = ?", shortURL).Find(&link)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "URL not found",
				})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": result.Error.Error(),
					})
			}
			return
		}
		c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
	})

	router.Run(":8101")
}

func generateShortURL() string {
	const alphabetUpperCase = "abcdefghijklmnopqrstuvwxyz"
	const alphabetLowerCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const number = "0123456789"
	const charSet = alphabetLowerCase + alphabetUpperCase + number

	rand.Seed(time.Now().UnixNano())

	var shortURL string
	for i := 0; i < 6; i++ {
		shortURL += string(charSet[rand.Intn(len(charSet))])
	}

	return shortURL
}

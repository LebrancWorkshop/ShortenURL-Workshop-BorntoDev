package controllers

import (
	"net/http"

	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/forms"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/models"
	"github.com/LebrancWorkshop/ShortenURL-Workshop-BorntoDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Model *models.Model
}

func NewController(model *models.Model) *Controller {
	return &Controller{Model: model}
}

func (url *Controller) Ping(c *gin.Context) {

}

func (c *Controller) CreateShortURL(ctx *gin.Context) {
	type RequestData struct {
			URL string `json:"url" binding: "required"`
		}
		var data RequestData

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var link forms.ShortlyURL
		result, err := c.Model.GetOriginalURL(&link)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				shortURL := utils.GenerateShortURL()
				link = forms.ShortlyURL{
					OriginalURL: data.URL,
					ShortURL: shortURL,
				}
				result, err = c.Model.CreateShortURL(&link)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": result.Error.Error(),
					})
					return
				}
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"short_url": link.ShortURL,
		})
}

func (c *Controller) GetShortURL(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")
	var link forms.ShortlyURL
	result, err := c.Model.GetShortURL(&link, shortURL)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "URL not found",
			})
		} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": result.Error.Error(),
				})
		}
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, link.OriginalURL)
}

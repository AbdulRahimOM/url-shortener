package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/AbdulRahimOM/misc-projects/url-shortener/config"
	"github.com/AbdulRahimOM/misc-projects/url-shortener/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	alreadyRegisteredMsg = "This URL is already registered"
	successfulMsg = "URL shortened successfully"
	urlNotFoundMsg = "URL not found"
)

func main() {
	e := echo.New()
	e.GET("/health-check", ping)
	e.POST("/generate", generateUrl)
	e.GET("/:shortUrl", getUrl)

	e.Logger.Fatal(e.Start(":" + config.Project.Port))
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}

func generateUrl(c echo.Context) error {
	longUrl := c.FormValue("longUrl")
	if longUrl == "" {
		return c.JSON(http.StatusBadRequest, domain.ErrorRes{
			Status:  false,
			Message: "Invalid URL",
			Error:   "URL cannot be empty",
		})
	}

	if err := isValidURL(longUrl); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorRes{
			Status:  false,
			Message: "Invalid URL",
			Error:   err.Error(),
		})
	}

	if strings.ContainsAny(longUrl, " ") {
		return c.JSON(http.StatusBadRequest, domain.ErrorRes{
			Status:  false,
			Message: "Invalid URL",
			Error:   "URL cannot contain spaces",
		})
	}

	fmt.Println("URL: ", longUrl)
	//check if the url already registered once
	var urlRecord domain.UrlRecord
	result := config.DB.Where("long_url = ?", longUrl).First(&urlRecord)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, domain.ErrorRes{
				Status:  false,
				Message: "Database error.",
				Error:   result.Error.Error(),
			})
		}
	} else {
		return c.JSON(http.StatusOK, domain.SuccessRes{
			Status:       true,
			Message:      alreadyRegisteredMsg,
			ShortenedUrl: config.Project.Host + ":" + config.Project.Port + "/" + urlRecord.ShortUrlPath,
		})
	}

	path := uuid.New().String()[:config.Url.Length]
	result = config.DB.Where("short_url_path = ?", path).First(&urlRecord)

	for result.Error != gorm.ErrRecordNotFound {
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, domain.ErrorRes{
				Status:  false,
				Message: "Database error",
				Error:   result.Error.Error(),
			})
		}
		path = uuid.New().String()[:config.Url.Length]
		result = config.DB.Where("short_url_path = ?", path).First(&urlRecord)
	}

	result = config.DB.Create(&domain.UrlRecord{
		LongUrl:      longUrl,
		ShortUrlPath: path,
	})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorRes{
			Status:  false,
			Message: "Database error",
			Error:   result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.SuccessRes{
		Status:       true,
		Message:      successfulMsg,
		ShortenedUrl: config.Project.Host + ":" + config.Project.Port + "/" + path,
	})
}
func isValidURL(u string) error {
	_, err := url.ParseRequestURI(u)
	return err
}
func getUrl(c echo.Context) error {
	path := c.Param("shortUrl")
	var urlRecord domain.UrlRecord
	result := config.DB.Where("short_url_path = ?", path).First(&urlRecord)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, domain.ErrorRes{
			Status:  false,
			Message: urlNotFoundMsg,
			Error:   result.Error.Error(),
		})
	}

	return c.Redirect(http.StatusMovedPermanently, urlRecord.LongUrl)
}

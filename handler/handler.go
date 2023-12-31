package handler

import (
	"UrlShortner/store"
	"UrlShortner/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

type CreateUrlRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
	UserId      string `json:"userId" binding:"required"`
}

type CreateCustomUrlRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
	CustomUrl   string `json:"customUrl" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var request CreateUrlRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		log.Println(request)
		return
	}

	shortUrl := utils.GenerateShortUrl(request.OriginalUrl, request.UserId)
	store.SaveUrlMapping(shortUrl, request.OriginalUrl)

	host := fmt.Sprintf("http://%v/", ctx.Request.Host)

	ctx.JSON(200, gin.H{
		"message":  "Short url successfully generated",
		"shortUrl": host + shortUrl,
	})
}

func CreateCustomShortUrl(ctx *gin.Context) {
	var request CreateCustomUrlRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	re := regexp.MustCompile("[^A-Za-z0-9-+]")
	customUrl := re.ReplaceAllString(request.CustomUrl, "")

	res := store.RetrieveInitialUrl(customUrl)
	if len(res) != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Link already in used", "message": "Try different link"})
		return
	}

	store.SaveUrlMapping(customUrl, request.OriginalUrl)

	host := fmt.Sprintf("http://%v/", ctx.Request.Host)

	ctx.JSON(200, gin.H{
		"message":  "Custom short url successfully generated",
		"shortUrl": host + customUrl,
	})
}

func RedirectToOriginalUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	originalUrl := store.RetrieveInitialUrl(shortUrl)

	if len(originalUrl) == 0 {
		ctx.JSON(404, gin.H{"error": "Failed to find the associated url", "message": "Short url doesn't exists"})
	}

	if len(originalUrl) < 7 ||
		originalUrl[:7] != "http://" ||
		originalUrl[:8] != "https://" {
		ctx.Redirect(302, "http://"+originalUrl)
	} else {
		ctx.Redirect(302, originalUrl)
	}
}

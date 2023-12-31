package main

import (
	"UrlShortner/handler"
	"UrlShortner/store"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	store.InitializeStore()
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "Ok",
		})
	})

	r.POST("create/short/url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.POST("create/custom/short/url", func(ctx *gin.Context) {
		handler.CreateCustomShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.RedirectToOriginalUrl(ctx)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"usd/api"
	"usd/internal/module"
	"usd/internal/repository"
)

func main() {
	URLRepo := repository.NewURLRepository()
	shortenerModule := module.NewShortener(URLRepo)

	port := ":8101"
	httpHandler := api.NewHandler(shortenerModule, "localhost"+port)

	router := gin.Default()
	router.POST("/shortener", httpHandler.Shorten)
	router.GET("/r/:hash", httpHandler.Resolve)
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}

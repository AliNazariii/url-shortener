package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"usd/internal/module"
)

type HandlerImpl struct {
	shortenerModule module.Shortener
	hostname        string
}

func NewHandler(shortenerModule module.Shortener, hostname string) *HandlerImpl {
	return &HandlerImpl{
		shortenerModule: shortenerModule,
		hostname:        hostname,
	}
}

type ShortenerRequest struct {
	URL string `json:"url"`
}

func (h *HandlerImpl) Shorten(context *gin.Context) {
	var request ShortenerRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if request.URL == "" {
		err := errors.New("url is empty")
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	shortURL, err := h.shortenerModule.Shorten(request.URL)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"short_url": h.hostname + "/r/" + shortURL,
	})
}

func (h *HandlerImpl) Resolve(context *gin.Context) {
	hash := context.Param("hash")
	if hash == "" {
		err := errors.New("url is empty")
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resolve, err := h.shortenerModule.Resolve(hash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.Redirect(http.StatusFound, resolve)
}

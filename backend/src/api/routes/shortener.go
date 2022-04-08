package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	services "dcardHw/src/services"
)

type ShortenerReq struct {
	Url      string    `json:"url" form:"url" binding:"required"`
	ExpireAt time.Time `json:"expireAt" form:"expireAt" binding:"required" time_format:"2021-02-08T09:20:41Z"`
}

func Shortener(server *gin.Engine) *gin.Engine {
	server.POST("/api/v1/urls", func(ctx *gin.Context) {
		var shortenerReq ShortenerReq
		err := ctx.ShouldBind(&shortenerReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		status, shortenUrl := services.GenerateShortenUrl(shortenerReq.Url, shortenerReq.ExpireAt)
		if status == 0 {
			ctx.JSON(http.StatusAccepted, gin.H{"id": shortenerReq.Url, "shortUrl": shortenUrl})
		} else if status == 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url"})
		} else if status == 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expireAt"})
		}

	})
	return server
}

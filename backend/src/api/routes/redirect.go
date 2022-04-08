package routes

import (
	"dcardHw/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(server *gin.Engine) *gin.Engine {
	server.GET("/:short", func(ctx *gin.Context) {
		shorturl := ctx.Params.ByName("short")
		fmt.Println(shorturl)
		status, ori := services.RedirectUrl(shorturl)
		if status == 0 {
			ctx.Redirect(http.StatusFound, ori)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid url"})
		}
	})
	return server
}

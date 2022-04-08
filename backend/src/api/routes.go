package api

import (
	routes "dcardHw/src/api/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() gin.Engine {
	server := gin.Default()
	routes.Shortener(server)
	routes.Redirect(server)
	return *server
}

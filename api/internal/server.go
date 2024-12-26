package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/gonza56d/gauth/internal/routes"
)

func RunServer() {
	var server *gin.Engine = gin.Default()
	v1 := server.Group("/v1")
	routes.AddRoutes(v1)
	server.Run()
}

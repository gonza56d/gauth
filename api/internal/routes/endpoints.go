package routes

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	//"github.com/gonza56d/gauth/internal/app"
	//apimodel "github.com/gonza56d/gauth/pkg"
	"github.com/google/uuid"
)

func AddRoutes(rg *gin.RouterGroup) {
	helloRoute(rg)
	authRoute(rg)
}

func helloRoute(rg *gin.RouterGroup) {
	helloRoute := rg.Group("/hello")

	helloRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})
}


func isValidUUID(value string, ctx *gin.Context, action string) *uuid.UUID {
	result, err := uuid.Parse(value)
	if err != nil || result.Version() != 4 {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
			"error": "Invalid UUID for auth ID",
			"action": action,
		})
		return nil
	}
	return &result 
}

func authRoute(rg *gin.RouterGroup) {
	//authRoute := rg.Group("/auth")
}

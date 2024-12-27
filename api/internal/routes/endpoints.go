package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gonza56d/gauth/internal/app"
	apimodel "github.com/gonza56d/gauth/pkg"
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
	authRoute := rg.Group("/auth")

	// authenticate
	authRoute.POST("/login", func(ctx *gin.Context) {
		action := "LOGIN"
		var request apimodel.LoginRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Bad request",
				"error": err.Error(),
				"action": action,
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Bad request",
				"error": err.Error(),
				"action": action,
			})
			return
		}
		jwt_token := app.Login(&request)
		if jwt_token == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
				"error": "Invalid email or password",
				"action": action,
			})
			return
		}
		ctx.JSON(201, gin.H{
			"jwt_token": jwt_token,
		})
	})
}

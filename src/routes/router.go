package routes

import (
	"backend-go/src/controllers"
	"backend-go/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Rutas publicas
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	//Rutas protegidas
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/saludo", controllers.HolaMundo)

	}

	return r
}

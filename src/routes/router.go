package routes

import (
	"fmt"

	"backend-go/src/controllers"
	"backend-go/src/middleware"
	"backend-go/src/utils"

	"github.com/gin-gonic/gin"
)

// AddRoute agrega la ruta y loguea automáticamente
func AddRoute(group *gin.RouterGroup, method string, path string, handler gin.HandlerFunc, protected bool) {
	fullPath := group.BasePath() + path
	label := "→"
	if protected {
		label = "🔐"
	}

	switch method {
	case "GET":
		group.GET(path, handler)
	case "POST":
		group.POST(path, handler)
	case "PUT":
		group.PUT(path, handler)
	case "DELETE":
		group.DELETE(path, handler)
	// agrega más métodos si lo necesitas
	}

	utils.LogInfo(fmt.Sprintf("%s [%s] %s", label, method, fullPath))
}

// SetupRouter inicializa todas las rutas
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rutas públicas
	public := r.Group("/auth")
	AddRoute(public, "POST", "/login", controllers.Login, false)
	AddRoute(public, "POST", "/register", controllers.Register, false)

	// Rutas protegidas
	privateAuth := r.Group("/auth", middleware.AuthMiddleware())
	AddRoute(privateAuth, "GET", "/user", controllers.GetMe, true)

	privateAPI := r.Group("/api", middleware.AuthMiddleware())
	AddRoute(privateAPI, "GET", "/saludo", controllers.HolaMundo, true)

	return r
}

package main

import (
	"backend-go/src/config"
	"backend-go/src/routes"
	"backend-go/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.LoadEnv()

	if err := config.ConectarBaseDeDatos(); err != nil {
		utils.LogFatal("Error al conectar a la base de datos: " + err.Error())
	}

	utils.LogSuccess("Conexión exitosa a PostgreSQL")

	router := routes.SetupRouter()

	utils.LogInfo("Iniciando servidor en http://localhost:8080")
	utils.LogSuccess("Servidor corriendo en :8080")
	utils.LogWarn("Modo producción recomendado: export GIN_MODE=release")

	router.Run(":8080")
}

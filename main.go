package main

import (
    "backend-go/src/config"
    "backend-go/src/routes"
    "log"
)

func main() {
    config.LoadEnv()
    config.ConectarBaseDeDatos()
    router := routes.SetupRouter()
    log.Println("Servidor corriendo en :8080")
    router.Run(":8080")
}

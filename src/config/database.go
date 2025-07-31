package config

import (
	"fmt"
	"log"
	"os"

	"backend-go/src/models" // 👈 Agrega esta línea para importar el modelo

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBaseDeDatos() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	DB = db
	fmt.Println("✅ Conexión exitosa a la base de datos PostgreSQL")

	// 👇 Crea el esquema si no existe
	if err := DB.Exec("CREATE SCHEMA IF NOT EXISTS private").Error; err != nil {
		log.Fatalf("❌ Error creando esquema 'private': %v", err)
	}

	// 👇 Ejecuta la migración automática
	err = db.AutoMigrate(&models.AuthUser{})
	if err != nil {
		log.Fatalf("❌ Error al migrar modelos: %v", err)
	}
}

package config

import (
	"fmt"
	"os"

	"backend-go/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBaseDeDatos() error {
	// Cargar .env
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error cargando .env: %w", err)
	}

	// Configuración desde variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	DB = db

	// Crear esquema
	if err := DB.Exec("CREATE SCHEMA IF NOT EXISTS private").Error; err != nil {
		return fmt.Errorf("error creando esquema 'private': %w", err)
	}

	// Migraciones
	if err := DB.AutoMigrate(&models.AuthUser{}); err != nil {
		return fmt.Errorf("error al migrar modelos: %w", err)
	}

	return nil
}

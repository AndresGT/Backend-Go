package controllers

import (
	"net/http"
	"os"
	"time"

	"backend-go/src/config"
	"backend-go/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var input AuthUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}

	user := models.AuthUser{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo registrar el usuario"})
		return
	}

	// 🔐 Generar token JWT igual que en Login
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falta configurar JWT_SECRET"})
		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// 🍪 Establecer cookie segura
	c.SetCookie(
		"token",
		tokenString,
		60*60*24*3, // 3 días
		"/",
		"",
		true,
		true,
	)

	// 📦 También devolver token en JSON (por si es app móvil)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado correctamente",
		"token":   tokenString,
	})
}

func Login(c *gin.Context) {
	var input AuthUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	var user models.AuthUser
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Crear JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falta configurar JWT_SECRET"})
		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// 👉 Establecer token como cookie segura (para navegador)
	c.SetCookie(
		"token",     // nombre
		tokenString, // valor
		60*60*24*3,  // duración (3 días en segundos)
		"/",         // path
		"",          // dominio ("" = mismo host)
		true,        // Secure (solo por HTTPS, en dev puedes usar false)
		true,        // HttpOnly (no accesible por JS)
	)

	// 👉 También enviarlo por JSON (para apps móviles)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   tokenString,
	})
}

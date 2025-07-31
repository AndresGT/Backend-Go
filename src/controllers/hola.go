package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HolaMundo(c *gin.Context) {
	userID := c.MustGet("user_id").(float64)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola mundo autenticado!",
		"user_id": userID,
	})
}


package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthUser struct {
	ID       string `gorm:"primaryKey;type:uuid"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
}

// Hook para generar UUID antes de crear el registro
func (user *AuthUser) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}

// Tabla personalizada en esquema "private"
func (AuthUser) TableName() string {
	return "private.users_auth"
}

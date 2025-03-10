package models

import (
	// "database/sql"
	// "log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// User Model

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // "admin" or "user"
	ID       int    `json:"id"`
}

// PostgreSQL bağlantı bilgileri
const (
	DB_USER     = "your_db_user"
	DB_PASSWORD = "your_db_password"
	DB_NAME     = "your_db_name"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

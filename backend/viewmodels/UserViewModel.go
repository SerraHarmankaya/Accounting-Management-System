package viewmodels

import (
	"database/sql"
	//"github.com/serraharmankaya/go-login/models"
)

// UserViewModel yapısı
type UserViewModel struct {
	DB *sql.DB
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"net/http"
	"time"

	"go-login/backend/database"
	"go-login/backend/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key") // Gerçek projelerde .env kullan!

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// Kullanıcı Girişi (Login)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Geçersiz istek türü", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	// Kullanıcıyı veritabanından al
	var storedUser models.User
	query := `SELECT id, name, email, password, role FROM users WHERE email = $1`
	err = database.DB.QueryRow(query, user.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password, &storedUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Kullanıcı bulunamadı", http.StatusUnauthorized)
		} else {
			http.Error(w, "Veritabanı hatası", http.StatusInternalServerError)
		}
		return
	}

	// Şifreyi karşılaştır
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Yanlış şifre", http.StatusUnauthorized)
		return
	}

	// JWT Token oluştur
	expirationTime := time.Now().Add(24 * time.Hour) // Token süresi: 24 saat
	claims := &Claims{
		Email: storedUser.Email,
		Role:  storedUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Token oluşturulamadı", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"token":   tokenString,
		"role":    storedUser.Role,
	})
	// Yanıt olarak token gönder
	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"token": tokenString,})
}

// Yeni Kullanıcı Oluşturma (Create User)
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Geçersiz istek türü", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	// Şifreyi hash'le
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Şifre oluşturulamadı", http.StatusInternalServerError)
		return
	}

	if user.Role == "" {
		http.Error(w, "Role bilgisi eksik", http.StatusBadRequest)
		return
	}

	// Kullanıcıyı veritabanına ekle
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(query, user.Name, user.Email, hashedPassword, user.Role)
	if err != nil {
		http.Error(w, "Veritabanı hatası"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": ' ',
	})
}

func AuthCheckHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token eksik", http.StatusUnauthorized)
		return
	}

	// Header'dan token'ı al
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Yetkilendirme hatası: Token eksik", http.StatusUnauthorized)
		return
	}

	// "Bearer " kısmını kaldırarak token'ı al
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	parseToken(tokenString)

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func FetchClients(w http.ResponseWriter, r *http.Request) {
	// Yalnızca GET isteği kabul edilir
	if r.Method != http.MethodGet {
		http.Error(w, "Geçersiz istek türü", http.StatusMethodNotAllowed)
		return
	}

	// Authorization header'dan token alınır
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header bulunamadı", http.StatusUnauthorized)
		return
	}
	// "Bearer " kısmını kaldırarak token'ı al
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Token doğrulama ve Claims oluşturma
	claims, err := parseToken(tokenString)
	if err != nil {
		http.Error(w, "Geçersiz token", http.StatusUnauthorized)
		return
	}

	// Kullanıcı rolünü kontrol et, sadece admin rolü erişebilir
	if claims.Role != "admin" {
		http.Error(w, "Yalnızca admin müşterilere erişebilir", http.StatusForbidden)
		return
	}

	// Veritabanından kullanıcıları (müşterileri) alacak sorgu
	query := `
		SELECT id, name, email, role 
		FROM users 
		WHERE role = $1
	`
	rows, err := database.DB.Query(query, "user")
	if err != nil {
		log.Printf("Veritabanı sorgu hatası: %v", err)
		http.Error(w, "Veritabanı hatası", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clients []models.User
	for rows.Next() {
		var client models.User
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Role); err != nil {
			log.Printf("Veri tarama hatası: %v", err)
			http.Error(w, "Veri alınırken hata oluştu", http.StatusInternalServerError)
			return
		}
		clients = append(clients, client)
	}

	// rows.Next() döngüsünden sonra oluşabilecek hataları kontrol et
	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		http.Error(w, "Veri alınırken hata oluştu", http.StatusInternalServerError)
		return
	}

	// JSON olarak kullanıcıları döndürme
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "JSON dönüştürme hatası", http.StatusInternalServerError)
	}
}

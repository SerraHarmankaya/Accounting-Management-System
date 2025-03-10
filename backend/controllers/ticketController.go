package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go-login/backend/models"

	"github.com/dgrijalva/jwt-go"

	"go-login/backend/database"
)

// Ticket oluşturma fonksiyonu (Sadece "user" rolü için)
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Geçersiz istek türü", http.StatusMethodNotAllowed)
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

	// Token'ı doğrula ve çöz
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Yetkilendirme hatası: Geçersiz token", http.StatusUnauthorized)
		return
	}

	// Kullanıcıyı email üzerinden bul
	var userID int
	query := `SELECT id FROM users WHERE email = $1`
	err = database.DB.QueryRow(query, claims.Email).Scan(&userID)
	if err != nil {
		http.Error(w, "Kullanıcı bulunamadı", http.StatusUnauthorized)
		return
	}

	// Kullanıcı rolünü kontrol et
	if claims.Role != "user" {
		http.Error(w, "Sadece kullanıcılar ticket oluşturabilir", http.StatusForbidden)
		return
	}

	var ticket models.Ticket
	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	// Ticket'ı veritabanına ekle
	query = `INSERT INTO tickets (title, content, created_by) VALUES ($1, $2, $3) RETURNING id`
	err = database.DB.QueryRow(query, ticket.Title, ticket.Content, userID).Scan(&ticket.ID)
	if err != nil {
		http.Error(w, "Veritabanı hatası", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Ticket başarıyla oluşturuldu",
		"ticket_id": ticket.ID,
		"success":   true,
	})
}

// Admin'in tüm ticket'ları görüntülemesi
func GetAllTickets(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Yalnızca admin erişebilir", http.StatusForbidden)
		return
	}

	// Veritabanından tüm ticket'ları alacak sorgu
	rows, err := database.DB.Query(`SELECT id, title, content, created_at, status, created_by FROM tickets`)
	if err != nil {
		http.Error(w, "Veritabanı hatası", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Content, &ticket.CreatedAt, &ticket.Status, &ticket.CreatedBy); err != nil {
			http.Error(w, "Ticket verisi alınırken hata oluştu", http.StatusInternalServerError)
			return
		}
		tickets = append(tickets, ticket)
	}

	// JSON olarak ticket'ları döndürme
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// User kendi ticketlarini goruntuleme

func FetchUserTickets(w http.ResponseWriter, r *http.Request) {
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
	if claims.Role != "user" {
		http.Error(w, "Yalnızca user kendi ticketlarina erişebilir", http.StatusForbidden)
		return
	}

	// Kullanıcı emailini claims'den alıyoruz
	userEmail := claims.Email // Bu değeri, token'dan çıkararak kullanıcının emailini alıyoruz

	// Veritabanından kullanıcının ticket'larını alacak sorgu
	rows, err := database.DB.Query(`SELECT id, title, content, created_at, status, created_by FROM tickets WHERE created_by = (SELECT id FROM users WHERE email = $1)`, userEmail)
	if err != nil {
		log.Printf("Veritabanı sorgu hatası: %v", err)
		http.Error(w, "Veritabanı hatası", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Content, &ticket.CreatedAt, &ticket.Status, &ticket.CreatedBy); err != nil {
			http.Error(w, "Ticket verisi alınırken hata oluştu", http.StatusInternalServerError)
			return
		}
		tickets = append(tickets, ticket)
	}

	// JSON olarak ticket'ları döndürme
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// Token'ı çözme ve claims elde etme
func parseToken(tokenString string) (*Claims, error) {
	// Token doğrulama işlemi
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Burada JWT'nin doğrulama anahtarını belirtiyoruz
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	// Claims kısmını alıyoruz
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("token geçersiz")
}

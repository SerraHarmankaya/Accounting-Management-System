package controllers

import (
	"Account-Management-System/backend/database"
	"Account-Management-System/backend/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

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

package main

import (
	"fmt"
	"log"
	"net/http"

	"go-login/backend/database"
	"go-login/backend/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	// CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9001"}, // Frontend URL’si
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	port := "9000"
	fmt.Println("🚀 Server is running on port 9000")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// // PostgreSQL bağlantı dizesi
// connStr := "host=localhost port=5432 user=postgres password=123456 dbname=TicketSystem sslmode=disable"

// // Veritabanına bağlan
// var err error
// db, err = sql.Open("postgres", connStr)
// if err != nil {
// 	log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
// }

// // Bağlantıyı test et
// err = db.Ping()
// if err != nil {
// 	log.Println("Veritabanı hatası:", err)
// 	log.Fatal("Veritabanına bağlanılamadı:", err)
// } else {
// 	fmt.Println("✅ PostgreSQL bağlantısı başarılı!")
// }

// // HTTP server ayarları
// mux := http.NewServeMux()
// mux.HandleFunc("/api/login", controllers.CreateUser)

// // CORS ayarları
// corsHandler := cors.New(cors.Options{
// 	AllowedOrigins:   []string{"http://localhost:9000"}, // Sadece Vue frontend portu
// 	AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
// 	AllowedHeaders:   []string{"Content-Type"},
// 	AllowCredentials: true,
// }).Handler(mux)

// // HTTP sunucusunu başlat (port: 9000)
// port := ":9000"
// fmt.Println("🚀 Sunucu çalışıyor: http://localhost" + port)
// log.Fatal(http.ListenAndServe(port, corsHandler))

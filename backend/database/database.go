package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDB() error {
	// PostgreSQL bağlantı dizesi
	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=TicketSystem sslmode=disable"

	// Veritabanına bağlan
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
		return err
	}

	// Bağlantıyı test et
	err = DB.Ping()
	if err != nil {
		log.Println("Veritabanı hatası:", err)
		log.Fatal("Veritabanına bağlanılamadı:", err)
		return err
	} else {
		fmt.Println("✅ PostgreSQL bağlantısı başarılı!")
		return nil
	}

}

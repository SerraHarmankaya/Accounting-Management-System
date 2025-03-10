package routes

import (
	"go-login/backend/controllers"

	"github.com/gorilla/mux"
)

// Uygulama yönlendirmelerini başlatan fonksiyon
func SetupRoutes(router *mux.Router) {
	// Kullanıcı Kaydı
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	//http.HandleFunc("/register", controllers.CreateUser)
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")

	// User ticket olusturma
	router.HandleFunc("/create/user-tickets", controllers.CreateTicket).Methods("POST")

	// User kendi ticketlarini goruntuleme
	router.HandleFunc("/user/tickets", controllers.FetchUserTickets).Methods("GET")

	// Admin ticket goruntuleme
	router.HandleFunc("/admin-page", controllers.GetAllTickets).Methods("GET")

	// Check Auth token
	router.HandleFunc("/auth-check", controllers.AuthCheckHandler).Methods("GET")

	// Admin tarafında kullanıcıları gösterme
	router.HandleFunc("/clients", controllers.FetchClients).Methods("GET")

	// Deneme (Profil)
	//router.HandleFunc("/deneme", controllers.updateUser()).Methods("POST")

}

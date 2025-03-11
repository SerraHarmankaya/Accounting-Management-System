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

	// Check Auth token
	router.HandleFunc("/auth-check", controllers.AuthCheckHandler).Methods("GET")

	// Admin tarafında kullanıcıları gösterme
	router.HandleFunc("/clients/list", controllers.FetchClients).Methods("GET")

	// Admin tarafında ticketlari gösterme
	router.HandleFunc("/clients/ticket", controllers.GetAllTickets).Methods("GET")

	// Client(User) silme
	router.HandleFunc("/delete/{id}", controllers.DeleteUser).Methods("DELETE")

	// Deneme (Profil)
	//router.HandleFunc("/deneme", controllers.updateUser()).Methods("POST")

}

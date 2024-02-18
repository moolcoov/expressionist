package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes Настраивает эндпоинты
func SetupRoutes(r *mux.Router) {
	// GET ручки
	r.HandleFunc("/get", GetHandler).Methods("GET")
	r.HandleFunc("/list", ListHandler).Methods("GET")
	r.HandleFunc("/settings", SettingsHandler).Methods("GET")
	r.HandleFunc("/ping", PingHandler).Methods("GET")
	r.HandleFunc("/agents", AgentsHandler).Methods("GET")

	// POST ручки
	r.HandleFunc("/new", NewHandler).Methods("POST")
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/settings/update", SettingsUpdateHandler).Methods("POST")
	r.HandleFunc("/submit", SubmitHandler).Methods("POST")

	http.Handle("/", r)
}

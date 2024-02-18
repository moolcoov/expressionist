package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"orchestra/agent"
)

// RegisterHandler Хендлер для роута /register
// Регистрирует агента.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение body в структуру
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not correct", 500)
		fmt.Println(err)
		return
	}

	var a agent.Agent

	err = json.Unmarshal(body, &a)
	if err != nil {
		http.Error(w, "provided data is not correct", http.StatusNotAcceptable)
		fmt.Println(err)
		return
	}

	// Регистрация агента
	a.Register()
}

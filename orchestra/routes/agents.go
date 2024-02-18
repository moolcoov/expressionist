package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orchestra/agent"
)

// AgentsHandler Хендлер для роута /agents
// Возвращает агентов
func AgentsHandler(w http.ResponseWriter, r *http.Request) {
	// Запись json в ответ
	res, err := json.Marshal(agent.Agents)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting expressions", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

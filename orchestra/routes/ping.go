package routes

import (
	"fmt"
	"net/http"
	"time"

	"orchestra/agent"

	"github.com/google/uuid"
)

// PingHandler Хендлер для роута /ping
// Обновляет время пинга агента
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// Получение id агента из параметров запроса
	agentId := r.URL.Query().Get("id")

	if agentId == "" {
		http.Error(w, "id is required", http.StatusNotAcceptable)
		return
	}

	uid, err := uuid.Parse(agentId) // Парсинг uuid
	if err != nil {
		fmt.Println(err)
		http.Error(w, "id is not correct", http.StatusNotAcceptable)
		return
	}

	a := agent.Agent{Id: uid}

	// Обновляем время пинга агента, если он существует
	for _, ag := range agent.Agents.AgentsList {
		if ag.Id == a.Id {
			agent.Agents.Mu.Lock()
			ag.PingTime = time.Now()
			agent.Agents.Mu.Unlock()

			return
		}
	}

	// В противном случае регистрируем как нового
	a.Register()
}

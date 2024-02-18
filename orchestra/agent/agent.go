package agent

import (
	"github.com/google/uuid"
	"time"
)

type Agent struct {
	Id       uuid.UUID `json:"id"`       // id агента
	PingTime time.Time `json:"pingTime"` // последнее время запроса
	Status   string    `json:"status"`   // статус агента "active" | "inactive"
}

// Register регистрирует агента
func (a *Agent) Register() {
	// pingTime ставится на текущее время
	a.PingTime = time.Now()

	// Агент изначально активен
	a.Status = "active"

	// Добавление в слайс агентов
	Agents.AgentsList = append([]*Agent{a}, Agents.AgentsList...)
}

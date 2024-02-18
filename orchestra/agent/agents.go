package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"orchestra/lib"
	"orchestra/schemas"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Agents = &AgentsWithMutex{Mu: &sync.RWMutex{}}
)

type AgentsWithMutex struct {
	AgentsList []*Agent      `json:"agents"`
	Mu         *sync.RWMutex `json:"-"`
}

// CheckAgents проверяет агентов
func (a *AgentsWithMutex) CheckAgents() {
	lib.Settings.Mu.RLock()
	// время по истечении которого агент без пинга становится неактивным
	activeTime := 40 * time.Second
	// время по истечении которого агент удаляется
	inactiveTime := time.Duration(lib.Settings.InactiveAgentTime+40) * time.Second
	lib.Settings.Mu.RUnlock()

	a.Mu.RLock()
	for i, agent := range a.AgentsList {
		// Если агент активен, но время ожидания пинга превысило допустимое
		if agent.Status == "active" && time.Since(agent.PingTime) > activeTime {
			// Агент становится неактивным
			agent.Status = "inactive"

			// Достаем из бд выражения этого агента, которые вычисляются
			var expressions []schemas.Expression

			query := fmt.Sprintf("SELECT * from expressions WHERE agent_id::text='%s' AND status='in_progress'", agent.Id.String())

			lib.Pg.Mu.RLock()
			err := lib.Pg.Client.Select(&expressions, query)
			lib.Pg.Mu.RUnlock()

			fmt.Println(expressions)

			if err != nil {
				fmt.Println("ERR: Failed to get agent's expressions")
				fmt.Println(err)
				continue
			}

			// Каждое отправляем новому агенту через RabbitMQ
			for _, expression := range expressions {
				// Запись ответа в json
				res, err := json.Marshal(&expression)
				if err != nil {
					return
				}

				lib.RbMQ.Mu.Lock()

				queue, err := lib.RbMQ.Channel.QueueDeclare(
					lib.Getenv("RABBITMQ_QUEUE", "expressions"),
					false,
					false,
					false,
					false,
					nil,
				)

				if err != nil {
					fmt.Println("ERR: Failed to open a RabbitMQ queue")
					fmt.Println(err.Error())
					return
				}

				err = lib.RbMQ.Channel.PublishWithContext(context.Background(), "",
					queue.Name,
					false,
					false,
					amqp.Publishing{
						Type: "application/json",
						Body: res,
					})

				lib.RbMQ.Mu.Unlock()

				if err != nil {
					fmt.Println("ERR: Failed to send an expression via RabbitMQ queue")
					fmt.Println(err.Error())
					return
				}
			}
		}

		// Если агент неактивен и превысило время неактивного агента
		if agent.Status == "inactive" && time.Since(agent.PingTime) > inactiveTime {
			// Удаляем агента
			a.AgentsList = append(a.AgentsList[:i], a.AgentsList[i+1:]...)
		}
	}
	a.Mu.RUnlock()
}

package main

import (
	"agent/lib"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Подключаем .env файлы
	if os.Getenv("ENVIRONMENT") != "docker" {
		godotenv.Load("../.env")
		godotenv.Load("../.env.local")
	}

	lib.Setup()

	// открываем очередь
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
		log.Fatal(err.Error())
	}

	msgs, err := lib.RbMQ.Channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("ERR: Failed to register a RabbitMQ consumer")
		log.Fatal(err.Error())
	}

	// получаем число - количество горутин
	goroutinesCount, err := strconv.Atoi(lib.Getenv("AGENT_GOROUTINES", "5"))
	if err != nil {
		goroutinesCount = 5
	}

	for i := 0; i < goroutinesCount; i++ {
		go func() {
			for {
				// получаем сообщение
				message := <-msgs

				body, err := io.ReadAll(bytes.NewReader(message.Body))
				if err != nil {
					return
				}

				var e lib.Expression

				err = json.Unmarshal(body, &e)
				if err != nil {
					return
				}

				// делаем статус вычисляется
				e.Status = "in_progress"
				e.AgentId = lib.Id
				e.Submit()

				// вычисляем
				e.Calculate()
				e.Submit()
			}
		}()
	}

	// цикл, который каждые 30 секунд пингует бэкенд и обновляет настройки
	for {
		time.Sleep(30 * time.Second)

		lib.UpdateSettings()
		lib.Ping()
	}
}

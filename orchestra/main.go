package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"orchestra/agent"
	"orchestra/lib"
	"orchestra/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Подключаем .env файлы
	if os.Getenv("ENVIRONMENT") != "docker" {
		godotenv.Load("../.env")
	}

	// Настраиваем сервисы
	lib.Setup()

	// Настраиваем эндпоинты
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	fmt.Print("\nSUCCESS: Server has successfully started\n\n")

	// Запускаем проверку агентов
	go func() {
		for {
			agent.Agents.CheckAgents()
			time.Sleep(10 * time.Second)
		}
	}()

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}

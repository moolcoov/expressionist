package lib

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

var Settings = &SettingsWithMutex{Mu: &sync.RWMutex{}}

type SettingsWithMutex struct {
	// Встраиваем поля с настройками
	settings

	Mu *sync.RWMutex `json:"-"`
}

type settings struct {
	// Время операций
	AdditionTime       int `json:"additionTime"`
	SubtractionTime    int `json:"subtractionTime"`
	MultiplicationTime int `json:"multiplicationTime"`
	DivisionTime       int `json:"divisionTime"`

	// Время отображения неактивного агента
	InactiveAgentTime int `json:"inactiveAgentTime"`
}

func UpdateSettings() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", Getenv("BACKEND_ADDRESS", "http://orchestra:8080")+"/settings", nil)
	if err != nil {
		return
	}

	client := http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	Settings.Mu.Lock()
	err = json.Unmarshal(body, Settings)
	if err != nil {
		return
	}
	Settings.Mu.Unlock()
}

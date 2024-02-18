package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var Id uuid.UUID

func Register() {
	Id = uuid.New()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	values := map[string]string{"id": Id.String()}
	body, err := json.Marshal(values)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", Getenv("BACKEND_ADDRESS", "http://orchestra:8080")+"/register", bytes.NewReader(body))
	if err != nil {
		return
	}

	client := http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode == 406 || res.StatusCode == 500 {
		log.Fatal("FATAL ERR: id is incorrect")
	}

	fmt.Println("SETUP: Successfully registered!")
}

func Ping() {
	req, err := http.NewRequestWithContext(context.Background(), "GET", Getenv("BACKEND_ADDRESS", "http://orchestra:8080")+"/ping?id="+Id.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := http.Client{Timeout: 5 * time.Second}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

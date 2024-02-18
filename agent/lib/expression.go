package lib

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	shuntingYard "github.com/mgenware/go-shunting-yard"
)

const (
	statusCalculated = "calculated"
	statusErrored    = "errored"
)

type Expression struct {
	Id              uuid.UUID    `json:"id"`                    // id выражения
	Expression      string       `json:"expression"`            // выражение
	DispatchTime    sql.NullTime `json:"dispatchTime"`          // время отправки оркестратору
	CalculationTime sql.NullTime `json:"calculationTime"`       // время, когда выражение было вычислено
	Result          string       `json:"res"`                   // результат
	Status          string       `json:"status"`                // статус "created" | "in_progress" | "calculated" | "errored"
	IsSent          bool         `json:"isSent" `               // отправлено ли выражение агентам
	AgentId         uuid.UUID    `json:"agentId" db:"agent_id"` // id агента
}

// Calculate вычисляет значение выражения
func (e *Expression) Calculate() {
	infix, _ := shuntingYard.Scan(e.Expression)

	// получаем постфиксную запись
	postfix, err := parse(infix)
	if err != nil {
		e.error()
		return
	}

	// вычисляем её
	res, err := evaluate(postfix)
	if err != nil {
		e.error()
		return
	}

	// определяем время вычисления
	calculationTime := 0

	for _, token := range infix {
		switch token {
		case "+":
			calculationTime += Settings.AdditionTime
		case "-":
			calculationTime += Settings.SubtractionTime
		case "*":
			calculationTime += Settings.MultiplicationTime
		case "/":
			calculationTime += Settings.DivisionTime
		}
	}

	// ждем и записываем результат
	time.Sleep(time.Duration(calculationTime) * time.Millisecond)

	e.Result = strconv.FormatFloat(res, 'g', -1, 64)
	e.Status = statusCalculated
}

// error делает выражение ошибочным
func (e *Expression) error() {
	e.Status = statusErrored
	e.CalculationTime.Time = time.Now()
}

// Submit отправляет код на бэкенд
func (e *Expression) Submit() {
	body, err := json.Marshal(e)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", Getenv("BACKEND_ADDRESS", "http://orchestra:8080")+"/submit", bytes.NewReader(body))
	if err != nil {
		return
	}

	client := http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return
	}
}

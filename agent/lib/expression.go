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
	statusCreated    = "created"
	statusInProgress = "in_progress"
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

func (e *Expression) Calculate() {
	infix, _ := shuntingYard.Scan(e.Expression)

	postfix, err := shuntingYard.Parse(infix)
	if err != nil {
		e.error()
		return
	}

	res, err := shuntingYard.Evaluate(postfix)
	if err != nil {
		e.error()
		return
	}

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

	time.Sleep(time.Duration(calculationTime) * time.Millisecond)

	e.Result = strconv.Itoa(res)
	e.Status = statusCalculated
}

func (e *Expression) error() {
	e.Status = statusErrored
	e.CalculationTime.Time = time.Now()
}

func (e *Expression) Submit() {
	body, err := json.Marshal(e)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", "http://localhost:8080/submit", bytes.NewReader(body))
	if err != nil {
		return
	}

	client := http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return
	}
}

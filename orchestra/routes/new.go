package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orchestra/lib"
	"orchestra/schemas"
	"strings"
	"time"

	"github.com/google/uuid"
	shuntingYard "github.com/mgenware/go-shunting-yard"
	amqp "github.com/rabbitmq/amqp091-go"
)

// NewHandler Хендлер для роута /new
// Читает из body выражение и добавляет в базу
// Возвращает Expression
func NewHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение body в структуру
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not correct", 500)
		fmt.Println(err)
		return
	}

	var e schemas.Expression

	err = json.Unmarshal(body, &e)
	if err != nil {
		http.Error(w, "provided data is not correct", 500)
		fmt.Println(err)
		return
	}

	// Парсинг начального выражения
	infix, _ := shuntingYard.Scan(e.Expression)
	e.Expression = strings.Join(infix, " ")

	// Добавление необходимых полей
	e.Id = uuid.New()
	e.Status = "created"

	// Проверка в кэше (Redis)
	func() {
		lib.Rdb.Mu.RLock()
		cache := lib.Rdb.Client.HGetAll(r.Context(), e.Expression)
		lib.Rdb.Mu.RUnlock()

		if cache.Err() != nil {
			fmt.Println("ERR: Failed to get from cache")
			fmt.Println(cache.Err())
			return
		}

		cacheInfo := cache.Val()

		// - Проверка результата из кэша
		res := cacheInfo["res"]
		if res != "" {
			e.Result = res
			e.Status = "calculated"
		}

		// - Проверка времени обработки из кэша
		cachedCalculationTime := cacheInfo["calculationTime"]
		calculationTime, err := time.Parse(time.RFC3339Nano, cachedCalculationTime)
		if err == nil {
			e.CalculationTime.Time = calculationTime
		}
	}()

	// Проверка времени отправки и вычисления
	if e.DispatchTime.Time.IsZero() {
		e.DispatchTime.Time = time.Now()
	}
	if e.CalculationTime.Time.IsZero() {
		e.CalculationTime.Time = time.Now()
	}

	e.DispatchTime.Valid = true
	e.CalculationTime.Valid = true

	// Запись ответа в json
	res, err := json.Marshal(&e)
	if err != nil {
		http.Error(w, "error while getting expression", 500)
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Добавление в очередь RabbitMQ
		func() {
			if e.Status == "calculated" {
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

			err = lib.RbMQ.Channel.PublishWithContext(ctx, "",
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
			e.IsSent = true
		}()

		// Insert в базу данных
		query := "INSERT INTO expressions (id, expression, dispatch_time, calculation_time, result, status, is_sent) VALUES ($1, $2, $3, $4, $5, $6, $7)"

		lib.Pg.Mu.Lock()
		_, err := lib.Pg.Client.Exec(query, e.Id, e.Expression, e.DispatchTime, e.CalculationTime, e.Result, e.Status, e.IsSent)
		if err != nil {
			fmt.Println("ERR: Failed to insert an expression to database")
			fmt.Println(err)
			return
		}
		lib.Pg.Mu.Unlock()
	}()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

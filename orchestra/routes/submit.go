package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"orchestra/lib"
	"orchestra/schemas"
	"time"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение body в структуру
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not correct", 500)
		fmt.Println(err)
		return
	}

	var submitted schemas.Expression

	err = json.Unmarshal(body, &submitted)
	if err != nil {
		http.Error(w, "provided data is not correct", 500)
		fmt.Println(err)
		return
	}

	// Обновление выражения в базе данных
	query := fmt.Sprintf("UPDATE expressions SET status = '%s', result = '%s', agent_id = '%s' WHERE id = '%s';", submitted.Status, submitted.Result, submitted.AgentId, submitted.Id.String())

	lib.Pg.Mu.RLock()
	_, err = lib.Pg.Client.Exec(query)
	lib.Pg.Mu.RUnlock()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting expressions", 500)
		return
	}

	// Добавление в кэш
	go func() {
		lib.Rdb.Mu.Lock()
		lib.Rdb.Client.HSet(context.Background(), submitted.Expression,
			[]string{
				"res", submitted.Result,
				"calculationTime", submitted.CalculationTime.Time.Format(time.RFC3339Nano),
			})
		lib.Rdb.Mu.Unlock()
	}()

}

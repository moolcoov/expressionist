package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"orchestra/lib"
	"orchestra/schemas"
)

// ListHandler Хендлер для роута /list
// Получает все выражения из базы данных.
// Возвращает []Expression
func ListHandler(w http.ResponseWriter, r *http.Request) {
	// Получение всех выражений из бд
	var expressions []schemas.Expression

	lib.Pg.Mu.RLock()
	err := lib.Pg.Client.Select(&expressions, "SELECT * from expressions ORDER BY dispatch_time DESC")
	lib.Pg.Mu.RUnlock()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting expressions", 500)
		return
	}

	// Запись json в ответ
	res, err := json.Marshal(expressions)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting expressions", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

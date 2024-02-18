package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orchestra/lib"
	"orchestra/schemas"

	"github.com/google/uuid"
)

// GetHandler Хендлер для роута /get?id={}
// Читает из параметров запроса id выражения и получает его из базы.
// Возвращает Expression
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Получение id выражения из параметров запроса
	paramsId := r.URL.Query().Get("id")

	if paramsId == "" {
		http.Error(w, "id is required", http.StatusNotAcceptable)
		return
	}

	uid, err := uuid.Parse(paramsId) // Парсинг uuid
	if err != nil {
		fmt.Println(err)
		http.Error(w, "id is not correct", http.StatusNotAcceptable)
		return
	}

	// Получение выражения из базы данных
	var e schemas.Expression

	query := fmt.Sprintf("SELECT * from expressions WHERE id::text = '%s'", uid.String())

	lib.Pg.Mu.RLock()
	err = lib.Pg.Client.Get(&e, query)
	lib.Pg.Mu.RUnlock()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting expressions", 500)
		return
	}

	// Запись результата в ответ как json
	res, err := json.Marshal(e)
	if err != nil {
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

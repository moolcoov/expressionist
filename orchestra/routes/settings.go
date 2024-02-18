package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"orchestra/lib"
)

// SettingsHandler Хендлер для роута /settings
// Возвращает значения настроек
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	lib.Settings.Mu.RLock()
	settings := lib.Settings
	lib.Settings.Mu.RUnlock()

	res, err := json.Marshal(settings)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while getting settings", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

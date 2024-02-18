package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"orchestra/lib"
	"orchestra/schemas"
)

// SettingsUpdateHandler Хендлер для роута /settings/update
// Обновляет значения настроек
func SettingsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение body в структуру
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body is not correct", 500)
		fmt.Println(err)
		return
	}

	var s schemas.Settings

	err = json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, "provided data is not correct", 500)
		fmt.Println(err)
		return
	}

	// Переназначаем настройки
	lib.Settings.Mu.Lock()
	lib.Settings.AdditionTime = lib.CheckSetting(s.AdditionTime)
	lib.Settings.SubtractionTime = lib.CheckSetting(s.SubtractionTime)
	lib.Settings.MultiplicationTime = lib.CheckSetting(s.MultiplicationTime)
	lib.Settings.DivisionTime = lib.CheckSetting(s.DivisionTime)
	lib.Settings.InactiveAgentTime = lib.CheckSetting(s.InactiveAgentTime)
	lib.Settings.Mu.Unlock()

	// Обновляем настройки в Redis
	lib.UpdateRedisSettings()
}

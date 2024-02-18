package lib

import (
	"context"
	"fmt"
	"orchestra/schemas"
	"strconv"
	"sync"
)

const (
	defaultSettingVal = 200 // Стандартное значение

	redisSettingsKey         = "_settings"          // Ключ настроек
	redisAdditionField       = "addTime"            // Ключ сложения
	redisSubtractionField    = "subtractTime"       // Ключ вычитания
	redisMultiplicationField = "multiplicationTime" // Ключ умножения
	redisDivisionField       = "divisionTime"       // Ключ деления
	redisInactiveAgentField  = "inactiveAgentTime"  // Ключ неактивного агента
)

var Settings = &SettingsWithMutex{}

type SettingsWithMutex struct {
	// Встраиваем поля с настройками
	schemas.Settings

	Mu *sync.RWMutex `json:"-"`
}

func SetupSettings() {
	// Берем настройки из Redis
	Rdb.Mu.RLock()
	settingsHash := Rdb.Client.HGetAll(context.Background(), redisSettingsKey)
	Rdb.Mu.RUnlock()

	settings := settingsHash.Val()

	// Подставляем их в структуру
	Settings.AdditionTime = getSetting(redisAdditionField, settings)
	Settings.SubtractionTime = getSetting(redisSubtractionField, settings)
	Settings.MultiplicationTime = getSetting(redisMultiplicationField, settings)
	Settings.DivisionTime = getSetting(redisDivisionField, settings)
	Settings.InactiveAgentTime = getSetting(redisInactiveAgentField, settings)

	Settings.Mu = &sync.RWMutex{}

	fmt.Println("SETUP: Settings have been configured")
}

func getSetting(key string, settings map[string]string) int {
	// Преобразуем значение из Redis
	val, err := strconv.Atoi(settings[key])
	if settings[key] == "" || err != nil {
		// В случае неудачи ставим стандартное значение и записываем его в Redis
		Rdb.Mu.Lock()
		Rdb.Client.HSetNX(context.Background(), redisSettingsKey, key, defaultSettingVal)
		Rdb.Mu.Unlock()

		return defaultSettingVal
	}
	return val
}

func CheckSetting(settingValue int) int {
	if settingValue >= 0 && settingValue <= 999_999_999 {
		return settingValue
	}
	return defaultSettingVal
}

func UpdateRedisSettings() {
	settings := map[string]interface{}{
		redisAdditionField:       Settings.AdditionTime,
		redisSubtractionField:    Settings.SubtractionTime,
		redisMultiplicationField: Settings.MultiplicationTime,
		redisDivisionField:       Settings.DivisionTime,
		redisInactiveAgentField:  Settings.InactiveAgentTime,
	}

	Rdb.Mu.Lock()
	Rdb.Client.HSet(context.Background(), redisSettingsKey, settings)
	Rdb.Mu.Unlock()
}

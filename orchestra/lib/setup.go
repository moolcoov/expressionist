package lib

// Setup Настраивает сервисы
func Setup() {
	SetupPostgres()
	SetupRedis()
	SetupRabbit()
	SetupSettings()
}

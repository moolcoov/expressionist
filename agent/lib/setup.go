package lib

import "fmt"

// Setup настраивает всё
func Setup() {
	SetupRabbit()
	Register()
	UpdateSettings()

	fmt.Print("\nSUCCESS: Agent has successfully started\n\n")
}

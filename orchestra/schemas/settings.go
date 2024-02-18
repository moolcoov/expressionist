package schemas

type Settings struct {
	// Время операций
	AdditionTime       int `json:"additionTime"`
	SubtractionTime    int `json:"subtractionTime"`
	MultiplicationTime int `json:"multiplicationTime"`
	DivisionTime       int `json:"divisionTime"`

	// Время отображения неактивного агента
	InactiveAgentTime int `json:"inactiveAgentTime"`
}

package schemas

import (
	"database/sql"
	"github.com/google/uuid"
)

type Expression struct {
	Id              uuid.UUID    `json:"id" db:"id"`                            // id выражения
	Expression      string       `json:"expression" db:"expression"`            // выражение
	DispatchTime    sql.NullTime `json:"dispatchTime" db:"dispatch_time"`       // время отправки оркестратору
	CalculationTime sql.NullTime `json:"calculationTime" db:"calculation_time"` // время, когда выражение было вычислено
	Result          string       `json:"res" db:"result"`                       // результат
	Status          string       `json:"status" db:"status"`                    // статус "created" | "in_progress" | "calculated"
	IsSent          bool         `json:"isSent" db:"is_sent"`                   // отправлено ли выражение агентам
	AgentId         uuid.UUID    `json:"agentId" db:"agent_id"`                 // id агента
}

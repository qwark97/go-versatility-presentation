package conf

import "github.com/google/uuid"

type configuration struct {
	ID          uuid.UUID `json:"id"`
	Method      string    `json:"method"`
	Addr        string    `json:"addr"`
	Frequency   string    `json:"frequency"`
	Description string    `json:"description"`
	Unit        string    `json:"unit"`
}

type verification struct {
	Success bool `json:"success"`
}

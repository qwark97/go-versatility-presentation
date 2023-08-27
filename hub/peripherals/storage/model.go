package storage

import "github.com/google/uuid"

type Configuration struct {
	ID          uuid.UUID `json:"id"`
	Method      string    `json:"method"`
	Addr        string    `json:"addr"`
	Frequency   string    `json:"frequency"`
	Description string    `json:"description"`
	Unit        string    `json:"unit"`
}

package storage

import "github.com/google/uuid"

type Configuration struct {
	ID          uuid.UUID `json:"id"`
	Method      string    `json:"method"`
	Addr        string    `json:"addr"`
	Frequency   string    `json:"frequency"`
	Description string    `json:"description"`
	Unit        Unit      `json:"unit"`
}

type Unit string

func (u Unit) String() string {
	if translatedUnit, known := predefinedUnits[u]; known {
		return translatedUnit
	} else {
		return string(u)
	}
}

var predefinedUnits = map[Unit]string{
	"celsius": "°C",
	"percent": "%",
}

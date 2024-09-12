package tdados

import (
	"time"
)

type EstructuraDoS interface {
	CrearArrAtacantes() []string
	AñadirVisita(string, time.Time)
}

package models

import (
	"time"
)

type Solicitud struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	Pasajero string    `json:"pasajero"`
	Chofer   string    `json:"chofer"`
	Origen   string    `json:"origen"`
	Destino  string    `json:"destino"`
	Estado   string    `json:"estado"`
	CreadoEn time.Time `json:"creado_en"`
}

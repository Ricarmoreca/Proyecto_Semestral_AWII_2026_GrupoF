package models

import (
	"errors"
	"time"
)

type DespachoDiario struct {
	ID                int    `gorm:"primaryKey;autoIncrement;column:id_despacho"`
	Fecha             string `gorm:"column:fecha"`
	NumeroCarrito     int    `gorm:"column:numero_carrito"`
	IDHorario         int    `gorm:"column:id_horario"`
	IDChofer          int    `gorm:"column:id_chofer"`
	PasajerosActuales int    `gorm:"column:pasajeros_actuales"`
}

func (DespachoDiario) TableName() string {
	return "despachos_diarios"
}

func (d *DespachoDiario) Validate() error {
	if d.ID < 0 {
		return errors.New("id_despacho no puede ser negativo")
	}
	if _, err := time.Parse("2006-01-02", d.Fecha); err != nil {
		return errors.New("fecha debe tener formato YYYY-MM-DD")
	}
	if d.NumeroCarrito <= 0 {
		return errors.New("numero_carrito debe ser un entero positivo")
	}
	if d.IDHorario <= 0 {
		return errors.New("id_horario debe ser un entero positivo")
	}
	if d.IDChofer <= 0 {
		return errors.New("id_chofer debe ser un entero positivo")
	}
	if d.PasajerosActuales < 0 {
		return errors.New("pasajeros_actuales no puede ser negativo")
	}
	return nil
}

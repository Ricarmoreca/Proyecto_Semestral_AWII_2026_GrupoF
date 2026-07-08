package models

import "errors"

type EstadoMantenimiento string

const (
	EstadoMantenimientoPendiente  EstadoMantenimiento = "Pendiente"
	EstadoMantenimientoEnProgreso EstadoMantenimiento = "En Progreso"
	EstadoMantenimientoCompletado EstadoMantenimiento = "Completado"
	EstadoMantenimientoCancelado  EstadoMantenimiento = "Cancelado"
)

func (e EstadoMantenimiento) Valid() bool {
	switch e {
	case EstadoMantenimientoPendiente, EstadoMantenimientoEnProgreso, EstadoMantenimientoCompletado, EstadoMantenimientoCancelado:
		return true
	default:
		return false
	}
}

type Mantenimiento struct {
	ID                  int                 `gorm:"primaryKey;autoIncrement;column:id_mantenimiento"`
	FechaMantenimiento  string              `gorm:"column:fecha_mantenimiento"`
	Descripcion         string              `gorm:"column:descripcion"`
	EstadoMantenimiento EstadoMantenimiento `gorm:"column:estado_mantenimiento"`
	NumeroCarrito       string              `gorm:"column:numero_carrito"`
}

func (Mantenimiento) TableName() string {
	return "mantenimientos"
}

func (m *Mantenimiento) Validate() error {
	if m.FechaMantenimiento == "" {
		return errors.New("fecha_mantenimiento es requerida")
	}
	if m.Descripcion == "" {
		return errors.New("descripcion es requerida")
	}
	if m.NumeroCarrito == "" {
		return errors.New("numero_carrito es requerido")
	}
	if !m.EstadoMantenimiento.Valid() {
		return errors.New("estado_mantenimiento no es válido")
	}
	return nil
}

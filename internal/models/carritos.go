package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type EstadoCarrito string

const (
	EstadoCarritoDisponible    EstadoCarrito = "Disponible"
	EstadoCarritoEnViaje       EstadoCarrito = "En Viaje"
	EstadoCarritoMantenimiento EstadoCarrito = "Mantenimiento"
)

func (e EstadoCarrito) Valid() bool {
	switch e {
	case EstadoCarritoDisponible, EstadoCarritoEnViaje, EstadoCarritoMantenimiento:
		return true
	default:
		return false
	}
}

type Carrito struct {
	Numero             int            `json:"numero_carrito" gorm:"primaryKey;column:numero_carrito"`
	Estado             EstadoCarrito  `json:"estado_carrito" gorm:"column:estado_carrito;default:Disponible"`
	CapacidadPasajeros int            `json:"capacidad_pasajeros" gorm:"column:capacidad_pasajeros"`
	Color              string         `json:"color" gorm:"column:color;default:Sin color"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

func (Carrito) TableName() string {
	return "carritos"
}

func (c *Carrito) Validate() error {
	if c.Numero <= 0 {
		return errors.New("numero_carrito debe ser un entero positivo")
	}
	if !c.Estado.Valid() {
		return errors.New("estado_carrito debe ser 'Disponible', 'En Viaje' o 'Mantenimiento'")
	}
	if c.CapacidadPasajeros <= 0 {
		return errors.New("capacidad_pasajeros debe ser un entero positivo")
	}
	if c.Color == "" {
		return errors.New("color es requerido")
	}
	return nil
}

type CarritoResponse struct {
	Numero             int           `json:"numero_carrito"`
	Estado             EstadoCarrito `json:"estado_carrito"`
	CapacidadPasajeros int           `json:"capacidad_pasajeros"`
	Color              string        `json:"color"`
	DeletedAt          *time.Time    `json:"deleted_at,omitempty"`
}

func (c *Carrito) ToResponse() CarritoResponse {
	resp := CarritoResponse{
		Numero:             c.Numero,
		Estado:             c.Estado,
		CapacidadPasajeros: c.CapacidadPasajeros,
		Color:              c.Color,
	}
	if c.DeletedAt.Valid {
		resp.DeletedAt = &c.DeletedAt.Time
	}
	return resp
}

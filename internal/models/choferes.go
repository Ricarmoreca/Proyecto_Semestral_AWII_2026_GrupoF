package models

import "errors"

type EstadoChofer string

const (
	EstadoChoferDisponible EstadoChofer = "Disponible"
	EstadoChoferEnRuta     EstadoChofer = "En Ruta"
	EstadoChoferDescanso   EstadoChofer = "Descanso"
)

func (e EstadoChofer) Valid() bool {
	switch e {
	case EstadoChoferDisponible, EstadoChoferEnRuta, EstadoChoferDescanso:
		return true
	default:
		return false
	}
}

type Chofer struct {
	ID       int          `gorm:"primaryKey;autoIncrement;column:id_chofer"`
	Nombre   string       `gorm:"column:nombre_chofer"`
	Licencia string       `gorm:"column:licencia;uniqueIndex"`
	Celular  string       `gorm:"column:celular"`
	Estado   EstadoChofer `gorm:"column:estado_chofer"`
}

func (Chofer) TableName() string {
	return "choferes"
}

func (c *Chofer) Validate() error {
	if c.Nombre == "" {
		return errors.New("nombre_chofer es requerido")
	}
	if c.Licencia == "" {
		return errors.New("licencia es requerida")
	}
	if c.Celular == "" {
		return errors.New("celular es requerido")
	}
	if !c.Estado.Valid() {
		return errors.New("estado_chofer no es valido")
	}
	return nil
}

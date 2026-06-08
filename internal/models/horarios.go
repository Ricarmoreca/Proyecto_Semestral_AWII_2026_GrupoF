package models

import "errors"

type Turno string

const (
	TurnoMatutina   Turno = "Matutina"
	TurnoVespertina Turno = "Vespertina"
)

func (t Turno) Valid() bool {
	switch t {
	case TurnoMatutina, TurnoVespertina:
		return true
	default:
		return false
	}
}

type Horario struct {
	ID         int    `gorm:"primaryKey;autoIncrement;column:id_horario"`
	Turno      Turno  `gorm:"column:turno"`
	HoraInicio string `gorm:"column:hora_inicio"`
	HoraFin    string `gorm:"column:hora_fin"`
}

func (Horario) TableName() string {
	return "horarios"
}

func (h *Horario) Validate() error {
	if !h.Turno.Valid() {
		return errors.New("turno debe ser 'Matutina' o 'Vespertina'")
	}
	if len(h.HoraInicio) != 5 || h.HoraInicio[2] != ':' {
		return errors.New("hora_inicio debe tener formato HH:MM")
	}
	if len(h.HoraFin) != 5 || h.HoraFin[2] != ':' {
		return errors.New("hora_fin debe tener formato HH:MM")
	}
	return nil
}

type CarritoHorarioRel struct {
	NumeroCarrito  int    `gorm:"column:numero_carrito;primaryKey"`
	IDHorario      int    `gorm:"column:id_horario;primaryKey"`
	HoraAsignacion string `gorm:"column:hora_asignacion"`
}

func (CarritoHorarioRel) TableName() string {
	return "carrito_horario"
}

type CarritoHorario struct {
	NumeroCarrito      int    `json:"numero_carrito,omitempty"`
	IDHorario          int    `json:"id_horario,omitempty"`
	HoraAsignacion     string `json:"hora_asignacion,omitempty"`
	Turno              string `json:"turno,omitempty"`
	HoraInicio         string `json:"hora_inicio,omitempty"`
	HoraFin            string `json:"hora_fin,omitempty"`
	EstadoCarrito      string `json:"estado_carrito,omitempty"`
	CapacidadPasajeros int    `json:"capacidad_pasajeros,omitempty"`
	Color              string `json:"color,omitempty"`
}

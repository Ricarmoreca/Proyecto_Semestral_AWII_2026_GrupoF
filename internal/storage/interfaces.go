package storage

import "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"

type CarritoRepository interface {
	GetAllCarritos() ([]models.Carrito, error)
	GetCarritoByNumero(numero int) (models.Carrito, error)
	CreateCarrito(c models.Carrito) error
	UpdateCarrito(c models.Carrito) error
	DeleteCarrito(numero int) error
	GetHorariosByCarrito(numero int) ([]models.CarritoHorario, error)
}

type ChoferRepository interface {
	GetAllChoferes() ([]models.Chofer, error)
	GetChoferByID(id int) (models.Chofer, error)
	GetChoferByLicencia(licencia string) (models.Chofer, error)
	CreateChofer(c models.Chofer) (models.Chofer, error)
	UpdateChofer(c models.Chofer) error
	DeleteChofer(id int) error
}

type HorarioRepository interface {
	GetAllHorarios() ([]models.Horario, error)
	GetHorarioByID(id int) (models.Horario, error)
	CreateHorario(h models.Horario) (models.Horario, error)
	UpdateHorario(h models.Horario) error
	DeleteHorario(id int) error
	GetCarritosByHorario(id int) ([]models.CarritoHorario, error)
}

type CarritoHorarioRepository interface {
	AsignarCarritoHorario(rel models.CarritoHorarioRel) error
	DesasignarCarritoHorario(rel models.CarritoHorarioRel) error
}

type DespachoDiarioRepository interface {
	GetAllDespachos() ([]models.DespachoDiario, error)
	GetDespachoByID(id int) (models.DespachoDiario, error)
	CreateDespacho(d models.DespachoDiario) (models.DespachoDiario, error)
	UpdateDespacho(d models.DespachoDiario) error
	DeleteDespacho(id int) error
}

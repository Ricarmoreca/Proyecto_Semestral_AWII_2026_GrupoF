package storage

import (
	"errors"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

var (
	ErrNotFound = errors.New("recurso no encontrado")
	ErrConflict = errors.New("el recurso ya existe")
)

type UsuariosRepository interface {
	ListarUsuarios() []models.Usuario
	BuscarUsuarioPorID(int) (models.Usuario, bool)
	CrearUsuario(models.Usuario) models.Usuario
	ActualizarUsuario(int, models.Usuario) (models.Usuario, bool)
	BorrarUsuario(int) bool
}

type SolicitudesRepository interface {
	ListarSolicitudes() []models.Solicitud
	BuscarSolicitudPorID(id int) (models.Solicitud, bool)
	CrearSolicitud(models.Solicitud) models.Solicitud
	AsignarChofer(id int, choferId string) (models.Solicitud, bool)
	ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool)
	BorrarSolicitud(id int) bool
}

type CarritosRepository interface {
	ListarCarritos() []models.Carrito
	BuscarCarritoPorID(id int) (models.Carrito, bool)
	CrearCarrito(models.Carrito) models.Carrito
	ActualizarCarrito(id int, datos models.Carrito) (models.Carrito, bool)
	BorrarCarrito(id int) bool
}

type ChoferesRepository interface {
	ListarChoferes() []models.Chofer
	BuscarChoferPorID(id int) (models.Chofer, bool)
	CrearChofer(models.Chofer) models.Chofer
	ActualizarChofer(id int, datos models.Chofer) (models.Chofer, bool)
	BorrarChofer(id int) bool
}

type MantenimientosRepository interface {
	ListarMantenimientos() []models.Mantenimiento
	BuscarMantenimientoPorID(id int) (models.Mantenimiento, bool)
	CrearMantenimiento(models.Mantenimiento) models.Mantenimiento
	ActualizarMantenimiento(id int, datos models.Mantenimiento) (models.Mantenimiento, bool)
	BorrarMantenimiento(id int) bool
}

type CarritoHorarioRepository interface {
	AsignarCarritoHorario(carritoID int, horarioID int) (models.CarritoHorario, bool)
	DeasignarCarritoHorario(carritoID int, horarioID int) bool
}

type DespachoDiarioRepository interface {
	ListarDespachosDiarios() []models.DespachoDiario
	BuscarDespachoDiarioPorID(id int) (models.DespachoDiario, bool)
	CrearDespachoDiario(models.DespachoDiario) models.DespachoDiario
	ActualizarDespachoDiario(id int, datos models.DespachoDiario) (models.DespachoDiario, bool)
	BorrarDespachoDiario(id int) bool
}

type HorariosRepository interface {
	ListarHorarios() []models.Horario
	BuscarHorarioPorID(id int) (models.Horario, bool)
	CrearHorario(models.Horario) models.Horario
	ActualizarHorario(id int, datos models.Horario) (models.Horario, bool)
	BorrarHorario(id int) bool
}

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type Almacen interface {
	UsuariosRepository
	SolicitudesRepository
	CarritosRepository
	ChoferesRepository
	MantenimientosRepository
	CarritoHorarioRepository
	DespachoDiarioRepository
	HorariosRepository
}

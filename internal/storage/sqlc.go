package storage

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// ActualizarCarrito implements [Almacen].
func (a *AlmacenSQLC) ActualizarCarrito(id int, datos models.Carrito) (models.Carrito, bool) {
	panic("unimplemented")
}

// ActualizarChofer implements [Almacen].
func (a *AlmacenSQLC) ActualizarChofer(id int, datos models.Chofer) (models.Chofer, bool) {
	panic("unimplemented")
}

// ActualizarDespachoDiario implements [Almacen].
func (a *AlmacenSQLC) ActualizarDespachoDiario(id int, datos models.DespachoDiario) (models.DespachoDiario, bool) {
	panic("unimplemented")
}

// AsignarCarritoHorario implements [Almacen].
func (a *AlmacenSQLC) AsignarCarritoHorario(carritoID int, horarioID int) (models.CarritoHorario, bool) {
	panic("unimplemented")
}

// BorrarCarrito implements [Almacen].
func (a *AlmacenSQLC) BorrarCarrito(id int) bool {
	panic("unimplemented")
}

// BorrarChofer implements [Almacen].
func (a *AlmacenSQLC) BorrarChofer(id int) bool {
	panic("unimplemented")
}

// BorrarDespachoDiario implements [Almacen].
func (a *AlmacenSQLC) BorrarDespachoDiario(id int) bool {
	panic("unimplemented")
}

// BuscarCarritoPorID implements [Almacen].
func (a *AlmacenSQLC) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	panic("unimplemented")
}

// BuscarChoferPorID implements [Almacen].
func (a *AlmacenSQLC) BuscarChoferPorID(id int) (models.Chofer, bool) {
	panic("unimplemented")
}

// BuscarDespachoDiarioPorID implements [Almacen].
func (a *AlmacenSQLC) BuscarDespachoDiarioPorID(id int) (models.DespachoDiario, bool) {
	panic("unimplemented")
}

// CrearCarrito implements [Almacen].
func (a *AlmacenSQLC) CrearCarrito(models.Carrito) models.Carrito {
	panic("unimplemented")
}

// CrearChofer implements [Almacen].
func (a *AlmacenSQLC) CrearChofer(models.Chofer) models.Chofer {
	panic("unimplemented")
}

// CrearDespachoDiario implements [Almacen].
func (a *AlmacenSQLC) CrearDespachoDiario(models.DespachoDiario) models.DespachoDiario {
	panic("unimplemented")
}

// DeasignarCarritoHorario implements [Almacen].
func (a *AlmacenSQLC) DeasignarCarritoHorario(carritoID int, horarioID int) bool {
	panic("unimplemented")
}

// ListarCarritos implements [Almacen].
func (a *AlmacenSQLC) ListarCarritos() []models.Carrito {
	panic("unimplemented")
}

// ListarChoferes implements [Almacen].
func (a *AlmacenSQLC) ListarChoferes() []models.Chofer {
	panic("unimplemented")
}

// ListarDespachosDiarios implements [Almacen].
func (a *AlmacenSQLC) ListarDespachosDiarios() []models.DespachoDiario {
	panic("unimplemented")
}

// ListarHorarios implements [Almacen].
func (a *AlmacenSQLC) ListarHorarios() []models.Horario {
	return nil
}

// BuscarHorarioPorID implements [Almacen].
func (a *AlmacenSQLC) BuscarHorarioPorID(id int) (models.Horario, bool) {
	return models.Horario{}, false
}

// CrearHorario implements [Almacen].
func (a *AlmacenSQLC) CrearHorario(models.Horario) models.Horario {
	return models.Horario{}
}

// ActualizarHorario implements [Almacen].
func (a *AlmacenSQLC) ActualizarHorario(id int, datos models.Horario) (models.Horario, bool) {
	return models.Horario{}, false
}

// BorrarHorario implements [Almacen].
func (a *AlmacenSQLC) BorrarHorario(id int) bool {
	return false
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

func aUsuariosDominio(u sqlcdb.Usuario) models.Usuario {
	return models.Usuario{
		ID:        int(u.ID),
		Nombre:    u.Nombre,
		Rol:       u.Rol,
		Matricula: u.Matricula,
	}
}

func nullStringToPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func aSolicitudesDominio(s sqlcdb.Solicitude) models.Solicitud {
	pasajeroID, _ := strconv.Atoi(s.Pasajero)
	creadoEn := time.Time{}
	if s.Creadoen.Valid {
		creadoEn = s.Creadoen.Time
	}
	return models.Solicitud{
		ID:       int(s.ID),
		Pasajero: pasajeroID,
		Chofer:   nullStringToPtr(s.Chofer),
		Origen:   s.Origen,
		Destino:  s.Destino,
		Estado:   s.Estado,
		CreadoEn: creadoEn,
	}
}

func (a *AlmacenSQLC) ListarUsuarios() []models.Usuario {
	filas, err := a.q.ListarUsuario(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Usuario, 0, len(filas))
	for _, f := range filas {
		out = append(out, aUsuariosDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	f, err := a.q.BuscarUsuarioPorID(context.Background(), int64(id))
	if err != nil {
		return models.Usuario{}, false
	}
	return aUsuariosDominio(f), true
}

func (a *AlmacenSQLC) CrearUsuario(u models.Usuario) models.Usuario {
	f, err := a.q.CrearUsuario(context.Background(), sqlcdb.CrearUsuarioParams{
		Nombre:    u.Nombre,
		Rol:       u.Rol,
		Matricula: u.Matricula,
	})
	if err != nil {
		return models.Usuario{}
	}
	return aUsuariosDominio(f)
}

func (a *AlmacenSQLC) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	f, err := a.q.ActualizarUsuario(context.Background(), sqlcdb.ActualizarUsuarioParams{
		Nombre:    datos.Nombre,
		Rol:       datos.Rol,
		Matricula: datos.Matricula,
		ID:        int64(id),
	})
	if err != nil {
		return models.Usuario{}, false
	}
	return aUsuariosDominio(f), true
}

func (a *AlmacenSQLC) BorrarUsuario(id int) bool {
	filas, err := a.q.BorrarUsuario(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

func (a *AlmacenSQLC) ListarSolicitudes() []models.Solicitud {
	filas, err := a.q.ListarSolicitudes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Solicitud, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSolicitudesDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	f, err := a.q.BuscarSolicitudPorID(context.Background(), int64(id))
	if err != nil {
		return models.Solicitud{}, false
	}
	return aSolicitudesDominio(f), true
}

func (a *AlmacenSQLC) CrearSolicitud(s models.Solicitud) models.Solicitud {
	f, err := a.q.CrearSolicitud(context.Background(), sqlcdb.CrearSolicitudParams{
		Pasajero: strconv.Itoa(s.Pasajero),
		Origen:   s.Origen,
		Destino:  s.Destino,
	})
	if err != nil {
		return models.Solicitud{}
	}
	return aSolicitudesDominio(f)
}

func (a *AlmacenSQLC) ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool) {
	chofer := sql.NullString{Valid: false}
	if datos.Chofer != nil {
		chofer = sql.NullString{String: *datos.Chofer, Valid: true}
	}

	f, err := a.q.ActualizarSolicitud(context.Background(), sqlcdb.ActualizarSolicitudParams{
		ID:     int64(id),
		Estado: datos.Estado,
		Chofer: chofer,
	})
	if err != nil {
		return models.Solicitud{}, false
	}
	return aSolicitudesDominio(f), true
}

func (a *AlmacenSQLC) BorrarSolicitud(id int) bool {
	filas, err := a.q.BorrarSolicitud(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

func (a *AlmacenSQLC) AsignarChofer(id int, choferId string) (models.Solicitud, bool) {
	s, err := a.q.AsignarChofer(context.Background(), sqlcdb.AsignarChoferParams{
		ID:     int64(id),
		Chofer: sql.NullString{String: choferId, Valid: true},
	})
	if err != nil {
		return models.Solicitud{}, false
	}

	return aSolicitudesDominio(s), true
}

// ListarMantenimientos implements [Almacen].
func (a *AlmacenSQLC) ListarMantenimientos() []models.Mantenimiento {
	return []models.Mantenimiento{}
}

// BuscarMantenimientoPorID implements [Almacen].
func (a *AlmacenSQLC) BuscarMantenimientoPorID(id int) (models.Mantenimiento, bool) {
	return models.Mantenimiento{}, false
}

// CrearMantenimiento implements [Almacen].
func (a *AlmacenSQLC) CrearMantenimiento(m models.Mantenimiento) models.Mantenimiento {
	return models.Mantenimiento{}
}

// ActualizarMantenimiento implements [Almacen].
func (a *AlmacenSQLC) ActualizarMantenimiento(id int, datos models.Mantenimiento) (models.Mantenimiento, bool) {
	return models.Mantenimiento{}, false
}

// BorrarMantenimiento implements [Almacen].
func (a *AlmacenSQLC) BorrarMantenimiento(id int) bool {
	return false
}

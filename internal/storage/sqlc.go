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

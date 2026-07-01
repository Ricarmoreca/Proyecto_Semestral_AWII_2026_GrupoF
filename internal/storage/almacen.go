package storage

import "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"

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

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type Almacen interface {
	UsuariosRepository
	SolicitudesRepository
}

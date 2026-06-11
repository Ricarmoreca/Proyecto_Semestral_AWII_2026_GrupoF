package storage

import "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"

type Almacen interface {
	ListarUsuarios() []models.Usuario
	BuscarUsuarioPorID(int) (models.Usuario, bool)
	CrearUsuario(models.Usuario) models.Usuario
	ActualizarUsuario(int, models.Usuario) (models.Usuario, bool)
	BorrarUsuario(int) bool

	ListarSolicitudes() []models.Solicitud
	BuscarSolicitudPorID(string) (models.Solicitud, bool)
	CrearSolicitud(models.Solicitud) models.Solicitud
	ActualizarSolicitud(string, models.Solicitud) (models.Solicitud, bool)
	BorrarSolicitud(string) bool
}

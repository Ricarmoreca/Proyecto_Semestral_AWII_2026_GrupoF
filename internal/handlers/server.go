package handlers

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type Server struct {
	Usuarios    *service.UsuarioService
	Solicitudes *service.SolicitudService
	Auth        *service.AuthService
	Almacen     storage.Almacen
}

func NewServer(usuarios *service.UsuarioService, solicitudes *service.SolicitudService, auth *service.AuthService, almacen storage.Almacen) *Server {
	return &Server{
		Usuarios:    usuarios,
		Solicitudes: solicitudes,
		Auth:        auth,
		Almacen:     almacen,
	}
}

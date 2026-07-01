package handlers

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
)

type Server struct {
	Usuarios    *service.UsuarioService
	Solicitudes *service.SolicitudService
	Auth        *service.AuthService
}

func NewServer(usuarios *service.UsuarioService, solicitudes *service.SolicitudService, auth *service.AuthService) *Server {
	return &Server{
		Usuarios:    usuarios,
		Solicitudes: solicitudes,
		Auth:        auth,
	}
}

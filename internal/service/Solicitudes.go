package service

import (
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type SolicitudService struct {
	repo storage.SolicitudesRepository
}

func NewSolicitudService(repo storage.SolicitudesRepository) *SolicitudService {
	return &SolicitudService{repo: repo}
}

func (s *SolicitudService) Listar() []models.Solicitud {
	return s.repo.ListarSolicitudes()
}

func (s *SolicitudService) Obtener(id int) (models.Solicitud, error) {
	solicitud, ok := s.repo.BuscarSolicitudPorID(id)
	if !ok {
		return models.Solicitud{}, ErrSolicitudNoEncontrada
	}
	return solicitud, nil
}

func (s *SolicitudService) Crear(solicitud models.Solicitud) (models.Solicitud, error) {
	if err := validacionSolicitud(solicitud); err != nil {
		return models.Solicitud{}, err
	}
	return s.repo.CrearSolicitud(solicitud), nil
}

func (s *SolicitudService) Actualizar(id int, solicitud models.Solicitud) (models.Solicitud, error) {
	if err := validacionActualizacionSolicitud(solicitud); err != nil {
		return models.Solicitud{}, err
	}
	actualizado, ok := s.repo.ActualizarSolicitud(id, solicitud)
	if !ok {
		return models.Solicitud{}, ErrSolicitudNoEncontrada
	}
	return actualizado, nil
}

func (s *SolicitudService) Borrar(id int) error {
	if !s.repo.BorrarSolicitud(id) {
		return ErrSolicitudNoEncontrada
	}
	return nil
}

func validacionSolicitud(solicitud models.Solicitud) error {
	if solicitud.Pasajero <= 0 || strings.TrimSpace(solicitud.Origen) == "" || strings.TrimSpace(solicitud.Destino) == "" {
		return ErrCamposObligatorios
	}
	return nil
}

func validacionActualizacionSolicitud(solicitud models.Solicitud) error {
	if strings.TrimSpace(solicitud.Estado) == "" {
		return ErrCamposObligatorios
	}
	return nil
}

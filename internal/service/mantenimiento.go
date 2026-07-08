package service

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type MantenimientoService struct {
	repo storage.MantenimientosRepository
}

func NewMantenimientoService(repo storage.MantenimientosRepository) *MantenimientoService {
	return &MantenimientoService{repo: repo}
}

func (s *MantenimientoService) Listar() []models.Mantenimiento {
	return s.repo.ListarMantenimientos()
}

func (s *MantenimientoService) Obtener(id int) (models.Mantenimiento, error) {
	mantenimiento, ok := s.repo.BuscarMantenimientoPorID(id)
	if !ok {
		return models.Mantenimiento{}, ErrMantenimientoNoEncontrado
	}
	return mantenimiento, nil
}

func (s *MantenimientoService) Crear(mantenimiento models.Mantenimiento) (models.Mantenimiento, error) {
	if err := mantenimiento.Validate(); err != nil {
		return models.Mantenimiento{}, err
	}
	return s.repo.CrearMantenimiento(mantenimiento), nil
}

func (s *MantenimientoService) Actualizar(id int, datos models.Mantenimiento) (models.Mantenimiento, error) {
	if err := datos.Validate(); err != nil {
		return models.Mantenimiento{}, err
	}
	actualizado, ok := s.repo.ActualizarMantenimiento(id, datos)
	if !ok {
		return models.Mantenimiento{}, ErrMantenimientoNoEncontrado
	}
	return actualizado, nil
}

func (s *MantenimientoService) Borrar(id int) error {
	if !s.repo.BorrarMantenimiento(id) {
		return ErrMantenimientoNoEncontrado
	}
	return nil
}

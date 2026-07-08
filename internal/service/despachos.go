package service

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type DespachoService struct {
	repo storage.DespachoDiarioRepository
}

func NewDespachoService(repo storage.DespachoDiarioRepository) *DespachoService {
	return &DespachoService{repo: repo}
}

func (s *DespachoService) Listar() []models.DespachoDiario {
	return s.repo.ListarDespachosDiarios()
}

func (s *DespachoService) Obtener(id int) (models.DespachoDiario, error) {
	despacho, ok := s.repo.BuscarDespachoDiarioPorID(id)
	if !ok {
		return models.DespachoDiario{}, ErrNotFound
	}
	return despacho, nil
}

func (s *DespachoService) Crear(despacho models.DespachoDiario) (models.DespachoDiario, error) {
	if err := despacho.Validate(); err != nil {
		return models.DespachoDiario{}, err
	}
	return s.repo.CrearDespachoDiario(despacho), nil
}

func (s *DespachoService) Actualizar(id int, despacho models.DespachoDiario) (models.DespachoDiario, error) {
	if err := despacho.Validate(); err != nil {
		return models.DespachoDiario{}, err
	}
	actualizado, ok := s.repo.ActualizarDespachoDiario(id, despacho)
	if !ok {
		return models.DespachoDiario{}, ErrNotFound
	}
	return actualizado, nil
}

func (s *DespachoService) Borrar(id int) error {
	if !s.repo.BorrarDespachoDiario(id) {
		return ErrNotFound
	}
	return nil
}

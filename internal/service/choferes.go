package service

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type ChoferService struct {
	repo storage.ChoferesRepository
}

func NewChoferService(repo storage.ChoferesRepository) *ChoferService {
	return &ChoferService{repo: repo}
}

func (s *ChoferService) Listar() []models.Chofer {
	return s.repo.ListarChoferes()
}

func (s *ChoferService) Obtener(id int) (models.Chofer, error) {
	chofer, ok := s.repo.BuscarChoferPorID(id)
	if !ok {
		return models.Chofer{}, ErrChoferNoEncontrado
	}
	return chofer, nil
}

func (s *ChoferService) Crear(chofer models.Chofer) (models.Chofer, error) {
	if err := chofer.Validate(); err != nil {
		return models.Chofer{}, err
	}
	return s.repo.CrearChofer(chofer), nil
}

func (s *ChoferService) Actualizar(id int, datos models.Chofer) (models.Chofer, error) {
	if err := datos.Validate(); err != nil {
		return models.Chofer{}, err
	}
	actualizado, ok := s.repo.ActualizarChofer(id, datos)
	if !ok {
		return models.Chofer{}, ErrChoferNoEncontrado
	}
	return actualizado, nil
}

func (s *ChoferService) Borrar(id int) error {
	if !s.repo.BorrarChofer(id) {
		return ErrChoferNoEncontrado
	}
	return nil
}

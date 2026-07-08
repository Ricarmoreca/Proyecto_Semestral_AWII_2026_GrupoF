package service

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type HorarioService struct {
	repo storage.HorariosRepository
}

func NewHorarioService(repo storage.HorariosRepository) *HorarioService {
	return &HorarioService{repo: repo}
}

func (s *HorarioService) Listar() []models.Horario {
	return s.repo.ListarHorarios()
}

func (s *HorarioService) Obtener(id int) (models.Horario, error) {
	horario, ok := s.repo.BuscarHorarioPorID(id)
	if !ok {
		return models.Horario{}, ErrNotFound
	}
	return horario, nil
}

func (s *HorarioService) Crear(horario models.Horario) (models.Horario, error) {
	if err := horario.Validate(); err != nil {
		return models.Horario{}, err
	}
	return s.repo.CrearHorario(horario), nil
}

func (s *HorarioService) Actualizar(id int, horario models.Horario) (models.Horario, error) {
	if err := horario.Validate(); err != nil {
		return models.Horario{}, err
	}
	actualizado, ok := s.repo.ActualizarHorario(id, horario)
	if !ok {
		return models.Horario{}, ErrNotFound
	}
	return actualizado, nil
}

func (s *HorarioService) Borrar(id int) error {
	if !s.repo.BorrarHorario(id) {
		return ErrNotFound
	}
	return nil
}

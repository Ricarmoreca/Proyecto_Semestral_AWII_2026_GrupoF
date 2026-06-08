package service

import (
	"errors"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type ChoferService struct {
	repo storage.ChoferRepository
}

func NewChoferService(repo storage.ChoferRepository) *ChoferService {
	return &ChoferService{repo: repo}
}

func (s *ChoferService) CrearChofer(chofer models.Chofer) (models.Chofer, error) {
	if err := chofer.Validate(); err != nil {
		return models.Chofer{}, err
	}
	if _, err := s.repo.GetChoferByLicencia(chofer.Licencia); err == nil {
		return models.Chofer{}, errors.New("ya existe un chofer con la licencia " + chofer.Licencia)
	}
	return s.repo.CreateChofer(chofer)
}

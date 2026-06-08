package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

func (a *AlmacenSQLite) GetAllChoferes() ([]models.Chofer, error) {
	var choferes []models.Chofer
	if err := a.db.Find(&choferes).Error; err != nil {
		return []models.Chofer{}, nil
	}
	if choferes == nil {
		return []models.Chofer{}, nil
	}
	return choferes, nil
}

func (a *AlmacenSQLite) GetChoferByID(id int) (models.Chofer, error) {
	var c models.Chofer
	if err := a.db.First(&c, id).Error; err != nil {
		if isNotFoundError(err) {
			return models.Chofer{}, ErrNotFound
		}
		return models.Chofer{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) GetChoferByLicencia(licencia string) (models.Chofer, error) {
	var c models.Chofer
	if err := a.db.Where("licencia = ?", licencia).First(&c).Error; err != nil {
		if isNotFoundError(err) {
			return models.Chofer{}, ErrNotFound
		}
		return models.Chofer{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) CreateChofer(c models.Chofer) (models.Chofer, error) {
	if err := a.db.Create(&c).Error; err != nil {
		if isConstraintError(err) {
			return models.Chofer{}, ErrConflict
		}
		return models.Chofer{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) UpdateChofer(c models.Chofer) error {
	result := a.db.Model(&models.Chofer{}).Where("id_chofer = ?", c.ID).Updates(map[string]interface{}{
		"nombre_chofer": c.Nombre,
		"licencia":      c.Licencia,
		"celular":       c.Celular,
		"estado_chofer": c.Estado,
	})
	if result.Error != nil {
		if isConstraintError(result.Error) {
			return ErrConflict
		}
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) DeleteChofer(id int) error {
	result := a.db.Delete(&models.Chofer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

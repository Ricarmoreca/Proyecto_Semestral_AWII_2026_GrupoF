package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

func (a *AlmacenSQLite) GetAllDespachos() ([]models.DespachoDiario, error) {
	var despachos []models.DespachoDiario
	if err := a.db.Find(&despachos).Error; err != nil {
		return []models.DespachoDiario{}, nil
	}
	if despachos == nil {
		return []models.DespachoDiario{}, nil
	}
	return despachos, nil
}

func (a *AlmacenSQLite) GetDespachoByID(id int) (models.DespachoDiario, error) {
	var d models.DespachoDiario
	if err := a.db.First(&d, id).Error; err != nil {
		if isNotFoundError(err) {
			return models.DespachoDiario{}, ErrNotFound
		}
		return models.DespachoDiario{}, err
	}
	return d, nil
}

func (a *AlmacenSQLite) CreateDespacho(d models.DespachoDiario) (models.DespachoDiario, error) {
	if err := a.ensureCarritoExists(d.NumeroCarrito); err != nil {
		return models.DespachoDiario{}, err
	}
	if err := a.ensureHorarioExists(d.IDHorario); err != nil {
		return models.DespachoDiario{}, err
	}
	if err := a.ensureChoferExists(d.IDChofer); err != nil {
		return models.DespachoDiario{}, err
	}
	if err := a.db.Create(&d).Error; err != nil {
		return models.DespachoDiario{}, err
	}
	return d, nil
}

func (a *AlmacenSQLite) UpdateDespacho(d models.DespachoDiario) error {
	result := a.db.Model(&models.DespachoDiario{}).Where("id_despacho = ?", d.ID).Updates(map[string]interface{}{
		"fecha":              d.Fecha,
		"numero_carrito":     d.NumeroCarrito,
		"id_horario":         d.IDHorario,
		"id_chofer":          d.IDChofer,
		"pasajeros_actuales": d.PasajerosActuales,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) DeleteDespacho(id int) error {
	result := a.db.Delete(&models.DespachoDiario{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

func (a *AlmacenSQLite) GetAllHorarios() ([]models.Horario, error) {
	var horarios []models.Horario
	if err := a.db.Find(&horarios).Error; err != nil {
		return []models.Horario{}, nil
	}
	if horarios == nil {
		return []models.Horario{}, nil
	}
	return horarios, nil
}

func (a *AlmacenSQLite) GetHorarioByID(id int) (models.Horario, error) {
	var h models.Horario
	if err := a.db.First(&h, id).Error; err != nil {
		if isNotFoundError(err) {
			return models.Horario{}, ErrNotFound
		}
		return models.Horario{}, err
	}
	return h, nil
}

func (a *AlmacenSQLite) CreateHorario(h models.Horario) (models.Horario, error) {
	if err := a.db.Create(&h).Error; err != nil {
		return models.Horario{}, err
	}
	return h, nil
}

func (a *AlmacenSQLite) UpdateHorario(h models.Horario) error {
	result := a.db.Model(&models.Horario{}).Where("id_horario = ?", h.ID).Updates(map[string]interface{}{
		"turno":       h.Turno,
		"hora_inicio": h.HoraInicio,
		"hora_fin":    h.HoraFin,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) DeleteHorario(id int) error {
	result := a.db.Delete(&models.Horario{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) GetCarritosByHorario(id int) ([]models.CarritoHorario, error) {
	var exists int64
	a.db.Model(&models.Horario{}).Where("id_horario = ?", id).Count(&exists)
	if exists == 0 {
		return nil, ErrNotFound
	}

	var results []models.CarritoHorario
	query := `SELECT ch.numero_carrito, ch.id_horario, ch.hora_asignacion,
		c.estado_carrito, c.capacidad_pasajeros, c.color
		FROM carrito_horario ch
		JOIN carritos c ON c.numero_carrito = ch.numero_carrito AND c.deleted_at IS NULL
		WHERE ch.id_horario = ?
		ORDER BY c.numero_carrito`
	if err := a.db.Raw(query, id).Scan(&results).Error; err != nil {
		return nil, err
	}
	if results == nil {
		return []models.CarritoHorario{}, nil
	}
	return results, nil
}

func (a *AlmacenSQLite) AsignarCarritoHorario(rel models.CarritoHorarioRel) error {
	if err := a.ensureCarritoExists(rel.NumeroCarrito); err != nil {
		return err
	}
	if err := a.ensureHorarioExists(rel.IDHorario); err != nil {
		return err
	}
	if err := a.ensureRelacionNoExiste(rel.NumeroCarrito, rel.IDHorario); err != nil {
		return err
	}
	if err := a.db.Create(&rel).Error; err != nil {
		if isConstraintError(err) {
			return ErrConflict
		}
		return err
	}
	return nil
}

func (a *AlmacenSQLite) DesasignarCarritoHorario(rel models.CarritoHorarioRel) error {
	result := a.db.Where("numero_carrito = ? AND id_horario = ?", rel.NumeroCarrito, rel.IDHorario).Delete(&models.CarritoHorarioRel{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

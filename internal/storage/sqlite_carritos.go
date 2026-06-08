package storage

import (
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"gorm.io/gorm"
)

func (a *AlmacenSQLite) GetAllCarritos() ([]models.Carrito, error) {
	var carritos []models.Carrito
	if err := a.db.Find(&carritos).Error; err != nil {
		return []models.Carrito{}, nil
	}
	if carritos == nil {
		return []models.Carrito{}, nil
	}
	return carritos, nil
}

func (a *AlmacenSQLite) GetCarritoByNumero(numero int) (models.Carrito, error) {
	var c models.Carrito
	if err := a.db.First(&c, numero).Error; err != nil {
		if isNotFoundError(err) {
			return models.Carrito{}, ErrNotFound
		}
		return models.Carrito{}, err
	}
	return c, nil
}

func (a *AlmacenSQLite) CreateCarrito(c models.Carrito) error {
	if err := a.db.Create(&c).Error; err != nil {
		if isConstraintError(err) {
			return ErrConflict
		}
		return err
	}
	return nil
}

func (a *AlmacenSQLite) UpdateCarrito(c models.Carrito) error {
	result := a.db.Model(&models.Carrito{}).Where("numero_carrito = ?", c.Numero).Updates(map[string]interface{}{
		"estado_carrito":      c.Estado,
		"capacidad_pasajeros": c.CapacidadPasajeros,
		"color":               c.Color,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) DeleteCarrito(numero int) error {
	result := a.db.Delete(&models.Carrito{}, numero)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) GetHorariosByCarrito(numero int) ([]models.CarritoHorario, error) {
	var exists int64
	a.db.Unscoped().Model(&models.Carrito{}).Where("numero_carrito = ?", numero).Count(&exists)
	if exists == 0 {
		return nil, ErrNotFound
	}

	var results []models.CarritoHorario
	query := `SELECT ch.numero_carrito, ch.id_horario, ch.hora_asignacion,
		h.turno, h.hora_inicio, h.hora_fin
		FROM carrito_horario ch
		JOIN horarios h ON h.id_horario = ch.id_horario
		JOIN carritos c ON c.numero_carrito = ch.numero_carrito AND c.deleted_at IS NULL
		WHERE ch.numero_carrito = ?
		ORDER BY h.turno, h.hora_inicio`
	if err := a.db.Raw(query, numero).Scan(&results).Error; err != nil {
		return nil, err
	}
	if results == nil {
		return []models.CarritoHorario{}, nil
	}
	return results, nil
}

func isConstraintError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "UNIQUE constraint") || strings.Contains(msg, "PRIMARY KEY")
}

func isNotFoundError(err error) bool {
	return err == gorm.ErrRecordNotFound
}

func (a *AlmacenSQLite) ensureCarritoExists(numero int) error {
	var count int64
	a.db.Model(&models.Carrito{}).Where("numero_carrito = ?", numero).Count(&count)
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) ensureHorarioExists(id int) error {
	var count int64
	a.db.Model(&models.Horario{}).Where("id_horario = ?", id).Count(&count)
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) ensureChoferExists(id int) error {
	var count int64
	a.db.Model(&models.Chofer{}).Where("id_chofer = ?", id).Count(&count)
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *AlmacenSQLite) ensureRelacionNoExiste(numeroCarrito, idHorario int) error {
	var count int64
	a.db.Model(&models.CarritoHorarioRel{}).Where("numero_carrito = ? AND id_horario = ?", numeroCarrito, idHorario).Count(&count)
	if count > 0 {
		return ErrConflict
	}
	return nil
}

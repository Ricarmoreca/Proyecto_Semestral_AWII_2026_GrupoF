package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func (a *AlmacenSQLite) DB() *gorm.DB {
	return a.db
}

func NewSQLiteStorage(path string) (*AlmacenSQLite, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Carrito{}, &models.Chofer{}, &models.Horario{}, &models.CarritoHorarioRel{}, &models.DespachoDiario{}); err != nil {
		return nil, err
	}
	return NuevoAlmacenSQLite(db), nil
}

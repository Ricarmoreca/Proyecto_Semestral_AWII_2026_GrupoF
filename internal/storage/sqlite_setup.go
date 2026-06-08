package storage

import (
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteAlmacen(path string) (*AlmacenSQLite, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.Exec("PRAGMA foreign_keys = ON")

	store := NuevoAlmacenSQLite(db)
	if err := store.Migrate(); err != nil {
		return nil, err
	}
	store.Seed()
	return store, nil
}

func (a *AlmacenSQLite) Migrate() error {
	return a.db.AutoMigrate(
		&models.Carrito{},
		&models.Chofer{},
		&models.Horario{},
		&models.CarritoHorarioRel{},
		&models.DespachoDiario{},
	)
}

func (a *AlmacenSQLite) Seed() {
	if err := a.seedHorarios(); err != nil {
		return
	}
	a.seedCarritos()
	a.seedChoferes()
	a.seedDespachos()
	a.seedCarritoHorarios()
}

func (a *AlmacenSQLite) seedHorarios() error {
	var count int64
	a.db.Model(&models.Horario{}).Count(&count)
	if count > 0 {
		return nil
	}
	horarios := []models.Horario{
		{Turno: models.TurnoMatutina, HoraInicio: "07:00", HoraFin: "12:00"},
		{Turno: models.TurnoVespertina, HoraInicio: "13:00", HoraFin: "18:00"},
	}
	return a.db.Create(&horarios).Error
}

func (a *AlmacenSQLite) seedCarritos() {
	var count int64
	a.db.Unscoped().Model(&models.Carrito{}).Count(&count)
	if count > 0 {
		return
	}
	carritos := []models.Carrito{
		{Numero: 101, Estado: models.EstadoCarritoDisponible, CapacidadPasajeros: 15, Color: "Blanco"},
		{Numero: 102, Estado: models.EstadoCarritoEnViaje, CapacidadPasajeros: 20, Color: "Amarillo"},
		{Numero: 103, Estado: models.EstadoCarritoMantenimiento, CapacidadPasajeros: 12, Color: "Azul"},
		{Numero: 104, Estado: models.EstadoCarritoDisponible, CapacidadPasajeros: 18, Color: "Verde"},
	}
	a.db.Create(&carritos)
}

func (a *AlmacenSQLite) seedChoferes() {
	var count int64
	a.db.Model(&models.Chofer{}).Count(&count)
	if count > 0 {
		return
	}
	choferes := []models.Chofer{
		{Nombre: "Carlos Mena", Licencia: "A123456", Celular: "0991112223", Estado: models.EstadoChoferDisponible},
		{Nombre: "Luis Zambrano", Licencia: "B987654", Celular: "0992223334", Estado: models.EstadoChoferEnRuta},
		{Nombre: "Ana Vera", Licencia: "C456789", Celular: "0993334445", Estado: models.EstadoChoferDescanso},
	}
	a.db.Create(&choferes)
}

func (a *AlmacenSQLite) seedDespachos() {
	var count int64
	a.db.Model(&models.DespachoDiario{}).Count(&count)
	if count > 0 {
		return
	}
	despachos := []models.DespachoDiario{
		{Fecha: "2026-06-22", NumeroCarrito: 101, IDHorario: 1, IDChofer: 1, PasajerosActuales: 12},
		{Fecha: "2026-06-22", NumeroCarrito: 102, IDHorario: 2, IDChofer: 2, PasajerosActuales: 18},
	}
	a.db.Create(&despachos)
}

func (a *AlmacenSQLite) seedCarritoHorarios() {
	var count int64
	a.db.Model(&models.CarritoHorarioRel{}).Count(&count)
	if count > 0 {
		return
	}
	ahora := time.Now().Format("15:04")
	rels := []models.CarritoHorarioRel{
		{NumeroCarrito: 101, IDHorario: 1, HoraAsignacion: ahora},
		{NumeroCarrito: 102, IDHorario: 1, HoraAsignacion: ahora},
		{NumeroCarrito: 102, IDHorario: 2, HoraAsignacion: ahora},
		{NumeroCarrito: 104, IDHorario: 2, HoraAsignacion: ahora},
	}
	a.db.Create(&rels)
}

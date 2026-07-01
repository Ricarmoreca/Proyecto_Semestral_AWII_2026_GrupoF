package storage

import (
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func NewSQLiteStorage(path string) (*AlmacenSQLite, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Usuario{}, &models.Solicitud{}); err != nil {
		return nil, err
	}
	return NuevoAlmacenSQLite(db), nil
}

func (a *AlmacenSQLite) ListarUsuarios() []models.Usuario {
	var u []models.Usuario
	a.db.Find(&u)
	return u
}

func (a *AlmacenSQLite) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var u models.Usuario
	if err := a.db.First(&u, id).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

func (a *AlmacenSQLite) CrearUsuario(u models.Usuario) models.Usuario {
	a.db.Create(&u)
	return u
}

func (a *AlmacenSQLite) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	var existente models.Usuario
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Usuario{}, false
	}

	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarUsuario(id int) bool {
	res := a.db.Delete(&models.Usuario{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) ListarSolicitudes() []models.Solicitud {
	var s []models.Solicitud
	a.db.Find(&s)
	return s
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s models.Solicitud) models.Solicitud {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) AsignarChofer(id int, choferId string) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	s.Chofer = &choferId
	if err := a.db.Save(&s).Error; err != nil {
		return models.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitud(id int) bool {
	res := a.db.Delete(&models.Solicitud{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Usuario{}).Count(&n)
	if n > 0 {
		return
	}

	usuarios := []models.Usuario{
		{ID: 1, Nombre: "Ana Pérez", Rol: "estudiante", Matricula: "2026001"},
		{ID: 2, Nombre: "Carlos López", Rol: "docente", Matricula: "2026002"},
		{ID: 3, Nombre: "Luis Martínez", Rol: "estudiante", Matricula: "2026003"},
	}
	a.db.Create(&usuarios)

	chofer1 := "Manuel Rodriguez"
	chofer2 := "Albert Lopez"

	solicitudes := []models.Solicitud{
		{ID: 1, Pasajero: 1, Chofer: &chofer1, Origen: "Campus", Destino: "Biblioteca", Estado: "pendiente", CreadoEn: time.Now()},
		{ID: 2, Pasajero: 2, Chofer: &chofer2, Origen: "Biblioteca", Destino: "Campus", Estado: "aceptada", CreadoEn: time.Now()},
		{ID: 3, Pasajero: 3, Chofer: nil, Origen: "Residencia", Destino: "Facultad", Estado: "pendiente", CreadoEn: time.Now()},
	}
	a.db.Create(&solicitudes)
}

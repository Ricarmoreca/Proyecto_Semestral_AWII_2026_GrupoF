package storage

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLiteStorage struct {
	db *gorm.DB
}

func NewSQLiteStorage(dsn string) (*SQLiteStorage, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Usuario{}, &models.Solicitud{}); err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) ListarUsuarios() []models.Usuario {
	var usuarios []models.Usuario
	s.db.Find(&usuarios)
	return usuarios
}

func (s *SQLiteStorage) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var usuario models.Usuario
	err := s.db.First(&usuario, "id = ?", fmt.Sprintf("%d", id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Usuario{}, false
		}
		return models.Usuario{}, false
	}
	return usuario, true
}

func (s *SQLiteStorage) CrearUsuario(usuario models.Usuario) models.Usuario {
	if strings.TrimSpace(usuario.ID) == "" {
		var maxID int
		s.db.Raw("SELECT MAX(CAST(id AS INTEGER)) FROM usuarios").Scan(&maxID)
		usuario.ID = fmt.Sprintf("%d", maxID+1)
	}
	s.db.Create(&usuario)
	return usuario
}

func (s *SQLiteStorage) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	var usuario models.Usuario
	if err := s.db.First(&usuario, "id = ?", fmt.Sprintf("%d", id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Usuario{}, false
		}
		return models.Usuario{}, false
	}
	datos.ID = usuario.ID
	usuario.Nombre = datos.Nombre
	usuario.Rol = datos.Rol
	usuario.Matricula = datos.Matricula
	if err := s.db.Save(&usuario).Error; err != nil {
		return models.Usuario{}, false
	}
	return usuario, true
}

func (s *SQLiteStorage) BorrarUsuario(id int) bool {
	res := s.db.Delete(&models.Usuario{}, "id = ?", fmt.Sprintf("%d", id))
	return res.Error == nil && res.RowsAffected > 0
}

func (s *SQLiteStorage) ListarSolicitudes() []models.Solicitud {
	var solicitudes []models.Solicitud
	s.db.Find(&solicitudes)
	return solicitudes
}

func (s *SQLiteStorage) BuscarSolicitudPorID(id string) (models.Solicitud, bool) {
	var solicitud models.Solicitud
	err := s.db.First(&solicitud, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Solicitud{}, false
		}
		return models.Solicitud{}, false
	}
	return solicitud, true
}

func (s *SQLiteStorage) CrearSolicitud(solicitud models.Solicitud) models.Solicitud {
	var maxSeq int
	s.db.Raw("SELECT MAX(CAST(substr(id, 5) AS INTEGER)) FROM solicitudes").Scan(&maxSeq)
	solicitud.ID = fmt.Sprintf("SOL-%d", maxSeq+1)
	if solicitud.Estado == "" {
		solicitud.Estado = "pendiente"
	}
	if solicitud.CreadoEn.IsZero() {
		solicitud.CreadoEn = time.Now()
	}
	s.db.Create(&solicitud)
	return solicitud
}

func (s *SQLiteStorage) ActualizarSolicitud(id string, datos models.Solicitud) (models.Solicitud, bool) {
	var solicitud models.Solicitud
	if err := s.db.First(&solicitud, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Solicitud{}, false
		}
		return models.Solicitud{}, false
	}

	if strings.TrimSpace(datos.Estado) != "" {
		solicitud.Estado = datos.Estado
	}
	if strings.TrimSpace(datos.Chofer) != "" {
		solicitud.Chofer = datos.Chofer
	}

	if err := s.db.Save(&solicitud).Error; err != nil {
		return models.Solicitud{}, false
	}
	return solicitud, true
}

func (s *SQLiteStorage) BorrarSolicitud(id string) bool {
	res := s.db.Delete(&models.Solicitud{}, "id = ?", id)
	return res.Error == nil && res.RowsAffected > 0
}

func (s *SQLiteStorage) SembrarSiVacio() {
	var count int64
	s.db.Model(&models.Usuario{}).Count(&count)
	if count > 0 {
		return
	}
	usuarios := []models.Usuario{
		{ID: "1", Nombre: "Ana Pérez", Rol: "estudiante", Matricula: "2026001"},
		{ID: "2", Nombre: "Carlos López", Rol: "docente", Matricula: "2026002"},
	}
	s.db.CreateInBatches(&usuarios, 10)
	solicitudes := []models.Solicitud{
		{ID: "SOL-1", Pasajero: "1", Chofer: "Manuel Rodriguez", Origen: "Campus", Destino: "Biblioteca", Estado: "pendiente", CreadoEn: time.Now()},
		{ID: "SOL-2", Pasajero: "2", Chofer: "Albert Lopez", Origen: "Biblioteca", Destino: "Campus", Estado: "aceptada", CreadoEn: time.Now()},
	}
	s.db.CreateInBatches(&solicitudes, 10)
}

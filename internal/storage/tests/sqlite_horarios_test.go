package storage_test

import (
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupHorarioDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Horario{}))
	return db
}

func TestHorarioSQLite_CrearYObtener(t *testing.T) {
	db := setupHorarioDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	horario := models.Horario{Turno: "Matutina", HoraInicio: "06:00", HoraFin: "11:00"}
	creado, err := repo.CreateHorario(horario)

	require.NoError(t, err)
	assert.Greater(t, creado.ID, 0)
	assert.Equal(t, "Matutina", string(creado.Turno))

	obtenido, err := repo.GetHorarioByID(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, "06:00", obtenido.HoraInicio)
}

func TestHorarioSQLite_GetHorarioByID_NoEncontrado(t *testing.T) {
	db := setupHorarioDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	_, err := repo.GetHorarioByID(999)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestHorarioSQLite_ListarVacio(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_horario_list?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Horario{}))
	repo := storage.NuevoAlmacenSQLite(db)

	lista, err := repo.GetAllHorarios()
	require.NoError(t, err)
	assert.Empty(t, lista)
}

func TestHorarioSQLite_Actualizar(t *testing.T) {
	db := setupHorarioDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateHorario(models.Horario{Turno: "Matutina", HoraInicio: "06:00", HoraFin: "11:00"})
	creado.HoraFin = "12:00"
	err := repo.UpdateHorario(creado)
	require.NoError(t, err)

	obtenido, _ := repo.GetHorarioByID(creado.ID)
	assert.Equal(t, "12:00", obtenido.HoraFin)
}

func TestHorarioSQLite_Eliminar(t *testing.T) {
	db := setupHorarioDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateHorario(models.Horario{Turno: "Vespertina", HoraInicio: "13:00", HoraFin: "18:00"})
	err := repo.DeleteHorario(creado.ID)
	require.NoError(t, err)

	_, err = repo.GetHorarioByID(creado.ID)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

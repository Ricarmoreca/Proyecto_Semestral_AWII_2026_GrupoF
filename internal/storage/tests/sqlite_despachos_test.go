package storage_test

import (
	"testing"
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupDespachoDB(t *testing.T, name string) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:"+name+"?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Carrito{}, &models.Horario{}, &models.Chofer{}, &models.DespachoDiario{}))

	db.Create(&models.Carrito{Numero: 10, Estado: "Disponible", CapacidadPasajeros: 15, Color: "Azul"})
	db.Create(&models.Horario{Turno: "Matutina", HoraInicio: "06:00", HoraFin: "11:00"})
	db.Create(&models.Chofer{Nombre: "Test Chofer", Licencia: "LCHOFER1", Celular: "0990000111", Estado: "Disponible"})

	return db
}

func TestDespachoSQLite_CrearYObtener(t *testing.T) {
	db := setupDespachoDB(t, "mem_desp_01")
	repo := storage.NuevoAlmacenSQLite(db)

	despacho := models.DespachoDiario{
		Fecha:             time.Now().Format("2006-01-02"),
		NumeroCarrito:     10,
		IDHorario:         1,
		IDChofer:          1,
		PasajerosActuales: 8,
	}
	creado, err := repo.CreateDespacho(despacho)

	require.NoError(t, err)
	assert.Greater(t, creado.ID, 0)
	assert.Equal(t, 8, creado.PasajerosActuales)

	obtenido, err := repo.GetDespachoByID(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, 10, obtenido.NumeroCarrito)
}

func TestDespachoSQLite_GetDespachoByID_NoEncontrado(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_desp_02?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.DespachoDiario{}))
	repo := storage.NuevoAlmacenSQLite(db)

	_, err = repo.GetDespachoByID(999)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestDespachoSQLite_ListarVacio(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_desp_03?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.DespachoDiario{}))
	repo := storage.NuevoAlmacenSQLite(db)

	lista, err := repo.GetAllDespachos()
	require.NoError(t, err)
	assert.Empty(t, lista)
}

func TestDespachoSQLite_Actualizar(t *testing.T) {
	db := setupDespachoDB(t, "mem_desp_04")
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateDespacho(models.DespachoDiario{
		Fecha:             time.Now().Format("2006-01-02"),
		NumeroCarrito:     10,
		IDHorario:         1,
		IDChofer:          1,
		PasajerosActuales: 5,
	})
	creado.PasajerosActuales = 12
	err := repo.UpdateDespacho(creado)
	require.NoError(t, err)

	obtenido, _ := repo.GetDespachoByID(creado.ID)
	assert.Equal(t, 12, obtenido.PasajerosActuales)
}

func TestDespachoSQLite_ActualizarNoEncontrado(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_desp_05?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.DespachoDiario{}))
	repo := storage.NuevoAlmacenSQLite(db)

	err = repo.UpdateDespacho(models.DespachoDiario{ID: 999, Fecha: "2026-01-01", NumeroCarrito: 1, IDHorario: 1, IDChofer: 1, PasajerosActuales: 0})
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestDespachoSQLite_Eliminar(t *testing.T) {
	db := setupDespachoDB(t, "mem_desp_06")
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateDespacho(models.DespachoDiario{
		Fecha:             time.Now().Format("2006-01-02"),
		NumeroCarrito:     10,
		IDHorario:         1,
		IDChofer:          1,
		PasajerosActuales: 3,
	})
	err := repo.DeleteDespacho(creado.ID)
	require.NoError(t, err)

	_, err = repo.GetDespachoByID(creado.ID)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

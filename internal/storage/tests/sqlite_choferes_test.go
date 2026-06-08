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

func setupChoferDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Chofer{}))
	return db
}

func TestChoferSQLite_CrearYObtener(t *testing.T) {
	db := setupChoferDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	chofer := models.Chofer{Nombre: "Juan Perez", Licencia: "L123456", Celular: "0991234567", Estado: models.EstadoChoferDisponible}
	creado, err := repo.CreateChofer(chofer)

	require.NoError(t, err)
	assert.Greater(t, creado.ID, 0)
	assert.Equal(t, "Juan Perez", creado.Nombre)

	obtenido, err := repo.GetChoferByID(creado.ID)
	require.NoError(t, err)
	assert.Equal(t, "L123456", obtenido.Licencia)
}

func TestChofersQLite_GetChoferByLicencia_NoEncontrado(t *testing.T) {
	db := setupChoferDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	_, err := repo.GetChoferByLicencia("NO_EXISTE")
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestChofersSQLite_ListarVacio(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_chofer_list?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Chofer{}))
	repo := storage.NuevoAlmacenSQLite(db)

	lista, err := repo.GetAllChoferes()
	require.NoError(t, err)
	assert.Empty(t, lista)
}

func TestChoferSQLite_Actualizar(t *testing.T) {
	db := setupChoferDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateChofer(models.Chofer{Nombre: "Original", Licencia: "L999", Celular: "0990000000", Estado: models.EstadoChoferDisponible})
	creado.Nombre = "Actualizado"
	err := repo.UpdateChofer(creado)
	require.NoError(t, err)

	obtenido, _ := repo.GetChoferByID(creado.ID)
	assert.Equal(t, "Actualizado", obtenido.Nombre)
}

func TestChoferSQLite_Eliminar(t *testing.T) {
	db := setupChoferDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	creado, _ := repo.CreateChofer(models.Chofer{Nombre: "Eliminar", Licencia: "L000", Celular: "09900000001", Estado: "Disponible"})
	err := repo.DeleteChofer(creado.ID)
	require.NoError(t, err)

	_, err = repo.GetChoferByID(creado.ID)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

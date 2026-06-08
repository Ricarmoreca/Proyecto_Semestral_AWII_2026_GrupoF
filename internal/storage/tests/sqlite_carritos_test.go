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

func setupCarritoDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Carrito{}))
	return db
}

func TestCarritoSQLite_CrearYObtener(t *testing.T) {
	db := setupCarritoDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	carrito := models.Carrito{Numero: 1, Estado: models.EstadoCarritoDisponible, CapacidadPasajeros: 15, Color: "Rojo"}
	err := repo.CreateCarrito(carrito)
	require.NoError(t, err)

	obtenido, err := repo.GetCarritoByNumero(1)
	require.NoError(t, err)
	assert.Equal(t, 15, obtenido.CapacidadPasajeros)
	assert.Equal(t, "Rojo", obtenido.Color)
}

func TestCarritoSQLite_GetCarritoByNumero_NoEncontrado(t *testing.T) {
	db := setupCarritoDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	_, err := repo.GetCarritoByNumero(999)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestCarritoSQLite_ListarVacio(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:mem_carrito_list?mode=memory&cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.Carrito{}))
	repo := storage.NuevoAlmacenSQLite(db)

	lista, err := repo.GetAllCarritos()
	require.NoError(t, err)
	assert.Empty(t, lista)
}

func TestCarritoSQLite_Actualizar(t *testing.T) {
	db := setupCarritoDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	repo.CreateCarrito(models.Carrito{Numero: 5, Estado: "Disponible", CapacidadPasajeros: 10, Color: "Azul"})

	err := repo.UpdateCarrito(models.Carrito{Numero: 5, Estado: "En Viaje", CapacidadPasajeros: 10, Color: "Azul"})
	require.NoError(t, err)

	obtenido, _ := repo.GetCarritoByNumero(5)
	assert.Equal(t, models.EstadoCarrito("En Viaje"), obtenido.Estado)
}

func TestCarritoSQLite_Eliminar(t *testing.T) {
	db := setupCarritoDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	repo.CreateCarrito(models.Carrito{Numero: 7, Estado: "Disponible", CapacidadPasajeros: 8, Color: "Verde"})
	err := repo.DeleteCarrito(7)
	require.NoError(t, err)

	_, err = repo.GetCarritoByNumero(7)
	assert.ErrorIs(t, err, storage.ErrNotFound)
}

func TestCarritoSQLite_CrearDuplicatePK(t *testing.T) {
	db := setupCarritoDB(t)
	repo := storage.NuevoAlmacenSQLite(db)

	repo.CreateCarrito(models.Carrito{Numero: 1, Estado: "Disponible", CapacidadPasajeros: 10, Color: "Blanco"})
	err := repo.CreateCarrito(models.Carrito{Numero: 1, Estado: "Disponible", CapacidadPasajeros: 10, Color: "Blanco"})
	assert.Error(t, err)
}

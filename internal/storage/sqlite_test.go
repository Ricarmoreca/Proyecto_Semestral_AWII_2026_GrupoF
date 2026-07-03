package storage

import (
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestAlmacenSQLiteCreaYBuscaSolicitud(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos en memoria: %v", err)
	}

	if err := db.AutoMigrate(&models.Usuario{}, &models.Solicitud{}); err != nil {
		t.Fatalf("no se pudo migrar el esquema: %v", err)
	}

	repo := NuevoAlmacenSQLite(db)
	creada := repo.CrearSolicitud(models.Solicitud{Pasajero: 1, Origen: "A", Destino: "B"})
	if creada.ID == 0 {
		t.Fatal("se esperaba un ID asignado a la solicitud creada")
	}

	lista := repo.ListarSolicitudes()
	if len(lista) != 1 {
		t.Fatalf("se esperaban 1 elemento, se obtuvieron %d", len(lista))
	}

	sol, ok := repo.BuscarSolicitudPorID(creada.ID)
	if !ok {
		t.Fatalf("no se encontró la solicitud con ID %d", creada.ID)
	}
	if sol.Origen != "A" || sol.Destino != "B" {
		t.Fatalf("la solicitud no se guardó correctamente: %+v", sol)
	}
}

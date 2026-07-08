package storage

import (
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestAlmacenSQLiteCreaYBuscaEntidades(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos en memoria: %v", err)
	}

	if err := db.AutoMigrate(&models.Usuario{}, &models.Solicitud{}, &models.Carrito{}, &models.Chofer{}, &models.Horario{}, &models.DespachoDiario{}, &models.CarritoHorarioRel{}); err != nil {
		t.Fatalf("no se pudo migrar el esquema: %v", err)
	}

	repo := NuevoAlmacenSQLite(db)

	t.Run("carrito", func(t *testing.T) {
		creado := repo.CrearCarrito(models.Carrito{Numero: 1, Estado: models.EstadoCarritoDisponible, CapacidadPasajeros: 4, Color: "Azul"})
		if creado.Numero == 0 {
			t.Fatal("se esperaba un número asignado al carrito creado")
		}

		lista := repo.ListarCarritos()
		if len(lista) != 1 {
			t.Fatalf("se esperaban 1 carrito, se obtuvieron %d", len(lista))
		}

		carrito, ok := repo.BuscarCarritoPorID(creado.Numero)
		if !ok {
			t.Fatalf("no se encontró el carrito con número %d", creado.Numero)
		}
		if carrito.Color != "Azul" {
			t.Fatalf("el carrito no se guardó correctamente: %+v", carrito)
		}
	})

	t.Run("chofer", func(t *testing.T) {
		creado := repo.CrearChofer(models.Chofer{Nombre: "Ana", Licencia: "ABC123", Celular: "999999999", Estado: models.EstadoChoferDisponible})
		if creado.ID == 0 {
			t.Fatal("se esperaba un ID asignado al chofer creado")
		}

		lista := repo.ListarChoferes()
		if len(lista) != 1 {
			t.Fatalf("se esperaban 1 chofer, se obtuvieron %d", len(lista))
		}

		chofer, ok := repo.BuscarChoferPorID(creado.ID)
		if !ok {
			t.Fatalf("no se encontró el chofer con ID %d", creado.ID)
		}
		if chofer.Licencia != "ABC123" {
			t.Fatalf("el chofer no se guardó correctamente: %+v", chofer)
		}
	})

	t.Run("horario", func(t *testing.T) {
		creado := repo.CrearHorario(models.Horario{Turno: models.TurnoMatutina, HoraInicio: "08:00", HoraFin: "12:00"})
		if creado.ID == 0 {
			t.Fatal("se esperaba un ID asignado al horario creado")
		}

		lista := repo.ListarHorarios()
		if len(lista) != 1 {
			t.Fatalf("se esperaban 1 horario, se obtuvieron %d", len(lista))
		}

		horario, ok := repo.BuscarHorarioPorID(creado.ID)
		if !ok {
			t.Fatalf("no se encontró el horario con ID %d", creado.ID)
		}
		if horario.HoraInicio != "08:00" {
			t.Fatalf("el horario no se guardó correctamente: %+v", horario)
		}
	})
}

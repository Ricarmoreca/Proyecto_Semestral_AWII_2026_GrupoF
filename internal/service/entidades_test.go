package service

import (
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

type mockCarritoRepo struct {
	crearLlamado bool
}

func (m *mockCarritoRepo) ListarCarritos() []models.Carrito { return nil }
func (m *mockCarritoRepo) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	return models.Carrito{}, false
}
func (m *mockCarritoRepo) CrearCarrito(c models.Carrito) models.Carrito {
	m.crearLlamado = true
	return c
}
func (m *mockCarritoRepo) ActualizarCarrito(id int, c models.Carrito) (models.Carrito, bool) {
	return models.Carrito{}, false
}
func (m *mockCarritoRepo) BorrarCarrito(id int) bool { return false }

type mockChoferRepo struct {
	crearLlamado bool
}

func (m *mockChoferRepo) ListarChoferes() []models.Chofer { return nil }
func (m *mockChoferRepo) BuscarChoferPorID(id int) (models.Chofer, bool) {
	return models.Chofer{}, false
}
func (m *mockChoferRepo) CrearChofer(c models.Chofer) models.Chofer {
	m.crearLlamado = true
	return c
}
func (m *mockChoferRepo) ActualizarChofer(id int, c models.Chofer) (models.Chofer, bool) {
	return models.Chofer{}, false
}
func (m *mockChoferRepo) BorrarChofer(id int) bool { return false }

type mockHorarioRepo struct {
	crearLlamado bool
}

func (m *mockHorarioRepo) ListarHorarios() []models.Horario { return nil }
func (m *mockHorarioRepo) BuscarHorarioPorID(id int) (models.Horario, bool) {
	return models.Horario{}, false
}
func (m *mockHorarioRepo) CrearHorario(h models.Horario) models.Horario {
	m.crearLlamado = true
	return h
}
func (m *mockHorarioRepo) ActualizarHorario(id int, h models.Horario) (models.Horario, bool) {
	return models.Horario{}, false
}
func (m *mockHorarioRepo) BorrarHorario(id int) bool { return false }

type mockDespachoRepo struct {
	crearLlamado bool
}

func (m *mockDespachoRepo) ListarDespachosDiarios() []models.DespachoDiario { return nil }
func (m *mockDespachoRepo) BuscarDespachoDiarioPorID(id int) (models.DespachoDiario, bool) {
	return models.DespachoDiario{}, false
}
func (m *mockDespachoRepo) CrearDespachoDiario(d models.DespachoDiario) models.DespachoDiario {
	m.crearLlamado = true
	return d
}
func (m *mockDespachoRepo) ActualizarDespachoDiario(id int, d models.DespachoDiario) (models.DespachoDiario, bool) {
	return models.DespachoDiario{}, false
}
func (m *mockDespachoRepo) BorrarDespachoDiario(id int) bool { return false }

func TestCrearEntidadInvalidaNoLlegaRepositorio(t *testing.T) {
	casos := []struct {
		nombre string
		err    error
		llamar func(t *testing.T)
	}{
		{
			nombre: "carrito",
			llamar: func(t *testing.T) {
				repo := &mockCarritoRepo{}
				service := NewCarritoService(repo)
				_, err := service.Crear(models.Carrito{Numero: 0, Estado: models.EstadoCarritoDisponible, CapacidadPasajeros: 0, Color: ""})
				if err == nil {
					t.Fatal("debía devolver error")
				}
				if repo.crearLlamado {
					t.Fatal("NO debía llamar al repositorio")
				}
			},
		},
		{
			nombre: "chofer",
			llamar: func(t *testing.T) {
				repo := &mockChoferRepo{}
				service := NewChoferService(repo)
				_, err := service.Crear(models.Chofer{Nombre: "", Licencia: "", Celular: "", Estado: models.EstadoChoferDisponible})
				if err == nil {
					t.Fatal("debía devolver error")
				}
				if repo.crearLlamado {
					t.Fatal("NO debía llamar al repositorio")
				}
			},
		},
		{
			nombre: "horario",
			llamar: func(t *testing.T) {
				repo := &mockHorarioRepo{}
				service := NewHorarioService(repo)
				_, err := service.Crear(models.Horario{Turno: "", HoraInicio: "", HoraFin: ""})
				if err == nil {
					t.Fatal("debía devolver error")
				}
				if repo.crearLlamado {
					t.Fatal("NO debía llamar al repositorio")
				}
			},
		},
		{
			nombre: "despacho",
			llamar: func(t *testing.T) {
				repo := &mockDespachoRepo{}
				service := NewDespachoService(repo)
				_, err := service.Crear(models.DespachoDiario{ID: -1, Fecha: "fecha", NumeroCarrito: 0, IDHorario: 0, IDChofer: 0, PasajerosActuales: -1})
				if err == nil {
					t.Fatal("debía devolver error")
				}
				if repo.crearLlamado {
					t.Fatal("NO debía llamar al repositorio")
				}
			},
		},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, tc.llamar)
	}
}

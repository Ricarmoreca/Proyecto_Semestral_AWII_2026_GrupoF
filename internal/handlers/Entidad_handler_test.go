package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockAlmacen struct {
	mock.Mock
}

func (m *MockAlmacen) ListarChoferes() []models.Chofer {
	args := m.Called()
	return args.Get(0).([]models.Chofer)
}

func (m *MockAlmacen) BuscarChoferPorID(id int) (models.Chofer, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Chofer), args.Bool(1)
}

func (m *MockAlmacen) CrearChofer(c models.Chofer) models.Chofer {
	args := m.Called(c)
	return args.Get(0).(models.Chofer)
}

func (m *MockAlmacen) ActualizarChofer(id int, c models.Chofer) (models.Chofer, bool) {
	args := m.Called(id, c)
	return args.Get(0).(models.Chofer), args.Bool(1)
}

func (m *MockAlmacen) BorrarChofer(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Implementar métodos vacíos de otras interfaces
func (m *MockAlmacen) ListarUsuarios() []models.Usuario {
	return []models.Usuario{}
}

func (m *MockAlmacen) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	return models.Usuario{}, false
}

func (m *MockAlmacen) CrearUsuario(u models.Usuario) models.Usuario {
	return models.Usuario{}
}

func (m *MockAlmacen) ActualizarUsuario(id int, u models.Usuario) (models.Usuario, bool) {
	return models.Usuario{}, false
}

func (m *MockAlmacen) BorrarUsuario(id int) bool {
	return false
}

func (m *MockAlmacen) ListarSolicitudes() []models.Solicitud {
	return []models.Solicitud{}
}

func (m *MockAlmacen) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func (m *MockAlmacen) CrearSolicitud(s models.Solicitud) models.Solicitud {
	return models.Solicitud{}
}

func (m *MockAlmacen) AsignarChofer(id int, choferId string) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func (m *MockAlmacen) ActualizarSolicitud(id int, s models.Solicitud) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func (m *MockAlmacen) BorrarSolicitud(id int) bool {
	return false
}

func (m *MockAlmacen) ListarCarritos() []models.Carrito {
	return []models.Carrito{}
}

func (m *MockAlmacen) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	return models.Carrito{}, false
}

func (m *MockAlmacen) CrearCarrito(c models.Carrito) models.Carrito {
	return models.Carrito{}
}

func (m *MockAlmacen) ActualizarCarrito(id int, c models.Carrito) (models.Carrito, bool) {
	return models.Carrito{}, false
}

func (m *MockAlmacen) BorrarCarrito(id int) bool {
	return false
}

func (m *MockAlmacen) ListarMantenimientos() []models.Mantenimiento {
	return []models.Mantenimiento{}
}

func (m *MockAlmacen) BuscarMantenimientoPorID(id int) (models.Mantenimiento, bool) {
	return models.Mantenimiento{}, false
}

func (m *MockAlmacen) CrearMantenimiento(mt models.Mantenimiento) models.Mantenimiento {
	return models.Mantenimiento{}
}

func (m *MockAlmacen) ActualizarMantenimiento(id int, mt models.Mantenimiento) (models.Mantenimiento, bool) {
	return models.Mantenimiento{}, false
}

func (m *MockAlmacen) BorrarMantenimiento(id int) bool {
	return false
}

func (m *MockAlmacen) AsignarCarritoHorario(carritoID int, horarioID int) (models.CarritoHorario, bool) {
	return models.CarritoHorario{}, false
}

func (m *MockAlmacen) DeasignarCarritoHorario(carritoID int, horarioID int) bool {
	return false
}

func (m *MockAlmacen) ListarDespachosDiarios() []models.DespachoDiario {
	return []models.DespachoDiario{}
}

func (m *MockAlmacen) BuscarDespachoDiarioPorID(id int) (models.DespachoDiario, bool) {
	return models.DespachoDiario{}, false
}

func (m *MockAlmacen) CrearDespachoDiario(dd models.DespachoDiario) models.DespachoDiario {
	return models.DespachoDiario{}
}

func (m *MockAlmacen) ActualizarDespachoDiario(id int, dd models.DespachoDiario) (models.DespachoDiario, bool) {
	return models.DespachoDiario{}, false
}

func (m *MockAlmacen) BorrarDespachoDiario(id int) bool {
	return false
}

func (m *MockAlmacen) ListarHorarios() []models.Horario {
	return []models.Horario{}
}

func (m *MockAlmacen) BuscarHorarioPorID(id int) (models.Horario, bool) {
	return models.Horario{}, false
}

func (m *MockAlmacen) CrearHorario(h models.Horario) models.Horario {
	return models.Horario{}
}

func (m *MockAlmacen) ActualizarHorario(id int, h models.Horario) (models.Horario, bool) {
	return models.Horario{}, false
}

func (m *MockAlmacen) BorrarHorario(id int) bool {
	return false
}

func setupChoferHandlerTest(t *testing.T) (*chi.Mux, *MockAlmacen) {
	t.Helper()
	mockAlmacen := new(MockAlmacen)
	handler := NewChoferHandler(mockAlmacen)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/choferes", func(r chi.Router) {
			r.Get("/", handler.ListarChoferes)
			r.Post("/", handler.CrearChofer)
			r.Get("/{id}", handler.ObtenerChofer)
			r.Put("/{id}", handler.ActualizarChofer)
			r.Delete("/{id}", handler.EliminarChofer)
		})
	})

	return r, mockAlmacen
}

func TestChoferHandler_Listar(t *testing.T) {
	r, mockAlmacen := setupChoferHandlerTest(t)
	esperados := []models.Chofer{
		{ID: 1, Nombre: "Carlos", Licencia: "A123", Celular: "0990000001", Estado: "Disponible"},
		{ID: 2, Nombre: "Maria", Licencia: "B456", Celular: "0990000002", Estado: "En Ruta"},
	}
	mockAlmacen.On("ListarChoferes").Return(esperados)

	req := httptest.NewRequest("GET", "/api/v1/choferes", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, 200, rr.Code)
	var respuesta []models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Len(t, respuesta, 2)
	assert.Equal(t, "Carlos", respuesta[0].Nombre)
	mockAlmacen.AssertExpectations(t)
}

func TestChoferHanler_Crear(t *testing.T) {
	r, mockAlmacen := setupChoferHandlerTest(t)

	creado := models.Chofer{ID: 3, Nombre: "Pedro", Licencia: "C789", Celular: "0990000003", Estado: "Disponible"}
	mockAlmacen.On("CrearChofer", mock.MatchedBy(func(c models.Chofer) bool {
		return c.Nombre == "Pedro"
	})).Return(creado)

	body := `{"Nombre":"Pedro","Licencia":"C789","Celular":"0990000003","Estado":"Disponible"}`
	req := httptest.NewRequest("POST", "/api/v1/choferes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	var respuesta models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Equal(t, 3, respuesta.ID)
	mockAlmacen.AssertExpectations(t)
}

func TestChoferHandlerGetByID(t *testing.T) {
	r, mockAlmacen := setupChoferHandlerTest(t)

	esperado := models.Chofer{ID: 1, Nombre: "Luis", Licencia: "D001", Celular: "0990000004", Estado: "Disponible"}
	mockAlmacen.On("BuscarChoferPorID", 1).Return(esperado, true)

	req := httptest.NewRequest("GET", "/api/v1/choferes/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var respuesta models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Equal(t, "Luis", respuesta.Nombre)
	mockAlmacen.AssertExpectations(t)
}

func TestChoferHandlerGetByIDNoEncontrado(t *testing.T) {
	r, mockAlmacen := setupChoferHandlerTest(t)

	mockAlmacen.On("BuscarChoferPorID", 999).Return(models.Chofer{}, false)

	req := httptest.NewRequest("GET", "/api/v1/choferes/999", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	mockAlmacen.AssertExpectations(t)
}

func TestChoferHandlerCrearValidationError(t *testing.T) {
	r, _ := setupChoferHandlerTest(t)

	body := `{"Nombre":"","Licencia":"","Celular":"","Estado":"Invalido"}`
	req := httptest.NewRequest("POST", "/api/v1/choferes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

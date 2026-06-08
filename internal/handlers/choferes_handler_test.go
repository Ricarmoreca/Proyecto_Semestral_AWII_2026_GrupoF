package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockChoferRepository struct {
	mock.Mock
}

func (m *MockChoferRepository) GetAllChoferes() ([]models.Chofer, error) {
	args := m.Called()
	return args.Get(0).([]models.Chofer), args.Error(1)
}

func (m *MockChoferRepository) GetChoferByID(id int) (models.Chofer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Chofer), args.Error(1)
}

func (m *MockChoferRepository) GetChoferByLicencia(licencia string) (models.Chofer, error) {
	args := m.Called(licencia)
	return args.Get(0).(models.Chofer), args.Error(1)
}

func (m *MockChoferRepository) CreateChofer(c models.Chofer) (models.Chofer, error) {
	args := m.Called(c)
	return args.Get(0).(models.Chofer), args.Error(1)
}

func (m *MockChoferRepository) UpdateChofer(c models.Chofer) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockChoferRepository) DeleteChofer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupChoferHandlerTest(t *testing.T) (*chi.Mux, *MockChoferRepository) {
	t.Helper()
	mockRepo := new(MockChoferRepository)

	mockAlmacen := &struct{ storage.AlmacenMock }{
		AlmacenMock: storage.AlmacenMock{
			GetAllChoferesFunc: mockRepo.GetAllChoferes,
			GetChoferByIDFunc:  mockRepo.GetChoferByID,
			CreateChoferFunc:   func(c models.Chofer) (models.Chofer, error) { return mockRepo.CreateChofer(c) },
			UpdateChoferFunc:   func(c models.Chofer) error { return mockRepo.UpdateChofer(c) },
			DeleteChoferFunc:   mockRepo.DeleteChofer,
		},
	}

	_ = service.NewChoferService(mockAlmacen)
	handler := NewChoferHandler(mockAlmacen)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/choferes", func(r chi.Router) {
			r.Get("/", handler.List)
			r.Post("/", handler.Create)
			r.Get("/{id}", handler.GetByID)
			r.Put("/{id}", handler.Update)
			r.Delete("/{id}", handler.Delete)
		})
	})

	return r, mockRepo
}

func TestChoferHandler_Listar(t *testing.T) {
	r, mockRepo := setupChoferHandlerTest(t)
	esperados := []models.Chofer{
		{ID: 1, Nombre: "Carlos", Licencia: "A123", Celular: "0990000001", Estado: "Disponible"},
		{ID: 2, Nombre: "Maria", Licencia: "B456", Celular: "0990000002", Estado: "En Ruta"},
	}
	mockRepo.On("GetAllChoferes").Return(esperados, nil).Once()

	req := httptest.NewRequest("GET", "/api/v1/choferes", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, 200, rr.Code)
	var respuesta []models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Len(t, respuesta, 2)
	assert.Equal(t, "Carlos", respuesta[0].Nombre)
	mockRepo.AssertExpectations(t)
}

func TestChoferHanler_Crear(t *testing.T) {
	r, mockRepo := setupChoferHandlerTest(t)

	creado := models.Chofer{ID: 3, Nombre: "Pedro", Licencia: "C789", Celular: "0990000003", Estado: "Disponible"}
	mockRepo.On("CreateChofer", mock.MatchedBy(func(c models.Chofer) bool {
		return c.Nombre == "Pedro"
	})).Return(creado, nil).Once()

	body := `{"Nombre":"Pedro","Licencia":"C789","Celular":"0990000003","Estado":"Disponible"}`
	req := httptest.NewRequest("POST", "/api/v1/choferes", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	var respuesta models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Equal(t, 3, respuesta.ID)
	mockRepo.AssertExpectations(t)
}

func TestChoferHandlerGetByID(t *testing.T) {
	r, mockRepo := setupChoferHandlerTest(t)

	esperado := models.Chofer{ID: 1, Nombre: "Luis", Licencia: "D001", Celular: "0990000004", Estado: "Disponible"}
	mockRepo.On("GetChoferByID", 1).Return(esperado, nil).Once()

	req := httptest.NewRequest("GET", "/api/v1/choferes/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var respuesta models.Chofer
	json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.Equal(t, "Luis", respuesta.Nombre)
	mockRepo.AssertExpectations(t)
}

func TestChoferHandlerGetByIDNoEncontrado(t *testing.T) {
	r, mockRepo := setupChoferHandlerTest(t)

	mockRepo.On("GetChoferByID", 999).Return(models.Chofer{}, storage.ErrNotFound).Once()

	req := httptest.NewRequest("GET", "/api/v1/choferes/999", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	mockRepo.AssertExpectations(t)
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

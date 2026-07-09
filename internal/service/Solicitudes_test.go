package service

import (
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

type mockSolicitudRepo struct {
	crearLlamado bool
}

func (m *mockSolicitudRepo) ListarSolicitudes() []models.Solicitud {
	return nil
}

func (m *mockSolicitudRepo) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func (m *mockSolicitudRepo) CrearSolicitud(s models.Solicitud) models.Solicitud {
	m.crearLlamado = true
	return s
}

func (m *mockSolicitudRepo) ActualizarSolicitud(id int, s models.Solicitud) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func (m *mockSolicitudRepo) BorrarSolicitud(id int) bool {
	return false
}

func (m *mockSolicitudRepo) AsignarChofer(id int, chofer string) (models.Solicitud, bool) {
	return models.Solicitud{}, false
}

func TestCrearSolicitudInvalidaNoLlegaRepositorio(t *testing.T) {
	repo := &mockSolicitudRepo{}
	service := NewSolicitudService(repo)

	_, err := service.Crear(models.Solicitud{
		Pasajero: 0,
		Origen:   "",
		Destino:  "",
	})

	if err == nil {
		t.Fatal("debía devolver error")
	}

	if repo.crearLlamado {
		t.Fatal("NO debía llamar al repositorio")
	}
}

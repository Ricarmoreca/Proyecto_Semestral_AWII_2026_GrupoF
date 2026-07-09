package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func TestListarUsuariosRequiereAutenticacion(t *testing.T) {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(nil))
		r.Get("/usuarios", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba 401, se recibió %d", rr.Code)
	}
}

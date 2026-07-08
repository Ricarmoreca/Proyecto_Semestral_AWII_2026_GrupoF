package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func TestRutasDeEntidadesRequierenAutenticacion(t *testing.T) {
	casos := []struct {
		nombre string
		ruta   string
	}{
		{nombre: "carritos", ruta: "/api/v1/carritos"},
		{nombre: "choferes", ruta: "/api/v1/choferes"},
		{nombre: "horarios", ruta: "/api/v1/horarios"},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, func(t *testing.T) {
			router := chi.NewRouter()
			router.Route("/api/v1", func(r chi.Router) {
				r.Use(middleware.Auth(nil))
				r.Get(tc.ruta[len("/api/v1"):], func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(http.StatusOK)
				})
			})

			req := httptest.NewRequest(http.MethodGet, tc.ruta, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusUnauthorized {
				t.Fatalf("se esperaba 401, se recibió %d", rr.Code)
			}
		})
	}
}

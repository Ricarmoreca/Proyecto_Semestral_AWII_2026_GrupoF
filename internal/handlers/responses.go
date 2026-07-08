package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message})
}

func statusDeError(err error) int {
	switch err {
	case service.ErrUsuarioNoEncontrado, service.ErrSolicitudNoEncontrada, service.ErrChoferNoEncontrado, service.ErrMantenimientoNoEncontrado:
		return http.StatusNotFound
	case service.ErrUsuarioYaExiste, service.ErrSolicitudNoValida, service.ErrSolicitudYaCerrada, service.ErrEstadoInvalido, service.ErrCamposObligatorios:
		return http.StatusBadRequest
	case service.ErrOperacionNoPermitida:
		return http.StatusForbidden
	case service.ErrBaseDeDatos:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

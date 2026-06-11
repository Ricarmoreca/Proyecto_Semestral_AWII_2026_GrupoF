package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/go-chi/chi/v5"
)

// Estructura auxiliar para leer el cuerpo de la petición al crear viaje
type RequestCrear struct {
	PasajeroID string `json:"pasajero_id"`
	Origen     string `json:"origen"`
	Destino    string `json:"destino"`
}

// Estructura auxiliar para actualizar el estado o asignar chofer
type RequestActualizar struct {
	Estado   string `json:"estado"`
	ChoferID string `json:"chofer_id,omitempty"`
}

func (s *Server) ListarSolicitudes(w http.ResponseWriter, _ *http.Request) {
	solicitudes := s.Storage.ListarSolicitudes()
	RespondJSON(w, http.StatusOK, solicitudes)
}

func (s *Server) ObtenerSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if strings.TrimSpace(id) == "" {
		RespondError(w, http.StatusBadRequest, "id de solicitud obligatorio")
		return
	}

	solicitud, encontrada := s.Storage.BuscarSolicitudPorID(id)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "solicitud no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, solicitud)
}

func (s *Server) CrearSolicitud(w http.ResponseWriter, r *http.Request) {
	var req RequestCrear
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if strings.TrimSpace(req.PasajeroID) == "" || strings.TrimSpace(req.Origen) == "" || strings.TrimSpace(req.Destino) == "" {
		RespondError(w, http.StatusBadRequest, "pasajero, origen y destino son obligatorios")
		return
	}

	nuevaSolicitud := models.Solicitud{
		Pasajero: req.PasajeroID,
		Origen:   req.Origen,
		Destino:  req.Destino,
		Estado:   "pendiente",
	}

	creada := s.Storage.CrearSolicitud(nuevaSolicitud)
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if strings.TrimSpace(id) == "" {
		RespondError(w, http.StatusBadRequest, "id de solicitud obligatorio")
		return
	}

	var req RequestActualizar
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if strings.TrimSpace(req.Estado) == "" {
		RespondError(w, http.StatusBadRequest, "el estado es obligatorio")
		return
	}

	actualizada, encontrada := s.Storage.ActualizarSolicitud(id, models.Solicitud{Estado: req.Estado, Chofer: req.ChoferID})
	if !encontrada {
		RespondError(w, http.StatusNotFound, "solicitud no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) EliminarSolicitud(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if strings.TrimSpace(id) == "" {
		RespondError(w, http.StatusBadRequest, "id de solicitud obligatorio")
		return
	}

	if !s.Storage.BorrarSolicitud(id) {
		RespondError(w, http.StatusNotFound, "solicitud no encontrada")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

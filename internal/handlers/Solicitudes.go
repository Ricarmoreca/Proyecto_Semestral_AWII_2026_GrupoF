package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/go-chi/chi/v5"
)

// Estructura auxiliar para leer el cuerpo de la petición al crear viaje
type RequestCrear struct {
	PasajeroID int    `json:"pasajero_id"`
	Origen     string `json:"origen"`
	Destino    string `json:"destino"`
}

// Estructura auxiliar para actualizar el estado o asignar chofer
type RequestActualizar struct {
	Estado   string `json:"estado"`
	ChoferID string `json:"chofer_id,omitempty"`
}

func (s *Server) ListarSolicitudes(w http.ResponseWriter, _ *http.Request) {
	solicitudes := s.Solicitudes.Listar()
	RespondJSON(w, http.StatusOK, solicitudes)
}

func (s *Server) ObtenerSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id de solicitud obligatorio")
		return
	}

	solicitud, err := s.Solicitudes.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
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

	if req.PasajeroID == 0 || strings.TrimSpace(req.Origen) == "" || strings.TrimSpace(req.Destino) == "" {
		RespondError(w, http.StatusBadRequest, "pasajero, origen y destino son obligatorios")
		return
	}

	nuevaSolicitud := models.Solicitud{
		Pasajero: req.PasajeroID,
		Origen:   req.Origen,
		Destino:  req.Destino,
		Estado:   "pendiente",
	}

	creada, err := s.Solicitudes.Crear(nuevaSolicitud)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
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

	var choferPtr *string
	if strings.TrimSpace(req.ChoferID) != "" {
		choferPtr = &req.ChoferID
	}
	actualizada, err := s.Solicitudes.Actualizar(id, models.Solicitud{Estado: req.Estado, Chofer: choferPtr})
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) EliminarSolicitud(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id de solicitud obligatorio")
		return
	}

	if err := s.Solicitudes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

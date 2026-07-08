package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

type MantenimientoHandler struct {
	almacen storage.Almacen
}

func NewMantenimientoHandler(almacen storage.Almacen) *MantenimientoHandler {
	return &MantenimientoHandler{almacen: almacen}
}

func (h *MantenimientoHandler) ListarMantenimientos(w http.ResponseWriter, r *http.Request) {
	mantenimientos := h.almacen.ListarMantenimientos()
	RespondJSON(w, http.StatusOK, mantenimientos)
}

func (h *MantenimientoHandler) ObtenerMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	mantenimiento, ok := h.almacen.BuscarMantenimientoPorID(id)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, mantenimiento)
}

func (h *MantenimientoHandler) CrearMantenimiento(w http.ResponseWriter, r *http.Request) {
	var mantenimiento models.Mantenimiento
	if err := json.NewDecoder(r.Body).Decode(&mantenimiento); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := mantenimiento.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado := h.almacen.CrearMantenimiento(mantenimiento)
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *MantenimientoHandler) ActualizarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	var mantenimiento models.Mantenimiento
	if err := json.NewDecoder(r.Body).Decode(&mantenimiento); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	mantenimiento.ID = id
	if err := mantenimiento.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	actualizado, ok := h.almacen.ActualizarMantenimiento(id, mantenimiento)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *MantenimientoHandler) EliminarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if !h.almacen.BorrarMantenimiento(id) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

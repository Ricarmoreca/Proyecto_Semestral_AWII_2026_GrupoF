package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

type ChoferHandler struct {
	almacen storage.Almacen
}

func NewChoferHandler(almacen storage.Almacen) *ChoferHandler {
	return &ChoferHandler{almacen: almacen}
}

func (h *ChoferHandler) ListarChoferes(w http.ResponseWriter, r *http.Request) {
	choferes := h.almacen.ListarChoferes()
	RespondJSON(w, http.StatusOK, choferes)
}

func (h *ChoferHandler) ObtenerChofer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	chofer, ok := h.almacen.BuscarChoferPorID(id)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, chofer)
}

func (h *ChoferHandler) CrearChofer(w http.ResponseWriter, r *http.Request) {
	var chofer models.Chofer
	if err := json.NewDecoder(r.Body).Decode(&chofer); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := chofer.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado := h.almacen.CrearChofer(chofer)
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *ChoferHandler) ActualizarChofer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	var chofer models.Chofer
	if err := json.NewDecoder(r.Body).Decode(&chofer); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	chofer.ID = id
	if err := chofer.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	actualizado, ok := h.almacen.ActualizarChofer(id, chofer)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *ChoferHandler) EliminarChofer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if !h.almacen.BorrarChofer(id) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

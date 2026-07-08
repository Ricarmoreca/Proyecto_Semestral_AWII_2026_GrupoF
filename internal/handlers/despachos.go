package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

type DespachoDiarioHandler struct {
	almacen storage.Almacen
}

func NewDespachoDiarioHandler(almacen storage.Almacen) *DespachoDiarioHandler {
	return &DespachoDiarioHandler{almacen: almacen}
}

func (h *DespachoDiarioHandler) ListarDespachos(w http.ResponseWriter, r *http.Request) {
	despachos := h.almacen.ListarDespachosDiarios()
	RespondJSON(w, http.StatusOK, despachos)
}

func (h *DespachoDiarioHandler) ObtenerDespacho(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	despacho, ok := h.almacen.BuscarDespachoDiarioPorID(id)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, despacho)
}

func (h *DespachoDiarioHandler) CrearDespacho(w http.ResponseWriter, r *http.Request) {
	var despacho models.DespachoDiario
	if err := json.NewDecoder(r.Body).Decode(&despacho); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := despacho.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado := h.almacen.CrearDespachoDiario(despacho)
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *DespachoDiarioHandler) ActualizarDespacho(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	var despacho models.DespachoDiario
	if err := json.NewDecoder(r.Body).Decode(&despacho); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	despacho.ID = id
	if err := despacho.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	actualizado, ok := h.almacen.ActualizarDespachoDiario(id, despacho)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *DespachoDiarioHandler) EliminarDespacho(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if !h.almacen.BorrarDespachoDiario(id) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

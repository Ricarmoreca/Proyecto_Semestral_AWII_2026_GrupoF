package handlers

import (
	"encoding/json"
	"errors"
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

func (h *DespachoDiarioHandler) List(w http.ResponseWriter, r *http.Request) {
	despachos, err := h.almacen.GetAllDespachos()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, despachos)
}

func (h *DespachoDiarioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	despacho, err := h.almacen.GetDespachoByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, despacho)
}

func (h *DespachoDiarioHandler) Create(w http.ResponseWriter, r *http.Request) {
	var despacho models.DespachoDiario
	if err := json.NewDecoder(r.Body).Decode(&despacho); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := despacho.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado, err := h.almacen.CreateDespacho(despacho)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *DespachoDiarioHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	if err := h.almacen.UpdateDespacho(despacho); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, despacho)
}

func (h *DespachoDiarioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if err := h.almacen.DeleteDespacho(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

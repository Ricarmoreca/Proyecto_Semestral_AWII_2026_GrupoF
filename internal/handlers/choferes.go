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

type ChoferHandler struct {
	almacen storage.Almacen
}

func NewChoferHandler(almacen storage.Almacen) *ChoferHandler {
	return &ChoferHandler{almacen: almacen}
}

func (h *ChoferHandler) List(w http.ResponseWriter, r *http.Request) {
	choferes, err := h.almacen.GetAllChoferes()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, choferes)
}

func (h *ChoferHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	chofer, err := h.almacen.GetChoferByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, chofer)
}

func (h *ChoferHandler) Create(w http.ResponseWriter, r *http.Request) {
	var chofer models.Chofer
	if err := json.NewDecoder(r.Body).Decode(&chofer); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := chofer.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado, err := h.almacen.CreateChofer(chofer)
	if err != nil {
		if errors.Is(err, storage.ErrConflict) {
			RespondError(w, http.StatusConflict, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *ChoferHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	if err := h.almacen.UpdateChofer(chofer); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, storage.ErrConflict) {
			RespondError(w, http.StatusConflict, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, chofer)
}

func (h *ChoferHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if err := h.almacen.DeleteChofer(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

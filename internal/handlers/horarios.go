package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

type HorarioHandler struct {
	almacen storage.Almacen
}

func NewHorarioHandler(almacen storage.Almacen) *HorarioHandler {
	return &HorarioHandler{almacen: almacen}
}

func (h *HorarioHandler) ListarHorarios(w http.ResponseWriter, r *http.Request) {
	horarios := h.almacen.ListarHorarios()
	RespondJSON(w, http.StatusOK, horarios)
}

func (h *HorarioHandler) ObtenerHorario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	horario, ok := h.almacen.BuscarHorarioPorID(id)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, horario)
}

func (h *HorarioHandler) CrearHorario(w http.ResponseWriter, r *http.Request) {
	var horario models.Horario
	if err := json.NewDecoder(r.Body).Decode(&horario); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := horario.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado := h.almacen.CrearHorario(horario)
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *HorarioHandler) ActualizarHorario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	var horario models.Horario
	if err := json.NewDecoder(r.Body).Decode(&horario); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	horario.ID = id
	if err := horario.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	actualizado, ok := h.almacen.ActualizarHorario(id, horario)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *HorarioHandler) EliminarHorario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	if !h.almacen.BorrarHorario(id) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HorarioHandler) GetCarritos(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un entero")
		return
	}
	carritos := h.almacen.ListarCarritos()
	_ = id
	RespondJSON(w, http.StatusOK, carritos)
}

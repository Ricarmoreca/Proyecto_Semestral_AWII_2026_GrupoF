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

type CarritoHandler struct {
	almacen storage.Almacen
}

func NewCarritoHandler(almacen storage.Almacen) *CarritoHandler {
	return &CarritoHandler{almacen: almacen}
}

func (h *CarritoHandler) List(w http.ResponseWriter, r *http.Request) {
	carritos, err := h.almacen.GetAllCarritos()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := make([]models.CarritoResponse, 0, len(carritos))
	for _, c := range carritos {
		response = append(response, c.ToResponse())
	}
	RespondJSON(w, http.StatusOK, response)
}

func (h *CarritoHandler) GetByNumero(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	carrito, err := h.almacen.GetCarritoByNumero(numero)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, carrito.ToResponse())
}

func (h *CarritoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Carrito
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := c.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.almacen.CreateCarrito(c); err != nil {
		if errors.Is(err, storage.ErrConflict) {
			RespondError(w, http.StatusConflict, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, c.ToResponse())
}

func (h *CarritoHandler) Update(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	var c models.Carrito
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	c.Numero = numero
	if err := c.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.almacen.UpdateCarrito(c); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, c.ToResponse())
}

func (h *CarritoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	if err := h.almacen.DeleteCarrito(numero); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CarritoHandler) GetHorarios(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	horarios, err := h.almacen.GetHorariosByCarrito(numero)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, horarios)
}

type asignarHorarioRequest struct {
	IDHorario int `json:"id_horario"`
}

func (h *CarritoHandler) AsignarHorario(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	var req asignarHorarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	rel := models.CarritoHorarioRel{
		NumeroCarrito: numero,
		IDHorario:     req.IDHorario,
	}
	if err := h.almacen.AsignarCarritoHorario(rel); err != nil {
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
	RespondJSON(w, http.StatusCreated, rel)
}

func (h *CarritoHandler) DesasignarHorario(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	idHorario, err := strconv.Atoi(chi.URLParam(r, "idHorario"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id_horario debe ser un entero")
		return
	}
	rel := models.CarritoHorarioRel{
		NumeroCarrito: numero,
		IDHorario:     idHorario,
	}
	if err := h.almacen.DesasignarCarritoHorario(rel); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

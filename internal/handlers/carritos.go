package handlers

import (
	"encoding/json"
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

func (h *CarritoHandler) ListarCarritos(w http.ResponseWriter, r *http.Request) {
	carritos := h.almacen.ListarCarritos()
	RespondJSON(w, http.StatusOK, carritos)
}

func (h *CarritoHandler) ObtenerCarrito(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	carrito, ok := h.almacen.BuscarCarritoPorID(numero)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}

	RespondJSON(w, http.StatusOK, carrito)
}

func (h *CarritoHandler) CrearCarrito(w http.ResponseWriter, r *http.Request) {
	var c models.Carrito
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if err := c.Validate(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	creado := h.almacen.CrearCarrito(c)
	RespondJSON(w, http.StatusCreated, creado)
}

func (h *CarritoHandler) ActualizarCarrito(w http.ResponseWriter, r *http.Request) {
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
	actualizado, ok := h.almacen.ActualizarCarrito(numero, c)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (h *CarritoHandler) EliminarCarrito(w http.ResponseWriter, r *http.Request) {
	numero, err := strconv.Atoi(chi.URLParam(r, "numero"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "numero_carrito debe ser un entero")
		return
	}
	if !h.almacen.BorrarCarrito(numero) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
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
	horarios := h.almacen.ListarHorarios()
	_ = numero
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
	_, ok := h.almacen.AsignarCarritoHorario(numero, req.IDHorario)
	if !ok {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
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
	if !h.almacen.DeasignarCarritoHorario(numero, idHorario) {
		RespondError(w, http.StatusNotFound, storage.ErrNotFound.Error())
		return
	}
	_ = rel
	w.WriteHeader(http.StatusNoContent)
}

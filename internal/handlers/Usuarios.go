package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) ListarUsuarios(w http.ResponseWriter, _ *http.Request) {
	usuarios := s.Usuarios.Listar()
	RespondJSON(w, http.StatusOK, usuarios)
}

func (s *Server) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	usuario, err := s.Usuarios.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, usuario)
}

func (s *Server) CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	creado, err := s.Usuarios.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	actualizada, err := s.Usuarios.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Usuarios.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

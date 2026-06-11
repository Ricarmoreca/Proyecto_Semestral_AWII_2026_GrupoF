package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Storage storage.Almacen
}

func NewServer(s storage.Almacen) *Server {
	return &Server{Storage: s}
}

func (s *Server) ListarUsuarios(w http.ResponseWriter, _ *http.Request) {
	usuarios := s.Storage.ListarUsuarios()
	RespondJSON(w, http.StatusOK, usuarios)
}

func (s *Server) ObtenerUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	usuario, encontrado := s.Storage.BuscarUsuarioPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
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

	creado := s.Storage.CrearUsuario(nuevo)
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

	actualizada, encontrada := s.Storage.ActualizarUsuario(id, datos)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
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

	if !s.Storage.BorrarUsuario(id) {
		RespondError(w, http.StatusNotFound, "usuario no encontrado")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

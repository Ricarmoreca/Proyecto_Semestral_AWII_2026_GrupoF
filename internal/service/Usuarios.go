package service

import (
	"strings"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type UsuarioService struct {
	repo storage.UsuariosRepository
}

func NewUsuarioService(repo storage.UsuariosRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) Listar() []models.Usuario {
	return s.repo.ListarUsuarios()
}

func (s *UsuarioService) Obtener(id int) (models.Usuario, error) {
	u, ok := s.repo.BuscarUsuarioPorID(id)
	if !ok {
		return models.Usuario{}, ErrUsuarioNoEncontrado
	}
	return u, nil
}

func (s *UsuarioService) Crear(u models.Usuario) (models.Usuario, error) {
	if err := validacionUsuario(u); err != nil {
		return models.Usuario{}, err
	}
	return s.repo.CrearUsuario(u), nil
}

func (s *UsuarioService) Actualizar(id int, p models.Usuario) (models.Usuario, error) {
	if err := validacionUsuario(p); err != nil {
		return models.Usuario{}, err
	}
	actualizado, ok := s.repo.ActualizarUsuario(id, p)
	if !ok {
		return models.Usuario{}, ErrUsuarioNoEncontrado
	}
	return actualizado, nil
}

func (s *UsuarioService) Borrar(id int) error {
	if !s.repo.BorrarUsuario(id) {
		return ErrUsuarioNoEncontrado
	}
	return nil
}

func validacionUsuario(p models.Usuario) error {
	if strings.TrimSpace(p.Nombre) == "" || strings.TrimSpace(p.Rol) == "" {
		return ErrCamposObligatorios
	}
	return nil
}

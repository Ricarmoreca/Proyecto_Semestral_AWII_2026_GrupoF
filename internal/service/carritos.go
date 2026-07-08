package service

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

type CarritoService struct {
	repo storage.CarritosRepository
}

func NewCarritoService(repo storage.CarritosRepository) *CarritoService {
	return &CarritoService{repo: repo}
}

func (s *CarritoService) Listar() []models.Carrito {
	return s.repo.ListarCarritos()
}

func (s *CarritoService) Obtener(id int) (models.Carrito, error) {
	carrito, ok := s.repo.BuscarCarritoPorID(id)
	if !ok {
		return models.Carrito{}, ErrNotFound
	}
	return carrito, nil
}

func (s *CarritoService) Crear(carrito models.Carrito) (models.Carrito, error) {
	if err := carrito.Validate(); err != nil {
		return models.Carrito{}, err
	}
	return s.repo.CrearCarrito(carrito), nil
}

func (s *CarritoService) Actualizar(id int, carrito models.Carrito) (models.Carrito, error) {
	if err := carrito.Validate(); err != nil {
		return models.Carrito{}, err
	}
	actualizado, ok := s.repo.ActualizarCarrito(id, carrito)
	if !ok {
		return models.Carrito{}, ErrNotFound
	}
	return actualizado, nil
}

func (s *CarritoService) Borrar(id int) error {
	if !s.repo.BorrarCarrito(id) {
		return ErrNotFound
	}
	return nil
}

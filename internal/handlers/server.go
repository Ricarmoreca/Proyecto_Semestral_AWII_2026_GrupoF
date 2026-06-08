package handlers

import (
	"net/http"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Carrito        *CarritoHandler
	Horario        *HorarioHandler
	Chofer         *ChoferHandler
	DespachoDiario *DespachoDiarioHandler
}

func NewServer(almacen storage.Almacen) *Server {
	return &Server{
		Carrito:        NewCarritoHandler(almacen),
		Horario:        NewHorarioHandler(almacen),
		Chofer:         NewChoferHandler(almacen),
		DespachoDiario: NewDespachoDiarioHandler(almacen),
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/carritos", func(r chi.Router) {
			r.Get("/", s.Carrito.List)
			r.Post("/", s.Carrito.Create)
			r.Get("/{numero}", s.Carrito.GetByNumero)
			r.Put("/{numero}", s.Carrito.Update)
			r.Delete("/{numero}", s.Carrito.Delete)
			r.Get("/{numero}/horarios", s.Carrito.GetHorarios)
			r.Post("/{numero}/horarios", s.Carrito.AsignarHorario)
			r.Delete("/{numero}/horarios/{idHorario}", s.Carrito.DesasignarHorario)
		})

		r.Route("/horarios", func(r chi.Router) {
			r.Get("/", s.Horario.List)
			r.Post("/", s.Horario.Create)
			r.Get("/{id}", s.Horario.GetByID)
			r.Put("/{id}", s.Horario.Update)
			r.Delete("/{id}", s.Horario.Delete)
			r.Get("/{id}/carritos", s.Horario.GetCarritos)
		})

		r.Route("/choferes", func(r chi.Router) {
			r.Get("/", s.Chofer.List)
			r.Post("/", s.Chofer.Create)
			r.Get("/{id}", s.Chofer.GetByID)
			r.Put("/{id}", s.Chofer.Update)
			r.Delete("/{id}", s.Chofer.Delete)
		})

		r.Route("/despachos-diarios", func(r chi.Router) {
			r.Get("/", s.DespachoDiario.List)
			r.Post("/", s.DespachoDiario.Create)
			r.Get("/{id}", s.DespachoDiario.GetByID)
			r.Put("/{id}", s.DespachoDiario.Update)
			r.Delete("/{id}", s.DespachoDiario.Delete)
		})
	})

	return r
}

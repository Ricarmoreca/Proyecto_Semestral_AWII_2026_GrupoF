package main

import (
	"log"
	"net/http"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/handlers"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	store, err := storage.NewSQLiteStorage("carritos.db")
	if err != nil {
		log.Fatalf("error al abrir la base de datos SQLite: %v", err)
	}
	servidor := handlers.NewServer(store)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/usuarios", servidor.ListarUsuarios)
		r.Post("/usuarios", servidor.CrearUsuario)
		r.Get("/usuarios/{id}", servidor.ObtenerUsuario)
		r.Put("/usuarios/{id}", servidor.ActualizarUsuario)
		r.Delete("/usuarios/{id}", servidor.EliminarUsuario)

		r.Get("/solicitudes", servidor.ListarSolicitudes)
		r.Post("/solicitudes", servidor.CrearSolicitud)
		r.Put("/solicitudes/{id}", servidor.ActualizarSolicitud)
		r.Get("/solicitudes/{id}", servidor.ObtenerSolicitud)
		r.Delete("/solicitudes/{id}", servidor.EliminarSolicitud)

		r.Get("/provocarerror", func(w http.ResponseWriter, r *http.Request) {
			panic("¡Error provocado desde el servidor!")
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

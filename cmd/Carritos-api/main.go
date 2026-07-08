package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver GORM (pure-Go)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/handlers"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/middleware"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

func main() {
	gdb, err := gorm.Open(sqlite.Open("carritos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(&models.Usuario{}, &models.Solicitud{}, &models.UsuarioRepo{}, &models.Carrito{}, &models.Chofer{}, &models.DespachoDiario{}, &models.Horario{}, &models.Mantenimiento{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo cafeteria.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "carritos.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	usuarioRepo := storage.NewUsuarioRepository(gdb)
	authService := service.NuevoAuthService(usuarioRepo)
	usuarioService := service.NewUsuarioService(almacen)
	solicitudService := service.NewSolicitudService(almacen)
	servidor := handlers.NewServer(usuarioService, solicitudService, authService, almacen)
	carritoHandler := handlers.NewCarritoHandler(almacen)
	choferHandler := handlers.NewChoferHandler(almacen)
	despachoHandler := handlers.NewDespachoDiarioHandler(almacen)
	horarioHandler := handlers.NewHorarioHandler(almacen)
	mantenimientoHandler := handlers.NewMantenimientoHandler(almacen)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/login", servidor.Login)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

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

			r.Route("/carritos", func(r chi.Router) {
				r.Get("/carritos", carritoHandler.ListarCarritos)
				r.Post("/carritos", carritoHandler.CrearCarrito)
				r.Get("/carritos/{numero}", carritoHandler.ObtenerCarrito)
				r.Put("/carritos/{numero}", carritoHandler.ActualizarCarrito)
				r.Delete("/carritos/{numero}", carritoHandler.EliminarCarrito)
				r.Get("/carritos/{numero}/horarios", carritoHandler.GetHorarios)
				r.Post("/carritos/{numero}/horarios", carritoHandler.AsignarHorario)
				r.Delete("/carritos/{numero}/horarios/{idHorario}", carritoHandler.DesasignarHorario)
			})

			r.Route("/horarios", func(r chi.Router) {
				r.Get("/horarios", horarioHandler.ListarHorarios)
				r.Post("/horarios", horarioHandler.CrearHorario)
				r.Get("/horarios/{id}", horarioHandler.ObtenerHorario)
				r.Put("/horarios/{id}", horarioHandler.ActualizarHorario)
				r.Delete("/horarios/{id}", horarioHandler.EliminarHorario)
				r.Get("/horarios/{id}/carritos", horarioHandler.GetCarritos)
			})

			r.Route("/choferes", func(r chi.Router) {
				r.Get("/choferes", choferHandler.ListarChoferes)
				r.Post("/choferes", choferHandler.CrearChofer)
				r.Get("/choferes/{id}", choferHandler.ObtenerChofer)
				r.Put("/choferes/{id}", choferHandler.ActualizarChofer)
				r.Delete("/choferes/{id}", choferHandler.EliminarChofer)
			})

			r.Route("/despachos", func(r chi.Router) {
				r.Get("/despachos", despachoHandler.ListarDespachos)
				r.Post("/despachos", despachoHandler.CrearDespacho)
				r.Get("/despachos/{id}", despachoHandler.ObtenerDespacho)
				r.Put("/despachos/{id}", despachoHandler.ActualizarDespacho)
				r.Delete("/despachos/{id}", despachoHandler.EliminarDespacho)
			})

			r.Route("/mantenimientos", func(r chi.Router) {
				r.Get("/mantenimientos", mantenimientoHandler.ListarMantenimientos)
				r.Post("/mantenimientos", mantenimientoHandler.CrearMantenimiento)
				r.Get("/mantenimientos/{id}", mantenimientoHandler.ObtenerMantenimiento)
				r.Put("/mantenimientos/{id}", mantenimientoHandler.ActualizarMantenimiento)
				r.Delete("/mantenimientos/{id}", mantenimientoHandler.EliminarMantenimiento)
			})
		})

		r.Get("/provocarerror", func(w http.ResponseWriter, r *http.Request) {
			panic("¡Error provocado desde el servidor!")
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

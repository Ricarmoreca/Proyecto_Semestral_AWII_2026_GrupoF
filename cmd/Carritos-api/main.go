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
	if err := gdb.AutoMigrate(&models.Usuario{}, &models.Solicitud{}, &models.UsuarioRepo{}); err != nil {
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
	servidor := handlers.NewServer(usuarioService, solicitudService, authService)

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
		})

		r.Get("/provocarerror", func(w http.ResponseWriter, r *http.Request) {
			panic("¡Error provocado desde el servidor!")
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

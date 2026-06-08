package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/handlers"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/service"
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/storage"
)

func main() {
	addr := flag.String("addr", ":8080", "direccion del servidor")
	dbPath := flag.String("db", "./data/transporte.db", "ruta de la base de datos")
	flag.Parse()

	almacen, err := initDB(*dbPath)
	if err != nil {
		log.Fatal("no se pudo inicializar la base de datos: ", err)
	}

	usuarioRepo := storage.NewUsuarioRepository(almacen.(*storage.AlmacenSQLite).DB())
	authService := service.NuevoAuthService(usuarioRepo)
	usuarioService := service.NewUsuarioService(almacen)
	solicitudService := service.NewSolicitudService(almacen)
	servidor := handlers.NewServer(almacen, usuarioService, solicitudService, authService)

	log.Printf("Servidor escuchando en http://localhost%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, servidor.RegisterRoutes()))
}

func initDB(dbPath string) (storage.Almacen, error) {
	dir := "./data"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	gdb, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.Exec("PRAGMA foreign_keys = ON")

	if err := gdb.AutoMigrate(&models.Usuario{}, &models.Solicitud{}, &models.UsuarioRepo{}); err != nil {
		return nil, err
	}

	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", dbPath)
		if err != nil {
			return nil, err
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	if err := almacenGorm.Migrate(); err != nil {
		return nil, err
	}
	almacenGorm.Seed()

	return almacen, nil
}

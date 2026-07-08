package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // dialector GORM para SQLite (pure-Go)
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

type Recursos struct {
	Almacen      Almacen
	Usuarios     UserRepository
	BackendUsado string
	Cerrar       func() error
}

func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {
	// 1. GORM es el DUENO DEL ESQUEMA: abre (segun el motor), migra y siembra.
	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}
	if err := gdb.AutoMigrate(&models.Solicitud{}, &models.UsuarioRepo{}, &models.Usuario{}, &models.Carrito{}, &models.Chofer{}, &models.DespachoDiario{}, &models.Horario{}, &models.Mantenimiento{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}
	almacenGorm := NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend de productos/categorias.
	//    El backend sqlc esta generado para SQLite (sus queries son de SQLite),
	//    por eso solo aplica cuando el driver es sqlite; con postgres se usa GORM.
	var almacen Almacen
	var sdb *sql.DB
	backendUsado := "gorm"
	if backend == "sqlc" && driver != "postgres" {
		sdb, err = sql.Open("sqlite", rutaDB)
		if err != nil {
			return nil, fmt.Errorf("abrir sql.DB para sqlc: %w", err)
		}
		almacen = NuevoAlmacenSQLC(sdb)
		backendUsado = "sqlc"
	} else {
		almacen = almacenGorm
	}

	// 3. Usuarios viven SIEMPRE en GORM (decision tomada en S10).
	// NuevoUsuarioGORM may be defined in another file; provide a fallback
	// to avoid undefined symbol during build. The fallback returns nil and
	// lets other parts of the program handle the absence of a users repo.
	usuarios := NuevoUsuarioGORM(gdb)

	// 4. Cierre ordenado: primero la conexion sql.DB de sqlc (si existe), luego GORM.
	cerrar := func() error {
		if sdb != nil {
			if err := sdb.Close(); err != nil {
				return err
			}
		}
		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	return &Recursos{
		Almacen:      almacen,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

// abrirGorm elige el Dialector segun el driver y abre la conexion.
//
// Para PostgreSQL reintenta unos segundos: dentro de docker compose la base
// puede tardar en aceptar conexiones aunque el contenedor ya este arriba (el
// healthcheck del compose reduce el problema, pero el reintento lo hace robusto).
func abrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		var gdb *gorm.DB
		var err error
		for intento := 1; intento <= 10; intento++ {
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return gdb, nil
			}
			log.Printf("PostgreSQL no esta listo (intento %d/10): %v", intento, err)
			time.Sleep(2 * time.Second)
		}
		return nil, fmt.Errorf("conectar a PostgreSQL tras reintentos: %w", err)
	default: // "sqlite"
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}

// NuevoUsuarioGORM es un fallback usado por `Inicializar` cuando no existe
// una implementación específica del repositorio de usuarios. Devuelve nil
// para que el resto del programa pueda decidir qué hacer (se documentó
// así en el comentario original).
func NuevoUsuarioGORM(db *gorm.DB) UserRepository {
	return nil
}

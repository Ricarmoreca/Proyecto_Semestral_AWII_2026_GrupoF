package storage

import (
	"time"

	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

// ActualizarCarrito implements [Almacen].
func (a *AlmacenSQLite) ActualizarCarrito(id int, datos models.Carrito) (models.Carrito, bool) {
	var existente models.Carrito
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Carrito{}, false
	}
	datos.Numero = id
	a.db.Save(&datos)
	return datos, true
}

// ActualizarChofer implements [Almacen].
func (a *AlmacenSQLite) ActualizarChofer(id int, datos models.Chofer) (models.Chofer, bool) {
	var existente models.Chofer
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Chofer{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

// ActualizarDespachoDiario implements [Almacen].
func (a *AlmacenSQLite) ActualizarDespachoDiario(id int, datos models.DespachoDiario) (models.DespachoDiario, bool) {
	var existente models.DespachoDiario
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.DespachoDiario{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

// AsignarCarritoHorario implements [Almacen].
func (a *AlmacenSQLite) AsignarCarritoHorario(carritoID int, horarioID int) (models.CarritoHorario, bool) {
	rel := models.CarritoHorarioRel{NumeroCarrito: carritoID, IDHorario: horarioID}
	if err := a.db.Create(&rel).Error; err != nil {
		return models.CarritoHorario{}, false
	}
	return models.CarritoHorario{NumeroCarrito: carritoID, IDHorario: horarioID}, true
}

// BorrarCarrito implements [Almacen].
func (a *AlmacenSQLite) BorrarCarrito(id int) bool {
	res := a.db.Delete(&models.Carrito{}, id)
	return res.RowsAffected > 0
}

// BorrarChofer implements [Almacen].
func (a *AlmacenSQLite) BorrarChofer(id int) bool {
	res := a.db.Delete(&models.Chofer{}, id)
	return res.RowsAffected > 0
}

// BorrarDespachoDiario implements [Almacen].
func (a *AlmacenSQLite) BorrarDespachoDiario(id int) bool {
	res := a.db.Delete(&models.DespachoDiario{}, id)
	return res.RowsAffected > 0
}

// BuscarCarritoPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarCarritoPorID(id int) (models.Carrito, bool) {
	var carrito models.Carrito
	if err := a.db.First(&carrito, id).Error; err != nil {
		return models.Carrito{}, false
	}
	return carrito, true
}

// BuscarChoferPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarChoferPorID(id int) (models.Chofer, bool) {
	var chofer models.Chofer
	if err := a.db.First(&chofer, id).Error; err != nil {
		return models.Chofer{}, false
	}
	return chofer, true
}

// BuscarDespachoDiarioPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarDespachoDiarioPorID(id int) (models.DespachoDiario, bool) {
	var despacho models.DespachoDiario
	if err := a.db.First(&despacho, id).Error; err != nil {
		return models.DespachoDiario{}, false
	}
	return despacho, true
}

// CrearCarrito implements [Almacen].
func (a *AlmacenSQLite) CrearCarrito(c models.Carrito) models.Carrito {
	a.db.Create(&c)
	return c
}

// CrearChofer implements [Almacen].
func (a *AlmacenSQLite) CrearChofer(c models.Chofer) models.Chofer {
	a.db.Create(&c)
	return c
}

// CrearDespachoDiario implements [Almacen].
func (a *AlmacenSQLite) CrearDespachoDiario(d models.DespachoDiario) models.DespachoDiario {
	a.db.Create(&d)
	return d
}

// DeasignarCarritoHorario implements [Almacen].
func (a *AlmacenSQLite) DeasignarCarritoHorario(carritoID int, horarioID int) bool {
	res := a.db.Delete(&models.CarritoHorarioRel{}, "numero_carrito = ? AND id_horario = ?", carritoID, horarioID)
	return res.RowsAffected > 0
}

// ListarCarritos implements [Almacen].
func (a *AlmacenSQLite) ListarCarritos() []models.Carrito {
	var carritos []models.Carrito
	a.db.Find(&carritos)
	return carritos
}

// ListarChoferes implements [Almacen].
func (a *AlmacenSQLite) ListarChoferes() []models.Chofer {
	var choferes []models.Chofer
	a.db.Find(&choferes)
	return choferes
}

// ListarDespachosDiarios implements [Almacen].
func (a *AlmacenSQLite) ListarDespachosDiarios() []models.DespachoDiario {
	var despachos []models.DespachoDiario
	a.db.Find(&despachos)
	return despachos
}

// ListarHorarios implements [Almacen].
func (a *AlmacenSQLite) ListarHorarios() []models.Horario {
	var horarios []models.Horario
	a.db.Find(&horarios)
	return horarios
}

// BuscarHorarioPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarHorarioPorID(id int) (models.Horario, bool) {
	var horario models.Horario
	if err := a.db.First(&horario, id).Error; err != nil {
		return models.Horario{}, false
	}
	return horario, true
}

// CrearHorario implements [Almacen].
func (a *AlmacenSQLite) CrearHorario(h models.Horario) models.Horario {
	a.db.Create(&h)
	return h
}

// ActualizarHorario implements [Almacen].
func (a *AlmacenSQLite) ActualizarHorario(id int, datos models.Horario) (models.Horario, bool) {
	var existente models.Horario
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Horario{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

// BorrarHorario implements [Almacen].
func (a *AlmacenSQLite) BorrarHorario(id int) bool {
	res := a.db.Delete(&models.Horario{}, id)
	return res.RowsAffected > 0
}

// ListarMantenimientos implements [Almacen].
func (a *AlmacenSQLite) ListarMantenimientos() []models.Mantenimiento {
	var mantenimientos []models.Mantenimiento
	a.db.Find(&mantenimientos)
	return mantenimientos
}

// BuscarMantenimientoPorID implements [Almacen].
func (a *AlmacenSQLite) BuscarMantenimientoPorID(id int) (models.Mantenimiento, bool) {
	var mantenimiento models.Mantenimiento
	if err := a.db.First(&mantenimiento, id).Error; err != nil {
		return models.Mantenimiento{}, false
	}
	return mantenimiento, true
}

// CrearMantenimiento implements [Almacen].
func (a *AlmacenSQLite) CrearMantenimiento(m models.Mantenimiento) models.Mantenimiento {
	a.db.Create(&m)
	return m
}

// ActualizarMantenimiento implements [Almacen].
func (a *AlmacenSQLite) ActualizarMantenimiento(id int, datos models.Mantenimiento) (models.Mantenimiento, bool) {
	var existente models.Mantenimiento
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Mantenimiento{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

// BorrarMantenimiento implements [Almacen].
func (a *AlmacenSQLite) BorrarMantenimiento(id int) bool {
	res := a.db.Delete(&models.Mantenimiento{}, id)
	return res.RowsAffected > 0
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func NewSQLiteStorage(path string) (*AlmacenSQLite, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.Usuario{}, &models.Solicitud{}, &models.Mantenimiento{}); err != nil {
		return nil, err
	}
	return NuevoAlmacenSQLite(db), nil
}

func (a *AlmacenSQLite) ListarUsuarios() []models.Usuario {
	var u []models.Usuario
	a.db.Find(&u)
	return u
}

func (a *AlmacenSQLite) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var u models.Usuario
	if err := a.db.First(&u, id).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

func (a *AlmacenSQLite) CrearUsuario(u models.Usuario) models.Usuario {
	a.db.Create(&u)
	return u
}

func (a *AlmacenSQLite) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	var existente models.Usuario
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Usuario{}, false
	}

	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarUsuario(id int) bool {
	res := a.db.Delete(&models.Usuario{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) ListarSolicitudes() []models.Solicitud {
	var s []models.Solicitud
	a.db.Find(&s)
	return s
}

func (a *AlmacenSQLite) BuscarSolicitudPorID(id int) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) CrearSolicitud(s models.Solicitud) models.Solicitud {
	a.db.Create(&s)
	return s
}

func (a *AlmacenSQLite) AsignarChofer(id int, choferId string) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	s.Chofer = &choferId
	if err := a.db.Save(&s).Error; err != nil {
		return models.Solicitud{}, false
	}
	return s, true
}

func (a *AlmacenSQLite) ActualizarSolicitud(id int, datos models.Solicitud) (models.Solicitud, bool) {
	var s models.Solicitud
	if err := a.db.First(&s, id).Error; err != nil {
		return models.Solicitud{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitud(id int) bool {
	res := a.db.Delete(&models.Solicitud{}, id)
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Usuario{}).Count(&n)
	if n > 0 {
		return
	}

	usuarios := []models.Usuario{
		{ID: 1, Nombre: "Ana Pérez", Rol: "estudiante", Matricula: "2026001"},
		{ID: 2, Nombre: "Carlos López", Rol: "docente", Matricula: "2026002"},
		{ID: 3, Nombre: "Luis Martínez", Rol: "estudiante", Matricula: "2026003"},
	}
	a.db.Create(&usuarios)

	chofer1 := "Manuel Rodriguez"
	chofer2 := "Albert Lopez"

	solicitudes := []models.Solicitud{
		{ID: 1, Pasajero: 1, Chofer: &chofer1, Origen: "Campus", Destino: "Biblioteca", Estado: "pendiente", CreadoEn: time.Now()},
		{ID: 2, Pasajero: 2, Chofer: &chofer2, Origen: "Biblioteca", Destino: "Campus", Estado: "aceptada", CreadoEn: time.Now()},
		{ID: 3, Pasajero: 3, Chofer: nil, Origen: "Residencia", Destino: "Facultad", Estado: "pendiente", CreadoEn: time.Now()},
	}
	a.db.Create(&solicitudes)

	mantenimientos := []models.Mantenimiento{
		{ID: 1, FechaMantenimiento: "2026-07-08", Descripcion: "Revisión de aceite y filtros", EstadoMantenimiento: "Pendiente", NumeroCarrito: "001"},
		{ID: 2, FechaMantenimiento: "2026-07-09", Descripcion: "Cambio de llantas", EstadoMantenimiento: "En Progreso", NumeroCarrito: "002"},
		{ID: 3, FechaMantenimiento: "2026-07-10", Descripcion: "Revisión de frenos", EstadoMantenimiento: "Completado", NumeroCarrito: "003"},
	}
	a.db.Create(&mantenimientos)
}

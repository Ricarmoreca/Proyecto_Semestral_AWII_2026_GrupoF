package service

import "errors"

var (
	// Errores de usuario
	ErrUsuarioNoEncontrado = errors.New("usuario no encontrado")
	ErrUsuarioYaExiste     = errors.New("el usuario ya existe")

	ErrEmailEnUso            = errors.New("El email ya está registrado")
	ErrCredencialesInvalidas = errors.New("Email o contraseña incorrectos")

	// Errores de solicitud
	ErrSolicitudNoEncontrada = errors.New("solicitud no encontrada")
	ErrSolicitudNoValida     = errors.New("solicitud inválida")
	ErrSolicitudYaCerrada    = errors.New("la solicitud ya está cerrada")
	ErrEstadoInvalido        = errors.New("estado de solicitud inválido")

	// Errores de chofer / asignación
	ErrChoferNoAsignado   = errors.New("chofer no asignado")
	ErrChoferNoEncontrado = errors.New("chofer no encontrado")

	// Errores de mantenimiento
	ErrMantenimientoNoEncontrado = errors.New("mantenimiento no encontrado")

	// Errores genéricos
	ErrCamposObligatorios   = errors.New("faltan campos obligatorios")
	ErrOperacionNoPermitida = errors.New("operación no permitida")
	ErrBaseDeDatos          = errors.New("error en la base de datos")

	ErrNotFound = errors.New("recurso no encontrado")
	ErrConflict = errors.New("el recurso ya existe")
)

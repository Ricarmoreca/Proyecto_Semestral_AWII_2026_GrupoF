-- ========= USUARIOS =======================

-- name: ListarUsuario :many
SELECT id, nombre, rol, matricula FROM usuarios;

-- name: BuscarUsuarioPorID :one
SELECT id, nombre, rol, matricula FROM usuarios
WHERE id = ?;

-- name: CrearUsuario :one
INSERT INTO usuarios (nombre, rol, matricula)
Values (?, ?, ?)
RETURNING id, nombre, rol, matricula;

-- name: ActualizarUsuario :one
UPDATE usuarios
SET nombre = ?, rol = ?, matricula = ?
WHERE id = ?
RETURNING id, nombre, rol, matricula;

-- name: BorrarUsuario :execrows
DELETE FROM usuarios WHERE id = ?;

-- =========== SOLICITUDES =================

-- name: ListarSolicitudes :many
SELECT id, pasajero, chofer, origen, destino, estado, creadoen FROM solicitudes;

-- name: BuscarSolicitudPorID :one
SELECT id, pasajero, chofer, origen, destino, estado, creadoen FROM solicitudes
WHERE id = ?;

-- name: CrearSolicitud :one
INSERT INTO solicitudes (pasajero, origen, destino, estado)
VALUES (?, ?, ?, 'pendiente')
RETURNING id, pasajero, chofer, origen, destino, estado, creadoen;

-- name: AsignarChofer :one
UPDATE solicitudes
SET chofer = ?
WHERE id = ?
RETURNING id, pasajero, chofer, origen, destino, estado, creadoen;

-- name: ActualizarSolicitud :one 
UPDATE solicitudes
SET estado = ?, chofer = COALESCE(?, chofer)
WHERE id = ?
RETURNING id, pasajero, chofer, origen, destino, estado, creadoen;

-- name: BorrarSolicitud :execrows
DELETE FROM solicitudes WHERE id = ?;
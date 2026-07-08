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

-- =========== CHOFERES ====================

-- name: ListarChoferes :many
SELECT id_chofer, nombre_chofer, licencia, celular, estado_chofer FROM choferes;

-- name: BuscarChoferPorID :one
SELECT id_chofer, nombre_chofer, licencia, celular, estado_chofer FROM choferes
WHERE id_chofer = ?;

-- name: CrearChofer :one
INSERT INTO choferes (nombre_chofer, licencia, celular, estado_chofer)
VALUES (?, ?, ?, ?)
RETURNING id_chofer, nombre_chofer, licencia, celular, estado_chofer;

-- name: ActualizarChofer :one
UPDATE choferes
SET nombre_chofer = ?, licencia = ?, celular = ?, estado_chofer = ?
WHERE id_chofer = ?
RETURNING id_chofer, nombre_chofer, licencia, celular, estado_chofer;

-- name: BorrarChofer :execrows
DELETE FROM choferes WHERE id_chofer = ?;

-- =========== HORARIOS ====================

-- name: ListarHorarios :many
SELECT id_horario, turno, hora_inicio, hora_fin FROM horarios;

-- name: BuscarHorarioPorID :one
SELECT id_horario, turno, hora_inicio, hora_fin FROM horarios
WHERE id_horario = ?;

-- name: CrearHorario :one
INSERT INTO horarios (turno, hora_inicio, hora_fin)
VALUES (?, ?, ?)
RETURNING id_horario, turno, hora_inicio, hora_fin;

-- name: ActualizarHorario :one
UPDATE horarios
SET turno = ?, hora_inicio = ?, hora_fin = ?
WHERE id_horario = ?
RETURNING id_horario, turno, hora_inicio, hora_fin;

-- name: BorrarHorario :execrows
DELETE FROM horarios WHERE id_horario = ?;

-- =========== CARRITOS ====================

-- name: ListarCarritos :many
SELECT numero_carrito, estado_carrito, capacidad_pasajeros, color FROM carritos;

-- name: BuscarCarritoPorID :one
SELECT numero_carrito, estado_carrito, capacidad_pasajeros, color FROM carritos
WHERE numero_carrito = ?;

-- name: CrearCarrito :one
INSERT INTO carritos (numero_carrito, estado_carrito, capacidad_pasajeros, color)
VALUES (?, ?, ?, ?)
RETURNING numero_carrito, estado_carrito, capacidad_pasajeros, color;

-- name: ActualizarCarrito :one
UPDATE carritos
SET estado_carrito = ?, capacidad_pasajeros = ?, color = ?
WHERE numero_carrito = ?
RETURNING numero_carrito, estado_carrito, capacidad_pasajeros, color;

-- name: BorrarCarrito :execrows
DELETE FROM carritos WHERE numero_carrito = ?;

-- =========== DESPACHOS ==================

-- name: ListarDespachos :many
SELECT id_despacho, fecha, numero_carrito, id_horario, id_chofer, pasajeros_actuales FROM despachos_diarios;

-- name: BuscarDespachoPorID :one
SELECT id_despacho, fecha, numero_carrito, id_horario, id_chofer, pasajeros_actuales FROM despachos_diarios
WHERE id_despacho = ?;

-- name: CrearDespacho :one
INSERT INTO despachos_diarios (fecha, numero_carrito, id_horario, id_chofer, pasajeros_actuales)
VALUES (?, ?, ?, ?, ?)
RETURNING id_despacho, fecha, numero_carrito, id_horario, id_chofer, pasajeros_actuales;

-- name: ActualizarDespacho :one
UPDATE despachos_diarios
SET fecha = ?, numero_carrito = ?, id_horario = ?, id_chofer = ?, pasajeros_actuales = ?
WHERE id_despacho = ?
RETURNING id_despacho, fecha, numero_carrito, id_horario, id_chofer, pasajeros_actuales;

-- name: BorrarDespacho :execrows
DELETE FROM despachos_diarios WHERE id_despacho = ?;

-- =========== CARRITO-HORARIO ============

-- name: AsignarCarritoHorario :one
INSERT INTO carrito_horario (numero_carrito, id_horario, hora_asignacion)
VALUES (?, ?, ?)
RETURNING numero_carrito, id_horario, hora_asignacion;

-- name: DeasignarCarritoHorario :execrows
DELETE FROM carrito_horario WHERE numero_carrito = ? AND id_horario = ?;
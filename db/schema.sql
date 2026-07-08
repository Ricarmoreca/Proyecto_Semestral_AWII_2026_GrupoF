CREATE TABLE usuarios (
    id         INTEGER PRIMARY KEY,
    nombre     TEXT NOT NULL,
    rol        TEXT NOT NULL,
    matricula  TEXT NOT NULL
);

CREATE TABLE solicitudes (
    id          INTEGER PRIMARY KEY,
    pasajero    TEXT NOT NULL,
    chofer      TEXT,
    origen      TEXT NOT NULL,
    destino     TEXT NOT NULL,
    estado      TEXT NOT NULL DEFAULT 'pendiente',
    creadoen    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE choferes (
    id_chofer      INTEGER PRIMARY KEY,
    nombre_chofer  TEXT NOT NULL,
    licencia       TEXT NOT NULL UNIQUE,
    celular        TEXT NOT NULL,
    estado_chofer  TEXT NOT NULL DEFAULT 'Disponible'
);

CREATE TABLE horarios (
    id_horario    INTEGER PRIMARY KEY,
    turno         TEXT NOT NULL,
    hora_inicio   TEXT NOT NULL,
    hora_fin      TEXT NOT NULL
);

CREATE TABLE carritos (
    numero_carrito        INTEGER PRIMARY KEY,
    estado_carrito        TEXT NOT NULL DEFAULT 'Disponible',
    capacidad_pasajeros   INTEGER NOT NULL,
    color                 TEXT NOT NULL DEFAULT 'Sin color'
);

CREATE TABLE despachos_diarios (
    id_despacho          INTEGER PRIMARY KEY,
    fecha                TEXT NOT NULL,
    numero_carrito       INTEGER NOT NULL,
    id_horario           INTEGER NOT NULL,
    id_chofer            INTEGER NOT NULL,
    pasajeros_actuales   INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (numero_carrito) REFERENCES carritos(numero_carrito),
    FOREIGN KEY (id_horario) REFERENCES horarios(id_horario),
    FOREIGN KEY (id_chofer) REFERENCES choferes(id_chofer)
);

CREATE TABLE carrito_horario (
    numero_carrito   INTEGER NOT NULL,
    id_horario       INTEGER NOT NULL,
    hora_asignacion  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (numero_carrito, id_horario),
    FOREIGN KEY (numero_carrito) REFERENCES carritos(numero_carrito),
    FOREIGN KEY (id_horario) REFERENCES horarios(id_horario)
);
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
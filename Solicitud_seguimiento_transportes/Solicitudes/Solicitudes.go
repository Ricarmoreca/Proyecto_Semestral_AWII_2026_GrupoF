package solicitudes

import (
	"errors"
	"time"

	usuarios "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF.git/Solicitud_seguimiento_transportes/Usuarios"
)

type Solicitud struct {
	ID       string           `json:"id"`
	Pasajero usuarios.Usuario `json:"pasajero"`
	Chofer   usuarios.Usuario `json:"chofer,omitempty"`
	Origen   string           `json:"origen"`
	Destino  string           `json:"destino"`
	Estado   string           `json:"estado"` // "pendiente", "en_camino", "completado"
	CreadoEn time.Time        `json:"creado_en"`
}

var BDSolicitudes = make(map[string]Solicitud)

func CrearSolicitud(id string, pasajeroID string, origen string, destino string) (Solicitud, error) {
	pasajero, existe := usuarios.BDUsuarios[pasajeroID]
	if !existe || pasajero.Rol != "estudiante" {
		return Solicitud{}, errors.Error("usuario no válido para solicitar transporte")
	}

	nuevaSolicitud := Solicitud{
		ID:       id,
		Pasajero: pasajero,
		Origen:   origen,
		Destino:  destino,
		Estado:   "pendiente",
		CreadoEn: time.Now(),
	}

	BDSolicitudes[id] = nuevaSolicitud
	return nuevaSolicitud, nil
}

func ActualizarEstado(id string, nuevoEstado string, choferID string) (Solicitud, error) {
	solicitud, existe := BDSolicitudes[id]
	if !existe {
		return Solicitud{}, errors.New("solicitud no encontrada")
	}

	if choferID != "" {
		chofer, existeChofer := usuarios.BDUsuarios[choferID]
		if existeChofer && chofer.Rol == "chofer" {
			solicitud.Chofer = chofer
		}
	}

	solicitud.Estado = nuevoEstado
	BDSolicitudes[id] = solicitud
	return solicitud, nil
}

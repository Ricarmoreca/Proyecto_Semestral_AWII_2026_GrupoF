package solicitudes

import (
	"encoding/json"
	"net/http"
)

// Estructura auxiliar para leer el cuerpo de la petición al crear viaje
type RequestCrear struct {
	PasajeroID string `json:"pasajero_id"`
	Origen     string `json:"origen"`
	Destino    string `json:"destino"`
}

// Estructura auxiliar para actualizar el estado o asignar chofer
type RequestActualizar struct {
	Estado   string `json:"estado"`
	ChoferID string `json:"chofer_id,omitempty"`
}

// HandlerSolicitudes maneja la creación y listado de viajes
func HandlerSolicitudes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost: // ESTUDIANTE PIDE CARRITO
		var req RequestCrear
		// Decodificar el JSON que envía el cliente
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "JSON inválido"})
			return
		}

		// Generamos un ID aleatorio o incremental simulado para el viaje
		idViaje := "SOL-" + string(rune(len(BDSolicitudes)+65))

		nuevaSolicitud, err := CrearSolicitud(idViaje, req.PasajeroID, req.Origen, req.Destino)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevaSolicitud)

	case http.MethodGet: // VER TODAS LAS SOLICITUDES (Para paneles de control)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BDSolicitudes)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandlerActualizarEstado maneja el seguimiento del taxi (En camino, Completado, etc.)
// Ejemplo: /solicitudes/actualizar?id=SOL-A
func HandlerActualizarEstado(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idViaje := r.URL.Query().Get("id")
	var req RequestActualizar

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON inválido"})
		return
	}

	solicitudModificada, err := ActualizarEstado(idViaje, req.Estado, req.ChoferID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(solicitudModificada)
}

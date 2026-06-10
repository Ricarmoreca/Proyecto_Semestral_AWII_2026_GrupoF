package usuarios

import (
	"encoding/json"
	"net/http"
)

// ObtenerUsuarioHandler responde con los datos de un usuario según su ID en la URL
// Ejemplo: /usuarios?id=U1
func ObtenerUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// Configurar cabecera para responder en formato JSON
	w.Header().Set("Content-Type", "application/json")

	// Validar que sea un método GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"})
		return
	}

	// Obtener el ID desde los parámetros de la URL (?id=U1)
	id := r.URL.Query().Get("id")
	usuario, existe := BDUsuarios[id]

	if !existe {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"})
		return
	}

	// Responder con el usuario encontrado
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

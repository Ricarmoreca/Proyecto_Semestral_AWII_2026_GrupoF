package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	solicitudes "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF.git/Solicitud_seguimiento_transportes/Solicitudes"
	usuarios "github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF.git/Solicitud_seguimiento_transportes/Usuarios"
)

func main() {
	// 1. Definir las rutas (Endpoints) de la API
	http.HandleFunc("/usuarios", usuarios.ObtenerUsuarioHandler)
	http.HandleFunc("/solicitudes", solicitudes.HandlerSolicitudes)
	http.HandleFunc("/solicitudes/actualizar", solicitudes.HandlerActualizarEstado)

	// 2. Levantar el servidor
	puerto := ":8080"
	fmt.Printf("🚀 Servidor universitario de carritos de golf corriendo en http://localhost%s\n", puerto)

	err := http.ListenAndServe(puerto, nil)
	if err != nil {
		log.Fatalf("Error al levantar el servidor: %v", err)
	}
}

// Función auxiliar para imprimir bonito en consola
func imprimirJSON(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}

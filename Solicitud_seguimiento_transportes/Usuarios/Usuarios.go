package usuarios

type Usuario struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Rol       string `json:"rol"`       // "estudiante" o "chofer"
	Matricula string `json:"matricula"` // Código universitario
}

var BDUsuarios = map[string]Usuario{
	"U1": {ID: "U1", Nombre: "Carlos Gómez", Rol: "estudiante", Matricula: "20230001"},
	"U2": {ID: "U2", Nombre: "Ana Martínez", Rol: "chofer", Matricula: "20210542"},
}

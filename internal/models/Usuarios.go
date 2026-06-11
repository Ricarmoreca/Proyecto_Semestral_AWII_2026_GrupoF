package models

type Usuario struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Rol       string `json:"rol"`       // "estudiante" o "chofer"
	Matricula string `json:"matricula"` // Código universitario
}

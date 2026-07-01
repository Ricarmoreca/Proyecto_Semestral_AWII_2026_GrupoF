package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"

	"gorm.io/gorm"
)

type UsuarioRepoGORM struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepoGORM {
	return &UsuarioRepoGORM{db: db}
}

func (r *UsuarioRepoGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return models.Usuario{}, err
	}
	return u, nil
}

func (r *UsuarioRepoGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("nombre = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

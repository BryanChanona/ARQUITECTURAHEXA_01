package application

import "arquitecturaHexagonal/src/users/domain"

type ViewUserById struct {
	db domain.Iuser
}

// Constructor de ViewUserById para usuarios
func NewViewUserById(db domain.Iuser) *ViewUserById {
	return &ViewUserById{db: db}
}

// Método para ejecutar la búsqueda del usuario por ID
func (viewUser *ViewUserById) Execute(id int) (domain.User, error) {
	return viewUser.db.GetById(id)
}

package application

import "arquitecturaHexagonal/src/users/domain"

type ViewUsers struct {
	db domain.Iuser
}
func NewViewUsers(db domain.Iuser)*ViewUsers{
	return &ViewUsers{db:db}
}

func (useCase *ViewUsers) Execute()([]domain.User, error){
	return useCase.db.GetAll()
}


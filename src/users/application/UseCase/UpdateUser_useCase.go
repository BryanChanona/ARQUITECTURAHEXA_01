package application

import "arquitecturaHexagonal/src/users/domain"

type UpdateUser struct {
	db domain.Iuser
}

func NewUpdateUser(db domain.Iuser) *UpdateUser {
	return &UpdateUser{db: db}

}

func (useCase *UpdateUser) Execute(id int, user domain.User) (err error) {
	return useCase.db.UpdateUser(id, user)
}

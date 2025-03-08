package domain

type Iuser interface {
	SaveUser(user User)(err error)
	GetAll()([]User,error)
	DeleteUser(id int) error
	UpdateUser(id int,user User) error
	GetById(id int)(User, error)
}
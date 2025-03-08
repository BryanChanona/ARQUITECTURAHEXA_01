package domain

type Ibook interface {
	SaveBook(book Book) (err error)
	GetAll() ([]Book, error)
	DeleteBook(id int) error
	UpdateBook(id int, book Book) error
	GetById(id int)(Book, error)
}

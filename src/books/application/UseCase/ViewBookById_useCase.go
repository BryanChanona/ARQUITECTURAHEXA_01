package UseCase

import "arquitecturaHexagonal/src/books/domain"

type ViewBookById struct {
	db domain.Ibook
}

func NewViewProductById(db domain.Ibook) *ViewBookById{
	return &ViewBookById{db:db}
}

func (viewBook *ViewBookById) Execute(id int)(domain.Book, error){
	return viewBook.db.GetById(id)
}

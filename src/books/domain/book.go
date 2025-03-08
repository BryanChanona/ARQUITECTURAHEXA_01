package domain

type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
	
}

func NewCreateBook(title string, author string) *Book {
	return &Book{Title: title, Author: author}
}

func (book *Book) GetTitle() string {
	return book.Title
}

func (book *Book) SetTitle(title string) {
	book.Title = title
}

func (book *Book) GetAuthor() string {
	return book.Author
}

func (book *Book) SetAuthor(author string) {
	book.Author = author
}

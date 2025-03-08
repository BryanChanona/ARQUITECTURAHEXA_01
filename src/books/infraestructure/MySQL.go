package infraestructure

import (
	"database/sql"
	"fmt"
	"arquitecturaHexagonal/src/books/domain"
)

// Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}

// Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) SaveBook(book domain.Book) (err error) {
	// Preparar la sentencia SQL para insertar el libro
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO books (title, author) VALUES (?, ?)")

	if err != nil {
		return err
	}

	// Asegurarse de que la sentencia preparada se ejecute correctamente al final
	defer sentenciaPreparada.Exec()

	// Ejecutar la sentencia preparada con los valores del libro
	_, err = sentenciaPreparada.Exec(book.GetTitle(), book.GetAuthor())

	if err != nil {
		return err
	}
	return nil
}
func (mysql *MySQL) GetAll() ([]domain.Book, error) {
	data, err := mysql.db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var books []domain.Book
	// Itera sobre todas las filas devueltas por la consulta
	for data.Next() {
		var book domain.Book
		err := data.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
func (mysql *MySQL) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	result, err := mysql.db.Exec(query, id)
	if err != nil {
		return err
	}
	// Verificar si realmente se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el libro con ID %d", id)
	}
	fmt.Println("Libro eliminado correctamente")
	return nil
}
func (mysql *MySQL) UpdateBook(id int, book domain.Book) error {
	// Verificar si el libro con el ID dado existe
	var count int
	query := "SELECT COUNT(*) FROM books WHERE id = ?"
	err := mysql.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("error verificando si el libro existe con ID %d: %v", id, err)
	}
	if count == 0 {
		return fmt.Errorf("no se encontró el libro con ID %d", id)
	}

	// Proceder con la actualización si el libro existe
	bookTitle := book.GetTitle()
	bookAuthor := book.GetAuthor()

	updateQuery := "UPDATE books SET title = ?, author = ? WHERE id = ?"
	_, err = mysql.db.Exec(updateQuery, bookTitle, bookAuthor, id)
	if err != nil {
		return fmt.Errorf("error actualizando el libro con ID %d: %v", id, err)
	}

	return nil
}

func (mysql *MySQL) GetById(id int)(domain.Book, error){
	var bookById domain.Book
	query:= "SELECT id, title, author FROM books WHERE id =?"

	row := mysql.db.QueryRow(query,id)
	
	err := row.Scan(&bookById.ID,&bookById.Title,&bookById.Author)

	if err != nil {
		if err == sql.ErrNoRows{
			return bookById, fmt.Errorf("book con ID %d no encontrado", id)
		}
		return bookById, err
	}
	return bookById, nil
}






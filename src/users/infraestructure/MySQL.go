package infraestructure

import (
	"database/sql"
	"fmt"

	"arquitecturaHexagonal/src/users/domain"
)

// Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}

// Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) SaveUser(user domain.User) (err error) {
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO users ( name, email) VALUES (?,?)")

	if err != nil {
		return err
	}

	defer sentenciaPreparada.Exec()

	_, err = sentenciaPreparada.Exec(user.GetName(), user.GetEmail())

	if err != nil {
		return err
	}
	return nil

}

func (mysql *MySQL) GetAll() ([]domain.User, error) {

	data, err := mysql.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var users []domain.User
	// Itera sobre todas las filas devueltas por la consulta
	for data.Next() {
		var user domain.User
		err := data.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func (mysql *MySQL) DeleteUser(id int) error {

	query := "DELETE FROM users WHERE id = ?"
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
		return fmt.Errorf("no se encontró el usuario con ID %d", id)
	}
	fmt.Println("Producto eliminado correctamente")
	return nil

}

func (mysql *MySQL) UpdateUser(id int, user domain.User) error {
	// Verificar si el ID existe
	var count int
	query := "SELECT COUNT(*) FROM users WHERE id = ?"
	err := mysql.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("error al verificar el ID de usuario: %v", err)
	}

	// Si no existe el usuario con el ID proporcionado, devolver error
	if count == 0 {
		return fmt.Errorf("no se encontró el usuario con ID %d", id)
	}

	// Proceder a actualizar el usuario
	userName := user.Name
	emailUser := user.Email

	updateQuery := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err = mysql.db.Exec(updateQuery, userName, emailUser, id)
	if err != nil {
		return fmt.Errorf("error actualizando el usuario con ID %d: %v", id, err)
	}

	return nil
}
func (mysql *MySQL) GetById(id int) (domain.User, error) {
	var userById domain.User
	query := "SELECT id, name, email FROM users WHERE id = ?"

	// Ejecutar la consulta
	row := mysql.db.QueryRow(query, id)

	// Escanear los resultados de la consulta en el objeto userById
	err := row.Scan(&userById.ID, &userById.Name, &userById.Email)

	// Manejo de errores
	if err != nil {
		if err == sql.ErrNoRows {
			return userById, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		return userById, err
	}

	// Si no hubo errores, retornar el usuario encontrado
	return userById, nil
}



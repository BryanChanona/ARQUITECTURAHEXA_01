package helpers

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
	_"github.com/go-sql-driver/mysql" // Importa el driver de MySQL
)

// Me devuelve la conexión a la base de datos
func ConnectDB() (db *sql.DB, err error) {

	//Verificamos si las variabvles de entorno se cargan exitosamente
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudieron cargar las variables de entorno")
	}

	user_DB := os.Getenv("DB_USER")
	password_DB := os.Getenv("DB_PASSWORD")
	host_DB := os.Getenv("DB_HOST")
	name_DB := os.Getenv("DB_NAME")
	//Formateamos una cadena.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user_DB, password_DB, host_DB, name_DB)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Error al abrir conexión: %s\n", err.Error())
		return nil, err
	}
	//Manejo de conexión  poll
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		fmt.Printf("Error al verificar la conexión: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil

}

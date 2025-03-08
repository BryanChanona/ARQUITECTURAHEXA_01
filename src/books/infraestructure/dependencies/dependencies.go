package dependencies

import (
	service "arquitecturaHexagonal/src/books/application/Service"
	application "arquitecturaHexagonal/src/books/application/UseCase"
	"arquitecturaHexagonal/src/books/infraestructure"
	"arquitecturaHexagonal/src/books/infraestructure/adapters"
	"arquitecturaHexagonal/src/books/infraestructure/controller"
	"arquitecturaHexagonal/src/helpers"
	"log"
)

var (
	mySQL infraestructure.MySQL
	eventService *service.Event 
)

func Init(){
	db, err := helpers.ConnectDB()
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos: %v", err)
		}
	mySQL =*infraestructure.NewMySQL(db)
	rabbit := adapters.NewRabbit() // Crear la conexión Rabbit
	eventService = service.NewEvent(rabbit) // Inicializar el servicio Event con Rabbit

}

func GetCreateBookController() *controller.CreateBookController {
	caseCreateBook := application.NewCreateBook(&mySQL)
	return controller.NewCreateBookController(caseCreateBook, eventService)
}
func GetViewBooksController() *controller.ViewBooksController {
	caseViewBooks := application.NewViewBooks(&mySQL)
	return controller.NewViewBooksController(caseViewBooks)
}
func GetDeleteBookController() *controller.DeleteBookController {
	caseDeleteBook := application.NewDeleteBook(&mySQL)
	return controller.NewDeleteBookController(caseDeleteBook)
}
func GetUpdateBookController() *controller.UpdateBookController {
	caseUpdateBook := application.NewUpdateBook(&mySQL)
	return controller.NewUpdateBookController(caseUpdateBook)
}
func GetViewBookByIdController()*controller.ViewBookByIdController{
	caseViewById:= application.NewViewProductById(&mySQL)
	return controller.NewViewbyIdController(caseViewById)
}




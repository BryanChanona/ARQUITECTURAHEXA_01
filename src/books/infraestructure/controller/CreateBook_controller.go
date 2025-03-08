package controller

import (
	service "arquitecturaHexagonal/src/books/application/Service"
	"arquitecturaHexagonal/src/books/application/UseCase"

	"arquitecturaHexagonal/src/books/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookController struct {
	createBook *UseCase.CreateBook
	event      *service.Event
}

// Constructor que inicializa tanto el caso de uso como el servicio de eventos (RabbitMQ)
func NewCreateBookController(useCase *UseCase.CreateBook, event *service.Event) *CreateBookController {
	return &CreateBookController{
		createBook: useCase,
		event:      event,
	}
}

// Ejecutar el flujo de creación de libro y publicar evento en RabbitMQ
func (controller *CreateBookController) Execute(ctx *gin.Context) {
	var book domain.Book

	// Bind del JSON recibido desde el frontend
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ejecutar el caso de uso para crear el libro
	err := controller.createBook.Execute(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con un mensaje de éxito
	ctx.JSON(http.StatusCreated, gin.H{"message": "Book created"})

	// Publicar el evento en RabbitMQ
	err = controller.event.Execute(book)
	if err != nil {
		// Puedes registrar este error en los logs si el evento no se envió correctamente
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending event to RabbitMQ"})
	}
}

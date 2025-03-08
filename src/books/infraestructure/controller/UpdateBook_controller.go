package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/books/application/UseCase"
	"arquitecturaHexagonal/src/books/domain"
	"github.com/gin-gonic/gin"
)

type UpdateBookController struct {
	updateBook *application.UpdateBook
}

func NewUpdateBookController(useCase *application.UpdateBook) *UpdateBookController {
	return &UpdateBookController{updateBook: useCase}
}

func (controller *UpdateBookController) Execute(c *gin.Context) {
	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")

	// Convertir el ID a entero
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = controller.updateBook.Execute(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Libro actualizado correctamente"})
}

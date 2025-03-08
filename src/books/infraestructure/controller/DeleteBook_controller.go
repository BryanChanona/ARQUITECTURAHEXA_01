package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/books/application/UseCase"
	"github.com/gin-gonic/gin"
)

type DeleteBookController struct {
	deleteBook *application.DeleteBook
}

func NewDeleteBookController(useCase *application.DeleteBook) *DeleteBookController {
	return &DeleteBookController{deleteBook: useCase}
}

func (controller *DeleteBookController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = controller.deleteBook.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book eliminado correctamente"})
}

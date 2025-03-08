package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/books/application/UseCase"
	"github.com/gin-gonic/gin"
)

type ViewBookByIdController struct {
	viewBookById *application.ViewBookById
}

func NewViewbyIdController(useCase *application.ViewBookById) *ViewBookByIdController{
	return &ViewBookByIdController{viewBookById:useCase}
}

func (controllerViewBookById *ViewBookByIdController)Execute(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
	}
	book, err := controllerViewBookById.viewBookById.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"book": book})
}
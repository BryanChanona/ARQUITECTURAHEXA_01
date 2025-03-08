package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/users/application/UseCase"
	"github.com/gin-gonic/gin"
)

type ViewUserByIdController struct {
	viewUserById *application.ViewUserById
}

// Constructor para ViewUserByIdController
func NewViewUserByIdController(useCase *application.ViewUserById) *ViewUserByIdController {
	return &ViewUserByIdController{viewUserById: useCase}
}

// Ejecuta la lógica para obtener un usuario por ID
func (controllerViewUserById *ViewUserByIdController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")          // Obtener el parámetro de ID
	id, err := strconv.Atoi(idParam)    // Convertir el ID a entero
	if err != nil {                     // Si no es un número válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamada al caso de uso para obtener el usuario
	user, err := controllerViewUserById.viewUserById.Execute(id)
	if err != nil { // Si ocurre un error en la ejecución
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si todo es correcto, devolver el usuario en la respuesta
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

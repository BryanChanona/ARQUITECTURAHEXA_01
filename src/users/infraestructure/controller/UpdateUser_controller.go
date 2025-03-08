package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/users/application/UseCase"
	"arquitecturaHexagonal/src/users/domain"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateController(useCase *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: useCase}
}

func (controller *UpdateUserController) Execute(c *gin.Context) {
	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")

	// Convertir el ID a entero
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}


	err = controller.updateUser.Execute(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}





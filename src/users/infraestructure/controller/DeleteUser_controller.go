package controller

import (
	"net/http"
	"strconv"

	application "arquitecturaHexagonal/src/users/application/UseCase"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(useCase *application.DeleteUser) *DeleteUserController{
	return &DeleteUserController{deleteUser: useCase}
}

func (controller *DeleteUserController)Execute(ctx *gin.Context){
	idParam:= ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err =controller.deleteUser.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User eliminado correctamente"})

}

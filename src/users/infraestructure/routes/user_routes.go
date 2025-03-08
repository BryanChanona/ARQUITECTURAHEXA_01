package routes

import (
	"arquitecturaHexagonal/src/users/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine){
	routes := router.Group("/users")
	createUserController := dependencies.GetCreateUserController().Execute
	viewUsersController := dependencies.GetViewUsersController().Execute
	deleteUserController := dependencies.GetDeleteUserController().Execute
	updateUserController := dependencies.GetUpdateUserController().Execute
	viewUserByIdController := dependencies.GetViewUserByIdController().Execute



	routes.POST("/",createUserController)
	routes.GET("/",viewUsersController)
	routes.DELETE("/:id",deleteUserController)
	routes.PUT("/:id", updateUserController)
	routes.GET("/:id",viewUserByIdController)

}
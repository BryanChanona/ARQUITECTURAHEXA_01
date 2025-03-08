package routes

import (
	"arquitecturaHexagonal/src/books/infraestructure/dependencies"
	
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {
	routes := router.Group("/books")
	createBookController := dependencies.GetCreateBookController().Execute
	viewBooksController := dependencies.GetViewBooksController().Execute
	deleteBookController := dependencies.GetDeleteBookController().Execute
	updateBookController := dependencies.GetUpdateBookController().Execute
	viewBookbyIdController := dependencies.GetViewBookByIdController().Execute
	
	


	routes.POST("/", createBookController)
	routes.GET("/",viewBooksController)
	routes.DELETE("/:id",deleteBookController)
	routes.PUT("/:id",updateBookController)
	routes.GET("/:id",viewBookbyIdController)


}
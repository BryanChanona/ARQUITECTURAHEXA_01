package main

import (
	dependenciesBook "arquitecturaHexagonal/src/books/infraestructure/dependencies"
	bookRoutesResource "arquitecturaHexagonal/src/books/infraestructure/routes"
	"arquitecturaHexagonal/src/helpers"
	dependenciesUser "arquitecturaHexagonal/src/users/infraestructure/dependencies"
	userRoutesResource "arquitecturaHexagonal/src/users/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main(){

	dependenciesUser.Init()
	dependenciesBook.Init()


	r := gin.Default()

	helpers.InitCORS(r)
	
	// Configurar rutas
	
	userRoutesResource.UserRouter(r)
	bookRoutesResource.BookRouter(r)

	r.Run(":8083") 
}
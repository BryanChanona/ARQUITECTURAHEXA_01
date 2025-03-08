package dependencies

import (
	"arquitecturaHexagonal/src/helpers"
	application "arquitecturaHexagonal/src/users/application/UseCase"
	"arquitecturaHexagonal/src/users/infraestructure"
	"arquitecturaHexagonal/src/users/infraestructure/controller"
	"log"
)

var (
	mySQL infraestructure.MySQL
)

func Init(){
	db, err := helpers.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL =*infraestructure.NewMySQL(db)

}

func GetCreateUserController() *controller.CreateUserController{
	caseCreateUser := application.NewCreateUser(&mySQL)
	return controller.NewCreateUserController(caseCreateUser)
}
func GetViewUsersController()*controller.ViewUsersController{
	caseViewUsers := application.NewViewUsers(&mySQL)
	return controller.NewViewUsersController(caseViewUsers)
}
func GetDeleteUserController()*controller.DeleteUserController{
	caseDeleteUser := application.NewDeleteUser(&mySQL)
	return controller.NewDeleteUserController(caseDeleteUser)
}
func GetUpdateUserController()*controller.UpdateUserController{
	caseUpdateUser := application.NewUpdateUser(&mySQL)
	return controller.NewUpdateController(caseUpdateUser)
}
func GetViewUserByIdController()* controller.ViewUserByIdController{
	caseViewUserById := application.NewViewUserById(&mySQL)
	return controller.NewViewUserByIdController(caseViewUserById)
}

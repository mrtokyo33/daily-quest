package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/daily-quest/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/findUserById/:userId", userController.FindUserByID)
	r.GET("/findUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}

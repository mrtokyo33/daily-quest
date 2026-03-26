package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/daily-quest/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/findUserById/:userId", controller.FindUserById)
	r.GET("/findUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/daily-quest/src/configuration/validation"
	"github.com/mrtokyo33/daily-quest/src/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

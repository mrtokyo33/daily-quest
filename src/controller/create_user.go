package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/daily-quest/src/configuration/logger"
	"github.com/mrtokyo33/daily-quest/src/configuration/validation"
	"github.com/mrtokyo33/daily-quest/src/controller/model/request"
	"github.com/mrtokyo33/daily-quest/src/model"
	"github.com/mrtokyo33/daily-quest/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully!", zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")
}

package service

import (
	"fmt"

	"github.com/mrtokyo33/daily-quest/src/configuration/logger"
	"github.com/mrtokyo33/daily-quest/src/configuration/rest_err"
	"github.com/mrtokyo33/daily-quest/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	// regra de negócio
	userDomain.EncryptPassword()

	// debug
	fmt.Println(userDomain.GetPassword())

	return nil

}

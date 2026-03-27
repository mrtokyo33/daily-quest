package service

import (
	"github.com/mrtokyo33/daily-quest/src/configuration/rest_err"
	"github.com/mrtokyo33/daily-quest/src/model"
)

// construtor
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

// implementação privada
type userDomainService struct {
}

// Interface define o contrato do service
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

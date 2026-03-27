package service

import (
	"github.com/mrtokyo33/daily-quest/src/configuration/rest_err"
	"github.com/mrtokyo33/daily-quest/src/model"
)

func (*userDomainService) FindUser(string) (model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

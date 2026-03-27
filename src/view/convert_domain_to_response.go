package view

import (
	"github.com/mrtokyo33/daily-quest/src/controller/model/response"
	"github.com/mrtokyo33/daily-quest/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

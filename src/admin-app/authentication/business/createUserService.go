package business

import (
	genericConstants "ABCD/src/constants"
	"ABCD/src/utils/postgres"
	"authentication/common/constants"
	"authentication/models"
	"authentication/repositories"
	"context"
	"strings"
)

type CreateUserService struct {
	createUserRepositories repositories.CreateUserRepositories
}

func NewGetCreateUserService(createUserRepositories repositories.CreateUserRepositories) *CreateUserService {
	return &CreateUserService{
		createUserRepositories: createUserRepositories,
	}
}

func (service *CreateUserService) CreateUser(ctx context.Context, Span context.Context, bffCreateUserRequest models.BFFCreateUserRequest) error {

	postgresClient := postgres.GetDBInstance()

	condition := map[string]interface{}{
		genericConstants.Username: bffCreateUserRequest.Username, genericConstants.Name: bffCreateUserRequest.Name, genericConstants.Email: bffCreateUserRequest.Email,
		genericConstants.Password: bffCreateUserRequest.Password, genericConstants.PhoneNumber: bffCreateUserRequest.PhoneNumber, genericConstants.PanCard: bffCreateUserRequest.PanCard,
	}
	err := service.createUserRepositories.CreateUsrByCondition(ctx, postgresClient.GormDB, condition)
	if err != nil {
		if strings.Contains(err.Error(), constants.UniqueConstraintViolationError) {
			return err
		}
	}
	return nil

}

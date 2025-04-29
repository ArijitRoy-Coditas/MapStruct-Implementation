package business

import (
	"ABCD/src/utils/postgres"
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
		"username": bffCreateUserRequest.Username, "name": bffCreateUserRequest.Name, "email": bffCreateUserRequest.Email,
		"password": bffCreateUserRequest.Password, "phoneNumber": bffCreateUserRequest.PhoneNumber, "panCard": bffCreateUserRequest.PanCard,
	}
	err := service.createUserRepositories.CreateUsrByCondition(ctx, postgresClient.GormDB, condition)
	if err != nil {
		if strings.Contains(err.Error(), "Duplication key violation") {
			return err
		}
	}
	return nil

}

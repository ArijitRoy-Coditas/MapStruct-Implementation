package repositories

import (
	genericModels "ABCD/src/models"
	"context"

	"gorm.io/gorm"
)

type createUserRepositories struct {
}

type mockCreateUserRepositories struct{}

type CreateUserRepositories interface {
	CreateUsrByCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) error
}

func GetCreateUserRepositories() *createUserRepositories {
	return &createUserRepositories{}
}

func GetMockCreateUserRepositories() *mockCreateUserRepositories {
	return &mockCreateUserRepositories{}
}

func NewGetCreateUserRepositories(useDBMocks bool) CreateUserRepositories {
	if useDBMocks {
		return GetMockCreateUserRepositories()
	}
	return GetCreateUserRepositories()
}

func (cu *createUserRepositories) CreateUsrByCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) error {

	result := db.WithContext(ctx).Model(&genericModels.User{}).Create(condition)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cu *mockCreateUserRepositories) CreateUsrByCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) error {
	return nil
}

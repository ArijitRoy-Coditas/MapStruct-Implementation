package postgres

import (
	"ABCD/src/constants"
	"ABCD/src/models"
	"ABCD/src/utils/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresConfig *models.PostgresConfig
	postgresClient *models.DatabaseConfiguration
)

func InitPostgresClient(configPath string) error {
	if err := initPostgres(configPath); err != nil {
		return err
	}

	dsn := fmt.Sprintf(constants.DSNString, postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.DBName, postgresConfig.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf(constants.DBInitializationError, err)
	}

	setDBInstance(db)

	return nil
}

func setDBInstance(db *gorm.DB) {
	postgresClient = &models.DatabaseConfiguration{GormDB: db}
}

func GetDBInstance() *models.DatabaseConfiguration {
	return postgresClient
}

func initPostgres(configPath string) error {
	var err error
	postgresConfig, err = config.LoadConfig[models.PostgresConfig](configPath, constants.PostgresConfigName, constants.ConfigType)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"ABCD/src/constants"
	"ABCD/src/models"
	"ABCD/src/utils/postgres"
	"log"
)

func main() {

	if err := postgres.InitPostgresClient(constants.RootPath); err != nil {
		log.Fatalln(constants.DBInstanceFailedError)
	}

	postgresClient := postgres.GetDBInstance().GormDB

	if err := postgresClient.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf(constants.MigrationFailedError, err)
	}

	log.Println(constants.MigrationSuccessMessage)

}

package main

import (
	"ABCD/src/constants"
	"ABCD/src/utils/postgres"
	"authentication/router"
	"fmt"
	"log"
)

func main() {
	if err := postgres.InitPostgresClient(constants.BasePath); err != nil {
		log.Fatalln(constants.DBInstanceFailedError)
	}

	startRouter()
}

func startRouter() {
	router := router.GetRouter()
	router.Run(fmt.Sprintf(":%d", 8080))
}

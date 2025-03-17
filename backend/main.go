package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	engine := getEngine()
	serverErr := engine.Run(fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

/*
POST http://localhost:8585/api/v1/automations/workflows
{
	"name": "test-connection-Postgres-zxLRC7my"
	"workflowType": "TEST_CONNECTION",
	"request": {
		"serviceType": "Database",
		"connectionType": "Postgres",
		"connection": {
			"config": {
				"type": "Postgres",
				"scheme": "postgresql+psycopg2",
				"username": "postgres",
				"authType": {
					"password": "password"
				},
				.....
			}
		}
	}
}

POST http://localhost:8585/api/v1/automations/workflows/trigger/899e5895-1fab-4f80-b5bc-e9fe22009c09

GET http://localhost:8585/api/v1/automations/workflows/899e5895-1fab-4f80-b5bc-e9fe22009c09

DELETE http://localhost:8585/api/v1/automations/workflows/899e5895-1fab-4f80-b5bc-e9fe22009c09?hardDelete=true

*/

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

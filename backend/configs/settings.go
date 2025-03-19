package configs

import (
	"os"
	"log"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Settings struct {
	ServerHost string
	ServerPort int

	DatabaseHost string
	DatabasePort int
	DatabaseName string
	DatabaseUser string
	DatabasePassword string

	SystemVersion string
	SystemRevision string
	SystemTimestamp int
}

func NewSettings() *Settings {
	// Load env
	dotenvErr := godotenv.Load()
	
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	settings := &Settings{}

	// Server
	serverHost, ok := os.LookupEnv("SERVER_HOST")
	if ok {
		settings.ServerHost = serverHost
	} else {
		settings.ServerHost = "localhost"
	}

	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if ok {
		port, err := strconv.ParseInt(serverPort, 10, 64)
		if err != nil {
			log.Fatal("Invalid server port")
		}
		settings.ServerPort = int(port)
	} else {
		settings.ServerPort = 8000
	}

	// Database
	databaseHost, ok := os.LookupEnv("DATABASE_HOST")
	if ok {
		settings.DatabaseHost = databaseHost
	} else {
		settings.DatabaseHost = "localhost"
	}

	databasePort, ok := os.LookupEnv("DATABASE_PORT")
	if ok {
		port, err := strconv.ParseInt(databasePort, 10, 64)
		if err != nil {
			log.Fatal("Invalid database port")
		}
		settings.DatabasePort = int(port)
	} else {
		settings.DatabasePort = 5432
	}

	databaseName, ok := os.LookupEnv("DATABASE_NAME")
	if ok {
		settings.DatabaseName = databaseName
	} else {
		settings.DatabaseName = "go_social"
	}

	databaseUser, ok := os.LookupEnv("DATABASE_USER")
	if ok {
		settings.DatabaseUser = databaseUser
	} else {
		settings.DatabaseUser = "admin"
	}

	databasePassword, ok := os.LookupEnv("DATABASE_PASSWORD")
	if ok {
		settings.DatabasePassword = databasePassword
	} else {
		settings.DatabasePassword = "admin"
	}

	// System
	version, ok := os.LookupEnv("VERSION")
	if ok {
		settings.SystemVersion = version
	} else {
		settings.SystemVersion = "1.6.5"
	}

	revision, ok := os.LookupEnv("REVISION")
	if ok {
		settings.SystemRevision = revision
	} else {
		settings.SystemRevision = "c34abd832f2fd7acc08c9a8e833181587703f0f2"
	}

	timestamp, ok := os.LookupEnv("TIMESTAMP")
	if ok {
		timestamp, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			log.Fatal("Invalid system timestamp")
		}
		settings.SystemTimestamp = int(timestamp)
	} else {
		settings.SystemTimestamp = 1740754501563
	}

	return settings
}

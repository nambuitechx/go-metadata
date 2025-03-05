package configs

import (
	"log"
	"fmt"
	
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DatabaseConnection struct {
	DB *sqlx.DB
}

func NewDatabaseConnection(settings *Settings) *DatabaseConnection {
	databaseConnection := &DatabaseConnection{}

	db, err := sqlx.Connect(
		"pgx",
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v",
			settings.DatabaseUser,
			settings.DatabasePassword,
			settings.DatabaseHost,
			settings.DatabasePort,
			settings.DatabaseName,
		),
	)

    if err != nil {
        log.Fatalln(err)
    } else {
		log.Println("Database connected")
	}

	databaseConnection.DB = db

	return databaseConnection
}

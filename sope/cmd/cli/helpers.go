package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	sop.RootPath = path
	sop.DB.DataType = os.Getenv("DATABASE_TYPE")
}

func getDSN() string {
	dbType := sop.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	}
	return "mysql://" + sop.BuildDSN()
}

func showHelp() {
	color.Yellow(`Available commands:

	help                  		- show the help commands
	version               		- print application version
	migrate               		- runs all up migrations that have not been run previously
	migrate down          		- reverses the most recent migration
	migrate reset         		- runs all down migrations in reverse order, and then all up migrations
	make migration <name> 		- creates two new up and down migrations in the migrations folder
	make auth			  		- creates and runs migrations for authentication tables, and creates models and middleware
	make controller <name>		- creates a stub controller in the controllers directory
	make mode <name>			- creates a new model in data directory
	make session				- creates a table in the database as session store
	`)
}

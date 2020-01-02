package database

import (
	"os"

	"github.com/go-pg/pg/v9"
)

type envVarKey string

const (
	databaseUserEnv envVarKey = "DATABASE_USER"
	databasePassEnv envVarKey = "DATABASE_PASSWORD"
	databaseAddress envVarKey = "DATABASE_ADDRESS"
)

func NewDatabase() *pg.DB {
	user := getEnvOrDefault(databaseUserEnv, "postgres")
	password := getEnvOrDefault(databasePassEnv, "secretpass")
	address := getEnvOrDefault(databaseAddress, "localhost:5432")

	db := pg.Connect(&pg.Options{
		Addr:     address,
		User:     user,
		Password: password,
	})

	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	return db
}

func getEnvOrDefault(key envVarKey, defaultValue string) string {
	if environmentValue := os.Getenv(string(key)); environmentValue != "" {
		return environmentValue
	}
	return defaultValue
}

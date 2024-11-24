package posgresdb

import (
	"client/internal/utils/env"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	err_db_open = errors.New("can not open database : ")
)

func Database(env *env.Env) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	if env.DBDriver == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Brussels",
			env.DBHost,
			env.DBUser,
			env.DBPass,
			env.DBName,
			env.DBPort,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err_db_open, err)
		}
	} else {
		db, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
		if err != nil {
			log.Fatal(err_db_open, err)
		}
	}

	return db
}

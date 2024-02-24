package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DatabaseClient struct {
	DB *sqlx.DB
}

func Initialize() *DatabaseClient {
	db, err := sqlx.Connect("sqlite3", "hrm-local.db")
	if err != nil {
		log.Fatalln(err)
	}

	client := &DatabaseClient{DB: db}
	client.Migrate()

	return client
}

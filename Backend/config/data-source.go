package config

import (
	"api/app/utils"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var logg Logger
var env utils.LoadEnv

type DataSource struct {
	Host     string
	Port     string
	Password string
	User     string
	Database string
	SSLmode  string
}

type DataSourceInterface interface {
	New()
}

func (ds *DataSource) New() *DataSource {
	return &DataSource{
		Host:     os.Getenv("APP_ENV"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		SSLmode:  "disable",
	}
}

func (ds *DataSource) OpenConnection() (*sql.DB, error) {
	env.New()
	db := ds.New()

	conn, err := sql.Open(
		"postgres",
		"host="+db.Host+
			" port="+db.Port+
			" user="+db.User+
			" password="+db.Password+
			" dbname="+db.Database+
			" sslmode="+db.SSLmode)

	if err != nil {
		logg.New()
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		logg.New()
		panic(err)
	}

	return conn, err
}

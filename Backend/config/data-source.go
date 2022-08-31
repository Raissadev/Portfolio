package config

import (
	"api/app/utils"
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var env utils.LoadEnv
var Instance *gorm.DB

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
	env.New()
	return &DataSource{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		SSLmode:  "disable",
	}
}

func (ds *DataSource) Open() *gorm.DB {
	data := ds.New()
	dbURL := ("postgres://" + data.User + ":" + data.Password + "@" + data.Host + ":" + data.Port + "/" + data.Database)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Instance = db

	return db
}

func (ds *DataSource) Conn() (*sql.DB, error) {
	db := ds.Open()
	conn, err := db.DB()

	if err != nil {
		log.Fatalln(err)
	}

	Instance = db

	return conn, nil
}

func CloseDB() error {
	db, _ := Instance.DB()

	return db.Close()
}

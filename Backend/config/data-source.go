package config

import (
	"api/app/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logg Logger
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
	return &DataSource{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		SSLmode:  "disable",
	}
}

func (ds *DataSource) Open() {
	// dbURL := ("postgres://" + ds.User + ":" + ds.Password + "@" + ds.Host + ":" + ds.Port + "/" + ds.Database)
	dbURL := ("postgres://root:password@localhost:5432/portfolio")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Instance = db
}
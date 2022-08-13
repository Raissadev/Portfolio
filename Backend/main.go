package main

import (
	"api/app/routes"
	"api/config"
	"fmt"
)

var dataSource config.DataSource

func main() {
	fmt.Println("Hello World")

	dataSource.OpenConnection()

	routes.Router()
}

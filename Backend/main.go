package main

import (
	"api/app/routes"
	"api/config"
	"api/config/logger"
	"fmt"
	"os"
	"runtime"
)

var dataSource config.DataSource

func main() {
	fmt.Println("Hello World")

	dataSource.Open()

	routes.Router()

	logger.Log.Printf(
		"Server v pid=%d started with processes: %d", os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

}

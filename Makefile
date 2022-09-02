BINARY_NAME=application
exec = app

build:
	make
	cd ./Backend
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

run_app: build run

clean:
	go clean
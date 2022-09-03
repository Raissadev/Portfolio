BINARY_NAME = application
exec = application

build:
	ls;
	cd ./Backend && go mod download;
	cd ./Backend && go build -o bin/${BINARY_NAME}

run:
	cd ./Backend && ./bin/${BINARY_NAME}

run_app: build run

clean:
	go clean
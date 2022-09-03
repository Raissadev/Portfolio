FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o ./bin/application

EXPOSE 8000

CMD ["./bin/application"]
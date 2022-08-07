FROM golang:1.18.1-alpine

LABEL version="1.0.0"
LABEL description="My portfolio"

WORKDIR /app/

COPY . /

RUN go mod download

RUN go build -o api main.go

EXPOSE 4000

CMD ["/app/api"]
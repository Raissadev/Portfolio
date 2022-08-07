FROM golang:1.18.1-alpine

LABEL version="1.0.0"
LABEL description="My portfolio"

WORKDIR /app/
COPY . ./

RUN go mod download && \
    go build

# EXPOSE $PORT

CMD ["go", "run"]

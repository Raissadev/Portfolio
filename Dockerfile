FROM golang

LABEL version="1.0.0"
LABEL description="My portfolio"

WORKDIR /app/
COPY . ./

RUN go install && \
    go build

# EXPOSE $PORT

CMD ["go", "run"]

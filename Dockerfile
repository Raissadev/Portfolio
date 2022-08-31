FROM debian:10

LABEL version="1.0.0"
LABEL description="My portfolio"

RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

RUN apt update && \
    apt install -y golang-go

WORKDIR /app/

COPY . .

EXPOSE 4000

CMD ["/app/api"]
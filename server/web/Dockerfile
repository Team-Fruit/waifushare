FROM golang:stretch

ENV DOCKERIZE_VERSION v0.6.0

RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go get -u github.com/labstack/echo
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/jmoiron/sqlx

ENTRYPOINT ["dockerize", "-wait", "tcp://db:3306"]
CMD ["go", "run", "/var/app/server.go"]


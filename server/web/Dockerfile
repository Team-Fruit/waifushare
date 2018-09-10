FROM golang:stretch

RUN go get -u github.com/labstack/echo
RUN go get github.com/jmoiron/sqlx

CMD ["go", "run", "/var/app/server.go"]


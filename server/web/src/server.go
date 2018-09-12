package main

import (
    "log"
    "os"
    "net/http"
    "github.com/labstack/echo"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

func main() {
    db, err := sqlx.Connect("mysql", "waifushare:@tcp(db:3306)/waifushare_db")
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    defer db.Close()
    
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":8080"))
}


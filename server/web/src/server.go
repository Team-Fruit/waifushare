package main

import (
    "log"
    "net/http"
    // "starconv" funny typo !
    // "strconv"
    "database/sql"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

type (
    User struct {
        ID uint
        UserID string `db:"user_id"`
        DisplayName sql.NullString `db:"display_name"`
        PasswordHash sql.NullString `db:"password_hash"`
        SavedImageId sql.NullInt64 `db:"saved_image_id"`
        AccountStatus int `db:"account_status"`
    }
    Twimg struct {
        ID uint
        TwitterID string `db:"twitter_id"`
        FileName string `db:"file_name"`
    }
    Image struct {
        ID uint
        CreatedAt string `db:"created_at"`
        CreatedBy uint `db:"created_by"`
        TwitterID sql.NullString `db:"twitter_id"`
    }
    UserImageLike struct {
        UserID uint `db:"user_id"`
        ImageID uint `db:"image_id"`
        IsLike int `db:"is_like"`
    }
)

const (
    AdminPassword = "$2a$10$hM3xaS4f7i/fAH2pjQxRA.ylxGqE1X2MYUtWohSRuSgyFOCIkOvMe"
)

var (
    db *sqlx.DB
)

func main() {
    var err error
    db, err = sqlx.Connect("mysql", "waifushare:@tcp(db:3306)/waifushare_db")
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Close()
    
    e := echo.New()
    
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", hello)

    apiv1 := e.Group("/api/v1")
    
    apiv1.POST("/image", handler.UploadImage)
    apiv1.PUT("/user/password", handler.UpdateUserPassword)
    apiv1.PUT("/user", handler.UpdateUser)
    apiv1.POST("/user", handler.CreateUser)
    apiv1.DELETE("/user", handler.DeleteUser)
    
    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World")
    /* stmt, err := db.Preparex("SELECT COUNT(id) FROM user")
    if err != nil {
        return err
    }
    var cnt int
    err = stmt.Get(&cnt)
    if err != nil {
        return err
    }
    return c.String(http.StatusOK, strconv.Itoa(cnt)) */
}


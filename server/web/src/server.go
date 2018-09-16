package main

import (
    "log"
    "net/http"
    // "starconv" funny typo !
    "strconv"
    "database/sql"
    "golang.org/x/crypto/bcrypt"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

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
    e.POST("/api/v1/image", uploadImage)
    e.PUT("/api/v1/user/password", updateUserPassword)
    e.PUT("/api/v1/user", updateUser)
    e.POST("/api/v1/user", createUser)
    e.DELETE("/api/v1/user", deleteUser)
    
    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    stmt, err := db.Preparex("SELECT COUNT(id) FROM user")
    if err != nil {
        return err
    }
    var cnt int
    err = stmt.Get(&cnt)
    if err != nil {
        return err
    }
    return c.String(http.StatusOK, strconv.Itoa(cnt))
}

func uploadImage(c echo.Context) error {
    return c.String(http.StatusOK, "upload")
}

func updateUserPassword(c echo.Context) error {
    return c.String(http.StatusOK, "userPasswordUpdate")
}

func updateUser(c echo.Context) error {
    return c.String(http.StatusOK, "userUpdate")
}

func createUser(c echo.Context) error {
    return c.String(http.StatusOK, "createUser")
}

func deleteUser(c echo.Context) error {
    return c.String(http.StatusOK, "deleteUser")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


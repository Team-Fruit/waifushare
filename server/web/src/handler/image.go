package handler

import (
    "net/http"
    "os"
    "io"

    "github.com/labstack/echo"
)

type Message struct {
    Message string `json:"message"`
}

func UploadImage(c echo.Context) error {
    // username := c.FormValue("UserName")
    // password := c.FormValue("Password")
    // tweetid := c.FormValue("TweetID")

    file, err := c.FormFile("Image")
    if err != nil {
        return err
    }
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    dst, err := os.Create("../images/" + file.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()

    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    res := &Message{
        Message: "OK",
    }
    return c.JSON(http.StatusOK, res)
}


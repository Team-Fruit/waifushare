package handler

import (
    "net/http"
    "os"
    "io"

    "github.com/labstack/echo"
)

type (
    Image struct {
        Username string `form:"username" validate:"required"`
        Password string `form:"password" validate:"required"`
        Tweet_id string `form:"tweet_id" validate:"required,numeric"`
    }
    Message struct {
        Message string `json:"message"`
    }
)

func UploadImage(c echo.Context) error {
    cc := c.(*CustomContext)
    i := new(Image)
    if err := cc.BindValidate(i); err != nil {
        return err
    }
    
    file, err := cc.FormFile("image")
    if file == nil {
        return echo.NewHTTPError(http.StatusBadRequest)
    }
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

    res := &Message {
        Message: "OK",
    }
    return c.JSON(http.StatusOK, res)
}


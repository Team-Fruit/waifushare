package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func UploadImage(c echo.Context) error {
    return c.String(http.StatusOK, "upload")
}


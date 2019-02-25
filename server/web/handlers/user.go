package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func UpdateUserPassword(c echo.Context) error {
    return c.String(http.StatusOK, "userPasswordUpdate")
}

func UpdateUser(c echo.Context) error {
    return c.String(http.StatusOK, "userUpdate")
}

func CreateUser(c echo.Context) error {
    return c.String(http.StatusOK, "createUser")
}

func DeleteUser(c echo.Context) error {
    return c.String(http.StatusOK, "deleteUser")
}


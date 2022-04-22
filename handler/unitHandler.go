package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const unitTemplate = "unit.html"

func (h Handler) GetUnits(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) GetUnit(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) PostUnit(c echo.Context) error {
	log.Println("create")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) DeleteUnit(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) PutUnit(c echo.Context) error {
	log.Println("update")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

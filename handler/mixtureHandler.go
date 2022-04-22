package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const mixtureTemplate = "mixture.html"

func (h Handler) GetMixtures(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, mixtureTemplate, nil)
}

func (h Handler) GetMixture(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, mixtureTemplate, nil)
}

func (h Handler) PostMixture(c echo.Context) error {
	log.Println("create")
	return c.Render(http.StatusCreated, mixtureTemplate, nil)
}

func (h Handler) DeleteMixture(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, mixtureTemplate, nil)
}

func (h Handler) PutMixture(c echo.Context) error {
	log.Println("update")
	return c.Render(http.StatusOK, mixtureTemplate, nil)
}

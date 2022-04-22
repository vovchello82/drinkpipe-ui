package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const categoryTemplate = "category.html"

func (h Handler) GetCategories(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) GetCategory(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) PostCategory(c echo.Context) error {
	log.Println("create")
	return c.Render(http.StatusCreated, categoryTemplate, nil)
}

func (h Handler) DeleteCategory(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) PutCategory(c echo.Context) error {
	log.Println("update")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

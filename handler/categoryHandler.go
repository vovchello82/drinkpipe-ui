package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const categoryTemplate = "category.html"

type Category struct {
	Id   string `json:"id,omitempty" form:"id,omitempty"`
	Name string `json:"name" form:"name" validate:"required"`
	Type string `json:"type" form:"type" validate:"required"`
}

func (h Handler) GetCategories(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, categoryTemplate, categories)
}

func (h Handler) EditCategory(c echo.Context) error {
	log.Println("edit single")
	return c.Render(http.StatusOK, "editCategory.html", nil)
}

func (h Handler) GetCategory(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) PostCategory(c echo.Context) error {
	cat := new(Category)

	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("created %u", cat)
	categories = append(categories, cat)
	return c.Render(http.StatusCreated, categoryTemplate, categories)
}

func (h Handler) DeleteCategory(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) PutCategory(c echo.Context) error {
	log.Println("update")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

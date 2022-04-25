package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const categoryTemplate = "category.html"

type CategoryType string

const (
	UNKNOWN      CategoryType = "UNKNOWN"
	FOOD         CategoryType = "FOOD"
	NONE_FOOD    CategoryType = "NONE_FOOD"
	ALCOHOL      CategoryType = "ALCOHOL"
	NONE_ALCOHOL CategoryType = "NONE_ALCOHOL"
)

var categoryTypes = []CategoryType{
	UNKNOWN,
	FOOD,
	NONE_FOOD,
	ALCOHOL,
	NONE_ALCOHOL,
}

type Category struct {
	Id   string `json:"id,omitempty" param:"id" form:"id,omitempty"`
	Name string `json:"name" form:"name" validate:"required"`
	Type string `json:"type" form:"type" validate:"required"`
}

func (h Handler) GetCategories(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, categoryTemplate, categories)
}

func (h Handler) GetNewCategory(c echo.Context) error {
	log.Println("new category")
	return c.Render(http.StatusOK, "newCategory.html", map[string]interface{}{
		"types": categoryTypes,
	})
}

func (h Handler) GetEditCategory(c echo.Context) error {
	var id string
	echo.PathParamsBinder(c).String("id", &id)
	log.Println("edit single: ", id)

	category := categories[id]

	return c.Render(http.StatusOK, "editCategory.html", map[string]interface{}{
		"Category": category,
		"types":    categoryTypes,
	})
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

	cat.Id = uuid.New().String()
	categories[cat.Id] = cat
	log.Printf("created %s", cat)
	return c.Render(http.StatusCreated, categoryTemplate, categories)
}

func (h Handler) DeleteCategory(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h Handler) PutCategory(c echo.Context) error {
	cat := new(Category)

	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("put single: %s", cat)
	category := categories[c.Param("id")]
	category.Name = cat.Name
	category.Type = cat.Type

	return c.Render(http.StatusOK, categoryTemplate, categories)
}

package handler

import (
	"drinkpipe-ui/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var categories = map[string]*Category{}

const categoryTemplate = "category.html"

const categoryKeyPrefix = "cat"

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
	Piece
	Type string `json:"type" form:"type" validate:"required"`
}

type Piece struct {
	Id   string `json:"id,omitempty" param:"id" form:"id,omitempty"`
	Name string `json:"name" form:"name" validate:"required"`
}

func (c *Category) ConvertToJson() ([]byte, error) {
	value, err := json.Marshal(c)

	if err != nil {
		log.Println("marshal error")
		return nil, err
	}

	return value, nil
}

func (c *Category) ConvertFromJson(jsonString []byte) (store.Entity, error) {
	if err := json.Unmarshal(jsonString, &c); err != nil {
		return nil, err
	}

	return c, nil
}

func (p Piece) GetId() string {
	return p.Id
}

func (h *Handler) GetCategories(c echo.Context) error {

	if values, err := h.Repository.GetAll(categoryKeyPrefix); err == nil {
		log.Println("get all from redis")
		for _, v := range values {
			fmt.Println(v)
		}
	}

	return c.Render(http.StatusOK, categoryTemplate, categories)
}

func (h *Handler) GetNewCategory(c echo.Context) error {
	log.Println("new category")
	return c.Render(http.StatusOK, "newCategory.html", map[string]interface{}{
		"types": categoryTypes,
	})
}

func (h *Handler) GetEditCategory(c echo.Context) error {
	var id string
	echo.PathParamsBinder(c).String("id", &id)
	log.Println("edit single: ", id)

	category := categories[id]

	return c.Render(http.StatusOK, "editCategory.html", map[string]interface{}{
		"Category": category,
		"types":    categoryTypes,
	})
}

func (h *Handler) GetCategory(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h *Handler) PostCategory(c echo.Context) error {
	cat := new(Category)

	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cat.Id = uuid.New().String()
	categories[cat.Id] = cat
	log.Printf("created %s", cat)

	if err := h.Repository.Persist(cat, categoryKeyPrefix); err != nil {
		log.Println("Error on persisting", err.Error())
	}
	if values, err := h.Repository.GetAll(categoryKeyPrefix); err == nil {
		for _, v := range values {
			fmt.Println(v)
		}
	}

	return c.Render(http.StatusCreated, categoryTemplate, categories)
}

func (h *Handler) DeleteCategory(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, categoryTemplate, nil)
}

func (h *Handler) PutCategory(c echo.Context) error {
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

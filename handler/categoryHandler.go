package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

func convertCategoryFromJson(jsonString []byte) (*Category, error) {
	cat := new(Category)
	if err := json.Unmarshal(jsonString, &cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (p Piece) GetId() string {
	return p.Id
}

func (h *Handler) GetCategories(c echo.Context) error {

	if values, err := h.Repository.GetAll(categoryKeyPrefix); err == nil {
		cats := make(map[string]*Category, len(values))
		for _, v := range values {
			if cat, err := convertCategoryFromJson([]byte(v)); err == nil {
				cats[cat.Id] = cat
			}
		}
		return c.Render(http.StatusOK, categoryTemplate, cats)
	}

	return c.Render(http.StatusOK, categoryTemplate, []*Category{})
}

func (h *Handler) GetNewCategory(c echo.Context) error {
	log.Println("new category")
	return c.Render(http.StatusOK, "newCategory.html", map[string]interface{}{
		"types": categoryTypes,
	})
}

func (h *Handler) GetEditCategory(c echo.Context) error {
	var id string
	//TODO VZ path param object binding
	echo.PathParamsBinder(c).String("id", &id)
	log.Println("edit single: ", id)

	//TODO VZ error handling 404
	catJson, _ := h.Repository.GetById(id, categoryKeyPrefix)

	if cat, err := convertCategoryFromJson([]byte(catJson)); err == nil {
		return c.Render(http.StatusOK, "editCategory.html", map[string]interface{}{
			"Category": cat,
			"types":    categoryTypes,
		})
	}

	return h.GetCategories(c)
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
	if err := h.Repository.Persist(cat, categoryKeyPrefix); err != nil {
		log.Println("error on cat persisting", err.Error())
	} else {
		log.Printf("created %s", cat)
	}

	return h.GetCategories(c)
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
	//TODO VZ error handling 404
	categoryJson, _ := h.Repository.GetById(c.Param("id"), categoryKeyPrefix)

	if category, err := convertCategoryFromJson([]byte(categoryJson)); err == nil {
		category.Name = cat.Name
		category.Type = cat.Type
		if err := h.Repository.Persist(category, categoryKeyPrefix); err != nil {
			log.Println("error on cat persisting", err.Error())
		} else {
			log.Printf("updated %s", cat)
		}
	}

	return h.GetCategories(c)
}

package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const unitTemplate = "unit.html"

type Unit struct {
	Id           string `json:"id,omitempty" form:"id,omitempty"`
	Name         string `json:"name" form:"name" validate:"required"`
	EAN          string `json:"ean" form:"ean"`
	CategoryName string `json:"category_name" form:"category_name" validate:"required"`
	Flavour      string `json:"flavour" form:"flavour"`
}

func (h Handler) GetUnits(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) GetUnit(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) GetUpdateUnit(c echo.Context) error {
	log.Println("edit single")
	return c.Render(http.StatusOK, "updateUnit.html", categories)
}

func (h Handler) PostUnit(c echo.Context) error {
	u := new(Unit)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("created %u", u)
	return c.Render(http.StatusCreated, unitTemplate, nil)
}

func (h Handler) DeleteUnit(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) PutUnit(c echo.Context) error {
	log.Println("update")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

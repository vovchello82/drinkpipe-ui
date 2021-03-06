package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const unitTemplate = "unit.html"

const unitKeyPrefix = "unit"

type Unit struct {
	Piece
	EAN          string `json:"ean" form:"ean"`
	CategoryName string `json:"category_name" form:"category_name" validate:"required"`
	Flavour      string `json:"flavour" form:"flavour"`
}

func (u *Unit) ConvertToJson() ([]byte, error) {
	value, err := json.Marshal(u)

	if err != nil {
		log.Println("marshal error")
		return nil, err
	}

	return value, nil
}

func convertUnitFromJson(jsonString []byte) (*Unit, error) {
	cat := new(Unit)
	if err := json.Unmarshal(jsonString, &cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (h Handler) GetUnits(c echo.Context) error {
	if values, err := h.Repository.GetAll(unitKeyPrefix); err == nil {
		units := make(map[string]*Unit, len(values))
		for _, v := range values {
			if unit, err := convertUnitFromJson([]byte(v)); err == nil {
				units[unit.Id] = unit
			}
		}
		return c.Render(http.StatusOK, unitTemplate, units)
	}

	return c.Render(http.StatusOK, unitTemplate, []*Unit{})
}

func (h *Handler) GetNewUnit(c echo.Context) error {
	log.Println("new unit")
	if values, err := h.Repository.GetAll(categoryKeyPrefix); err == nil {
		categories := make(map[string]string, len(values))

		for _, v := range values {
			if cat, err := convertCategoryFromJson([]byte(v)); err == nil {
				categories[cat.Name] = cat.Type
			}

		}
		return c.Render(http.StatusOK, "newUnit.html", map[string]interface{}{
			"categories": categories,
		})
	}

	return c.Render(http.StatusOK, "newUnit.html", map[string]interface{}{})
}

func (h Handler) GetUnit(c echo.Context) error {
	log.Println("get single")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) GetEditUnit(c echo.Context) error {
	var id string
	//TODO VZ path param object binding
	echo.PathParamsBinder(c).String("id", &id)
	log.Println("edit single: ", id)

	//TODO VZ error handling 404
	unitJson, _ := h.Repository.GetById(id, unitKeyPrefix)

	log.Println("edit single:", id)

	categories := make(map[string]string)
	if values, err := h.Repository.GetAll(categoryKeyPrefix); err == nil {
		categories = make(map[string]string, len(values))

		for _, v := range values {
			if cat, err := convertCategoryFromJson([]byte(v)); err == nil {
				categories[cat.Name] = cat.Type
			}

		}
	}

	if unit, err := convertUnitFromJson([]byte(unitJson)); err == nil {
		return c.Render(http.StatusOK, "updateUnit.html", map[string]interface{}{
			"Unit":       unit,
			"categories": categories,
		})
	}

	return h.GetUnits(c)
}

func (h Handler) PostUnit(c echo.Context) error {
	u := new(Unit)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u.Id = uuid.New().String()
	if err := h.Repository.Persist(u, unitKeyPrefix); err != nil {
		log.Println("Error on persisting unit: ", err.Error())
		return c.Render(http.StatusInternalServerError, unitTemplate, nil)
	}

	return h.GetUnits(c)
}

func (h Handler) DeleteUnit(c echo.Context) error {
	log.Println("delete")
	return c.Render(http.StatusOK, unitTemplate, nil)
}

func (h Handler) PutUnit(c echo.Context) error {
	u := new(Unit)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//TODO VZ error handling 404
	unitJson, _ := h.Repository.GetById(c.Param("id"), unitKeyPrefix)

	if unit, err := convertUnitFromJson([]byte(unitJson)); err == nil {
		unit.Name = u.Name
		unit.EAN = u.EAN
		unit.Flavour = u.Flavour
		unit.CategoryName = u.CategoryName
		if err := h.Repository.Persist(unit, unitJson); err != nil {
			log.Println("error on cat persisting", err.Error())
		} else {
			log.Printf("updated %s", unit)
		}
	}

	return h.GetUnits(c)
}

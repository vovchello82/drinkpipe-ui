package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const mixtureTemplate = "mixture.html"

type Mixture struct {
	Piece
	Description string `json:"description" form:"description"`
	Ingredients []*Unit
}

func (m *Mixture) ConvertToJson() ([]byte, error) {
	value, err := json.Marshal(m)

	if err != nil {
		log.Println("marshal error")
		return nil, err
	}

	return value, nil
}

func convertMixtureFromJson(jsonString []byte) (*Mixture, error) {
	m := new(Mixture)
	if err := json.Unmarshal(jsonString, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (h Handler) GetEditMixture(c echo.Context) error {
	log.Println("get all")
	return c.Render(http.StatusOK, mixtureTemplate, nil)
}

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

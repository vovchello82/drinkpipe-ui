package main

import (
	"drinkpipe-ui/handler"
	"io"
	"net/http"

	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	h := &handler.Handler{}
	e.Use(middleware.Logger())
	mix := e.Group("/mixture")
	mix.GET("", h.GetMixtures)
	mix.POST("", h.PostMixture)
	mix.GET("/:id", h.GetMixture)
	mix.PUT("/:id", h.PutMixture)

	category := e.Group("/category")
	category.GET("/edit", h.EditCategory)
	category.GET("", h.GetCategories)
	category.POST("", h.PostCategory)
	category.GET("/:id", h.GetCategory)
	category.PUT("/:id", h.PutCategory)

	unit := e.Group("/unit")
	unit.GET("", h.GetUnits)
	unit.POST("", h.PostUnit)
	unit.GET("/:id", h.GetUnit)
	unit.PUT("/:id", h.PutUnit)

	e.Logger.Fatal(e.Start(":1323"))

}

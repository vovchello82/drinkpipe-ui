package main

import (
	"drinkpipe-ui/handler"
	"drinkpipe-ui/store"
	"io"
	"log"
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

	store, err := store.NewStoreRedis()

	if err != nil {
		log.Fatalln("redis store unavaiable")
	}
	h := &handler.Handler{
		Repository: store,
	}

	e.Use(middleware.Logger())
	main := e.Group("/dp")
	mix := main.Group("/mixture")
	mix.GET("", h.GetMixtures)
	mix.POST("", h.PostMixture)
	mix.GET("/:id", h.GetMixture)
	mix.PUT("/:id", h.PutMixture)

	category := main.Group("/category")
	category.GET("", h.GetCategories)
	category.GET("/new", h.GetNewCategory)
	category.POST("", h.PostCategory)
	category.GET("/:id", h.GetCategory)
	category.GET("/:id/edit", h.GetEditCategory)
	category.PUT("/:id", h.PutCategory)
	category.POST("/:id", h.PutCategory)

	unit := main.Group("/unit")
	unit.GET("/new", h.GetUpdateUnit)
	unit.GET("", h.GetUnits)
	unit.POST("", h.PostUnit)
	unit.GET("/:id", h.GetUnit)
	unit.PUT("/:id", h.PutUnit)

	e.Logger.Fatal(e.Start(":1323"))

}

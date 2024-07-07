package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(
	writer io.Writer,
	name string,
	data interface{},
	context echo.Context,
) error {
	return t.templates.ExecuteTemplate(writer, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	echo := echo.New()
	echo.Use(middleware.Logger())

	echo.Renderer = newTemplate()

	echo.GET("/", func(context echo.Context) error {

	})
}

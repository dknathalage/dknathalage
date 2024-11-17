package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/kp-wedding/*.html")),
	}

	e.Renderer = t

	e.GET("/", HandleIndex)

	e.Logger.Fatal(e.Start(":80"))
}

func HandleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "hello")
}

package main

import (
	"html/template"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoTemplate struct {
	echoTemplates *template.Template
}

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.echoTemplates.ExecuteTemplate(w, name, data)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &EchoTemplate{
		echoTemplates: template.Must(template.ParseGlob("*.gohtml")),
	}

	e.GET("/", func(c echo.Context) error { return c.Render(200, "index", nil) })
	e.GET("/dostuff", func(c echo.Context) error {
		time.Sleep(1 * time.Second)
		return c.Render(200, "dostuff", c.QueryString())
	})

	e.Logger.Fatal(e.Start(":8080"))
}

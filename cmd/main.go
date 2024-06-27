package main

import (
	"io"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type Round struct {
    OpponentName string
    Win int
    Loss int
}

func newRound(opponentName string, win int, loss int) Round {
    return Round{
        OpponentName: opponentName,
        Win: win,
        Loss: loss,
    }
}

type Rounds = []Round

type Data struct {
    Rounds Rounds
}

func newData() Data {
    return Data{
        Rounds: []Round{
            newRound("test1", 10, 0),
            newRound("test2", 0, 10),
        },
    }
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    data := newData()
    e.Renderer = newTemplate()

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })
    e.Logger.Fatal(e.Start(":2000"))
}


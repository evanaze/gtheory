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

type Battle struct {
    OpponentName string
    Score int
    ScoreHistory []int
}

type Battles = []Battle

type Data struct {
    Agent Actor
    Battles Battles
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    var agent quidProQuo
    var opponent alwaysCheat
    params := BattleParams{
        CooperateReward: 1,
        CheatReward: 2,
        Opacity: 0.05,
        NRounds: 200,
    }

    results := battle(agent, opponent, params)
    battle_results := Battle{
        OpponentName: opponent.name,
        Score: results.Agent1ScoreTotal,
        ScoreHistory: results.Agent1ScoreHistory,
    }

    battles := make([]Battle, 1)
    battles[0] = battle_results
    data := Data{
        Agent: agent,
        Battles: battles,
    }
    e.Renderer = newTemplate()

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })
    e.Logger.Fatal(e.Start(":2000"))
}


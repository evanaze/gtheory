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

func runBattles(agent Actor, opponents []Actor, params BattleParams) Battles {
    var nOpponents int = len(opponents)
    var battleResults = make([]Battle, nOpponents)
    for i, opponent := range opponents {
        result := battle(agent, opponent, params)
        battleResult := Battle{
            OpponentName: opponent.Name,
            Score: result.Agent1ScoreTotal,
            ScoreHistory: result.Agent1ScoreHistory,
        }
        battleResults[i] = battleResult
    }
    return battleResults
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    params := BattleParams{
        CooperateReward: 1,
        CheatReward: 2,
        Opacity: 0.05,
        NRounds: 200,
    }

    var agent = quidProQuo{"Quid Pro Quo"}
    var opponents = make([]Actor, 5)
    opponents[0] = alwaysCheat{"Always Cheat"}
    opponents[1] = alwaysCooperate{"Always Cooperate"}
    opponents[2] = copycat{"Copycat"}
    opponents[3] = grudger{"Grudger"}
    opponents[4] = detective{"Detective"}

    battleResults := runBattles(agent, opponents, params)

    data := Data{
        Agent: agent,
        Battles: battleResults,
    }
    e.Renderer = newTemplate()

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })
    e.Logger.Fatal(e.Start(":2000"))
}


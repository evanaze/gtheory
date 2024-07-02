package main


type Move int8

const (
    Cheat Move = iota
    Cooperate
)

type cpuAgent struct {
    name string
}

type agent interface {
    act() Move
}

type BattleResults struct {
    TotalRounds int
    AgentName string
    AgentWins int
    AgentWinPct float32
    OpponentName string
    OpponentWins int
    OpponentWinPct float32
}

func battle(agent1, agent2 agent, nrounds int) BattleResults {
    for round:=0; round<nrounds; round++ {
        println()
    }
    return BattleResults{
        TotalRounds: nrounds,
        AgentName: agent1.Name,
    }
}


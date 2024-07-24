package main

import (
	"errors"
	"math/rand"
)

type Move int8
type MoveHistory []Move

const (
	Cheat Move = iota
	Cooperate
)

func inverseMove(move Move) Move {
	if move == Cheat {
		return Cooperate
	} else {
		return Cheat
	}
}

type Agent struct {
	Name string
}

// An interface to the action of an Agent
type Actor interface {
	act(agentHistory, oppHistory MoveHistory) (Move, error)
	name() string
}

// Define different default agents
type quidProQuo Agent
type alwaysCooperate Agent
type alwaysCheat Agent
type copycat Agent
type grudger Agent
type detective Agent

func (a quidProQuo) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	var nHistory int = len(oppHistory)
	if nHistory == 0 {
		return Cooperate, nil
	}

	switch mostRecent := oppHistory[nHistory-1]; mostRecent {
	case Cheat:
		return Cheat, nil
	case Cooperate:
		return Cooperate, nil
	default:
		return 0, errors.New("invalid move in move history")
	}
}

// An agent that always cooperates
func (a alwaysCooperate) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	return Cooperate, nil
}

// An agent that always cheats
func (a alwaysCheat) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	return Cheat, nil
}

// A copycat agent
func (a copycat) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	return Cheat, nil
}

// A grudger agent
func (a grudger) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	return Cheat, nil
}

// A detective agent
func (a detective) act(agentHistory, oppHistory MoveHistory) (Move, error) {
	return Cheat, nil
}

func (a quidProQuo) name() string      { return a.Name }
func (a alwaysCooperate) name() string { return a.Name }
func (a alwaysCheat) name() string     { return a.Name }
func (a copycat) name() string         { return a.Name }
func (a grudger) name() string         { return a.Name }
func (a detective) name() string       { return a.Name }

type BattleResults struct {
	Agent1             Actor
	Agent1MoveHistory  []Move
	Agent1ScoreHistory []int
	Agent1ScoreTotal   int
	Agent2             Actor
	Agent2MoveHistory  []Move
	Agent2ScoreHistory []int
	Agent2ScoreTotal   int
}

type BattleParams struct {
	CooperateReward int
	CheatReward     int
	Opacity         float32
	NRounds         int
}

// Calculate the score of the round for each agent
func score(agent1Move, agent2Move Move, params BattleParams) (agent1Score, agent2Score int) {
	switch {
	case agent1Move == Cooperate && agent2Move == Cooperate:
		output := 1 + params.CooperateReward
		return output, output
	case agent1Move == Cooperate && agent2Move == Cheat:
		return -1, 1 + params.CheatReward
	case agent1Move == Cheat && agent2Move == Cooperate:
		return 1 + params.CheatReward, -1
	default:
		return 0, 0
	}
}

// Roll the possibility that the intended move gets obfuscated
func rollObfuscate(inputMove Move, p float32) Move {
	var roll = rand.Float32()
	if roll > p {
		return inputMove
	} else {
		var outputMove = inverseMove(inputMove)
		return outputMove
	}
}

// Execute the battle
func battle(agent1, agent2 Actor, params BattleParams) BattleResults {
	var agent1MoveHistory []Move
	var agent1ScoreHistory []int
	var agent1ScoreTotal int = 0

	var agent2MoveHistory []Move
	var agent2ScoreHistory []int
	var agent2ScoreTotal int = 0

	for round := 0; round < params.NRounds; round++ {
		agent1Move, _ := agent1.act(agent1MoveHistory, agent2MoveHistory)
		agent1Move = rollObfuscate(agent1Move, params.Opacity)
		agent1MoveHistory = append(agent1MoveHistory, agent1Move)

		agent2Move, _ := agent2.act(agent2MoveHistory, agent1MoveHistory)
		agent2Move = rollObfuscate(agent2Move, params.Opacity)
		agent2MoveHistory = append(agent1MoveHistory, agent2Move)

		agent1Score, agent2Score := score(agent1Move, agent2Move, params)
		agent1ScoreHistory = append(agent1ScoreHistory, agent1Score)
		agent1ScoreTotal += agent1Score

		agent2ScoreHistory = append(agent2ScoreHistory, agent2Score)
		agent2ScoreTotal += agent2Score
	}

	return BattleResults{
		Agent1:             agent1,
		Agent1MoveHistory:  agent1MoveHistory,
		Agent1ScoreHistory: agent1ScoreHistory,
		Agent1ScoreTotal:   agent1ScoreTotal,
		Agent2:             agent2,
		Agent2MoveHistory:  agent2MoveHistory,
		Agent2ScoreHistory: agent2ScoreHistory,
		Agent2ScoreTotal:   agent2ScoreTotal,
	}
}

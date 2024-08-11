package parser

import "sort"

type Game struct {
	Matches []*GameMatch
	Kills   map[string]int
}

type Player struct {
	Name  string
	Kills int
}

func NewGame() *Game {
	return &Game{
		Kills: map[string]int{},
	}
}

func (g *Game) IncrementKills(playerName string) {
	g.Kills[playerName]++
}

func (g *Game) GetRanking() []*Player {
	var players []*Player
	for name, score := range g.Kills {
		players = append(players, &Player{Name: name, Kills: score})
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Kills > players[j].Kills
	})

	return players
}

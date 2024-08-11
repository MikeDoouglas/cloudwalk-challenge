package parser

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesToString(t *testing.T) {
	matches := []*GameMatch{
		{
			Players: map[string]bool{"Player1": true, "Player2": true},
			Kills:   map[string]int{"Player1": 3, "Player2": 2},
			DeathType: map[DeathType]int{
				"MOD_RAILGUN":    4,
				"MOD_MACHINEGUN": 1,
			},
		},
	}

	report := &Report{Matches: matches}

	result, err := report.MatchesToString()
	assert.NoError(t, err)

	var resultJson []MatchJson
	err = json.Unmarshal([]byte(result), &resultJson)
	assert.NoError(t, err)

	assert.Len(t, resultJson, 1, "Expected 1 match")

	match := resultJson[0]
	expectedTotalKills := 5
	assert.Equal(t, expectedTotalKills, match.TotalKills, "Expected TotalKills to be %d", expectedTotalKills)
	assert.Len(t, match.Players, 2, "Expected 2 players")
	assert.Equal(t, 3, match.Kills["Player1"], "Expected Kills for Player1 to be 3")
	assert.Equal(t, 2, match.Kills["Player2"], "Expected Kills for Player2 to be 2")
}

func TestRankingToString(t *testing.T) {
	players := []*Player{
		{Name: "Player1", Kills: 10},
		{Name: "Player2", Kills: 5},
		{Name: "Player3", Kills: 3},
	}

	report := &Report{Ranking: players}
	result := report.RankingToString()

	expected := "1. Player1 (10)\n2. Player2 (5)\n3. Player3 (3)\n"
	assert.Equal(t, expected, result, "Expected:\n%s\nGot:\n%s", expected, result)
}

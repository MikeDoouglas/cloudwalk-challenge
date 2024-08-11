package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Report struct {
	Ranking []*Player
	Matches []*GameMatch
}

type MatchJson struct {
	TotalKills int               `json:"total_kills"`
	Players    []string          `json:"players"`
	Kills      map[string]int    `json:"kills"`
	DeathType  map[DeathType]int `json:"kills_by_means"`
}

func (r *Report) MatchesToString() (string, error) {
	matchesJson := make([]*MatchJson, 0, len(r.Matches))
	for _, m := range r.Matches {
		matchesJson = append(matchesJson, &MatchJson{
			TotalKills: m.GetTotalKills(),
			Players:    m.GetPlayerNames(),
			Kills:      m.Kills,
			DeathType:  m.DeathType,
		})
	}

	v, err := json.MarshalIndent(matchesJson, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to parse matches to string: %w", err)
	}

	return string(v), nil
}

func (r *Report) RankingToString() string {
	names := make([]string, 0, len(r.Ranking))
	scores := make([]int, 0, len(r.Ranking))
	for _, player := range r.Ranking {
		names = append(names, player.Name)
		scores = append(scores, player.Kills)
	}

	var builder strings.Builder
	counter := 1
	for i := range names {
		builder.WriteString(fmt.Sprintf("%d. %s (%d)\n", counter, names[i], scores[i]))
		counter++
	}

	return builder.String()
}

func (r *Report) PrintReport() error {
	matches, err := r.MatchesToString()
	if err != nil {
		return err
	}

	var builder strings.Builder
	builder.WriteString("------------- RANKING -------------\n")
	builder.WriteString(r.RankingToString())
	builder.WriteString("\n")
	builder.WriteString("------------- MATCHES -------------\n")
	builder.WriteString(matches)
	fmt.Println(builder.String())

	return nil
}

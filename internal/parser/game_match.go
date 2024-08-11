package parser

type GameMatch struct {
	Players    map[string]bool
	Kills      map[string]int
	WorldKills map[string]int
	DeathType  map[DeathType]int
}

func NewGameMatch() *GameMatch {
	return &GameMatch{
		Players:    map[string]bool{},
		Kills:      map[string]int{},
		WorldKills: map[string]int{},
		DeathType:  map[DeathType]int{},
	}
}

func (g *GameMatch) GetTotalKills() int {
	total := 0
	for i := range g.Kills {
		total += g.Kills[i]
	}

	for i := range g.WorldKills {
		total += g.WorldKills[i]
	}

	return total
}

func (g *GameMatch) GetPlayerNames() []string {
	names := make([]string, 0, len(g.Players))
	for name := range g.Players {
		names = append(names, name)
	}

	return names
}

func (g *GameMatch) IncrementWorldKill(killedKey string) {
	g.WorldKills[killedKey]++
}

func (g *GameMatch) IncrementKill(killedKey string) {
	g.Kills[killedKey]++
}

func (g *GameMatch) IncrementKillMethods(deathTypeString string) {
	deathType := GetDeathType(deathTypeString)
	g.DeathType[deathType]++
}

func (g *GameMatch) AddPlayer(name string) {
	g.Players[name] = true
}

func (g *GameMatch) CalculateWorldKills() {
	for killed, times := range g.WorldKills {
		_, exists := g.Kills[killed]
		if exists {
			g.Kills[killed] -= times
			if g.Kills[killed] <= 0 {
				delete(g.Kills, killed)
			}
		}
	}
}

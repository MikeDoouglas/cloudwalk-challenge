package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

const (
	regex              = `.*Kill:\s+\d+\s+\d+\s+\d+:\s+(.+?)\s+killed\s+(.+?)\s+by\s+(.+)`
	lineIdentifier     = "Kill:"
	endMatchIdentifier = "ShutdownGame"
)

type Service struct {
	filePath string
	worldTag string
	log      *zap.Logger
}

type LogEvent struct {
	Killer    string
	Victim    string
	DeathType string
}

func NewService(filePath string, worldTag string, log *zap.Logger) *Service {
	return &Service{filePath: filePath, worldTag: worldTag, log: log}
}

func (s *Service) ParseFile() (*Game, error) {
	gameLogs, err := readFile(s.filePath)
	if err != nil {
		return nil, err
	}

	s.log.Debug("file has been read", zap.String("file", s.filePath), zap.Int("gameMatches", len(gameLogs)))

	game := NewGame()
	for i := range gameLogs {
		gameMatch := NewGameMatch()
		for _, line := range gameLogs[i] {
			event := s.extractLogEvent(line)
			if event == nil {
				s.log.Warn("failed to match regex on text",
					zap.String("text", line),
					zap.String("regexPattern", regex))
				continue
			}

			if event.Killer == s.worldTag {
				gameMatch.IncrementWorldKill(event.Victim)
			} else {
				gameMatch.AddPlayer(event.Killer)
				gameMatch.IncrementKill(event.Killer)
				game.IncrementKills(event.Killer)
			}
			gameMatch.AddPlayer(event.Victim)
			gameMatch.IncrementKillMethods(event.DeathType)
		}

		gameMatch.CalculateWorldKills()
		game.Matches = append(game.Matches, gameMatch)
	}

	return game, nil
}

func (s *Service) extractLogEvent(text string) *LogEvent {
	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(text)

	if len(matches) == 4 {
		return &LogEvent{
			Killer:    matches[1],
			Victim:    matches[2],
			DeathType: matches[3],
		}
	}

	return nil
}

func readFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]string{}, fmt.Errorf("failed opening file: %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allLines := make([][]string, 0)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, lineIdentifier) {
			lines = append(lines, line)
		}

		if strings.Contains(line, endMatchIdentifier) && len(lines) > 0 {
			allLines = append(allLines, lines)
			lines = []string{}
		}
	}

	return allLines, nil
}

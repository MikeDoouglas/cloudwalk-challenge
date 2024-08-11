package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestExtractLogEvent(t *testing.T) {
	tests := []struct {
		input    string
		expected *LogEvent
	}{
		{
			input: "Kill: 1022 4 20: Player1 killed Player2 by MOD_RAILGUN",
			expected: &LogEvent{
				Killer:    "Player1",
				Victim:    "Player2",
				DeathType: "MOD_RAILGUN",
			},
		},
		{
			input:    "Invalid log line",
			expected: nil,
		},
	}

	s := &Service{}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := s.extractLogEvent(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestReadFile(t *testing.T) {
	fileContent := `
Kill: 1022 4 20: Player1 killed Player2 by MOD_RAILGUN
Kill: 1022 5 21: Player3 killed Player4 by MOD_MACHINEGUN
ShutdownGame:
Kill: 1022 6 22: Player5 killed Player6 by MOD_RAILGUN
ShutdownGame:
`
	tmpFile, err := os.CreateTemp("", "testlog.log")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(fileContent)
	assert.NoError(t, err)
	assert.NoError(t, tmpFile.Close())

	result, err := readFile(tmpFile.Name())
	assert.NoError(t, err)

	expectedLength := 2
	assert.Equal(t, expectedLength, len(result))

	assert.Equal(t, 2, len(result[0]))
	assert.Equal(t, 1, len(result[1]))
}

func TestParseFile(t *testing.T) {
	logger := zaptest.NewLogger(t)

	fileContent := `
Kill: 1022 4 20: Player1 killed Player2 by MOD_RAILGUN
Kill: 1022 5 21: Player3 killed Player4 by MOD_MACHINEGUN
ShutdownGame:
Kill: 1022 6 22: Player5 killed Player6 by MOD_RAILGUN
ShutdownGame:
`
	tmpFile, err := os.CreateTemp("", "testlog_*.log")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(fileContent)
	assert.NoError(t, err)
	assert.NoError(t, tmpFile.Close())

	s := NewService(tmpFile.Name(), "world", logger)
	game, err := s.ParseFile()
	assert.NoError(t, err)

	expectedMatches := 2
	assert.Equal(t, expectedMatches, len(game.Matches))
}

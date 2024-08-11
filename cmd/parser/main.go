package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mikedoouglas/cloudwalk-challenge/internal/parser"
	"go.uber.org/zap"
)

const filesDirectory = "files"

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("fail to initialize zap logger: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	fileName := os.Getenv("FILE_NAME")
	worldTag := os.Getenv("WORLD_TAG")

	filePath := fmt.Sprintf("%s/%s", filesDirectory, fileName)
	service := parser.NewService(filePath, worldTag, logger)
	game, err := service.ParseFile()
	if err != nil {
		fmt.Println("failed to parse file: %w", err)
	}

	report := &parser.Report{Ranking: game.GetRanking(), Matches: game.Matches}
	if err := report.PrintReport(); err != nil {
		fmt.Println("failed to report: %w", err)
	} else {
		logger.Info("Report has been successfully generated!")
	}
}

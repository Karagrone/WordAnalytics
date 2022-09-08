package main

import (
	"Projects/WordAnalytics/internal/telegram"
	"Projects/WordAnalytics/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	telegram.Bot()
	log.Info("Bot started")
}

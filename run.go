package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main/handlers"
	"os"
)

func connect(debug bool) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = debug
	log.Printf("Successfully authorized on account @%s", bot.Self.UserName)

	return bot
}

func getBotActions() handlers.ActiveHandlers {
	actions := handlers.ActiveHandlers{Handlers: []handlers.Handler{
		// Place your handlers here
	}}

	return actions
}

func main() {
	client := connect(true)
	actions := getBotActions()

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := client.GetUpdatesChan(updateConfig)
	for update := range updates {
		runRes := actions.HandleAll(update)
		fmt.Println("Run results: [ID|called]")
		fmt.Println(runRes)
	}
}

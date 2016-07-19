package main

import (
	"encoding/json"
	"fmt"
	"github.com/requilence/integram"
	"github.com/requilence/integram/services/trello"
	"os"
)

func main() {

	type Config struct {
		Debug         bool
		TrelloKey     string
		TrelloSecret  string
		TelegramToken string
	}

	configFile, _ := os.Open("config.json")
	decoder := json.NewDecoder(configFile)
	config := Config{}

	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("error:", err)
	}

	integram.Debug = config.Debug

	integram.Register(
		trello.Config{
			integram.OAuthProvider{
				ID:     config.TrelloKey,
				Secret: config.TrelloSecret,
			},
		},
		config.TelegramToken,
	)

	integram.Run()
}

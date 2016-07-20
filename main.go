package main

import (
	"encoding/json"
	"fmt"
	"github.com/requilence/integram"
	"github.com/requilence/integram/services/trello"
	"os"
)

type Config struct {
	Debug         bool
	TrelloKey     string
	TrelloSecret  string
	TelegramToken string
}

func loadConfig(filePath string) Config {

	configFile, _ := os.Open(filePath)
	decoder := json.NewDecoder(configFile)
	config := Config{}

	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("error:", err)
	}

	return config

}

func main() {

	confPath := os.Args[1]

	config := loadConfig(confPath)

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

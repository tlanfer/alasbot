package alasbot

import "fmt"

type Bot struct {
	Game Game
	Chat Chat
}

type Game interface {
	PlayerCount() (int, int, error)
}

type Chat interface {
	OnMessage(func(author, message string) string)
}

func (bot Bot) Start() {

	bot.Chat.OnMessage(func(author, message string) string {
		if message == "!server" {
			count, max, err := bot.Game.PlayerCount()
			if err != nil {
				return "Sorry, could not get player count: " + err.Error()
			}

			return fmt.Sprintf("There are %v/%v players connected", count, max)
		}
		return ""
	})
}

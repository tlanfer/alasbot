package alasbot

import (
	"fmt"
	"math"
	"time"
)

type Bot struct {
	Game Game
	Chat Chat
}

type Game interface {
	PlayerCount() (int, int, error)
	GameTime() (time.Duration, error)
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

			gameTime, err := bot.Game.GameTime()
			if err != nil {
				return "Sorry, could not get server time: " + err.Error()
			}

			days := math.Floor(gameTime.Hours()/24)
			time := gameTime - time.Duration(days)*24*time.Hour

			return fmt.Sprintf("There are %v/%v players connected. It is day %v, the time is %v", count, max, days, time)
		}
		return ""
	})
}

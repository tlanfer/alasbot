package alasbot

import (
	"fmt"
)

type Bot struct {
	Game Game
	Chat Chat
}

type Game interface {
	PlayerCount() (int, int, error)
	GameTime() (int, int, int, error)
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

			days, hours, minutes, err := bot.Game.GameTime()
			if err != nil {
				return "Sorry, could not get server time: " + err.Error()
			}

			nextBloodMoon := ((days/7)+1)*7

			return fmt.Sprintf("There are %v/%v players connected. Its day %v, the time is %v:%v. The next bloodmoon will be on day %v", count, max, days, hours, minutes, nextBloodMoon)
		}
		return ""
	})
}

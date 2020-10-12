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
	GameTime() (int, int, int, int, error)
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

			days, hours, minutes, bloodMoonFrequency, err := bot.Game.GameTime()
			if err != nil {
				return "Sorry, could not get server time: " + err.Error()
			}

			playersMessage := fmt.Sprintf("Server is online, %v/%v players.", count, max)
			timeMessage := fmt.Sprintf("Its day %v, the time is %02d:%02d.", days, hours, minutes)

			bloodMoonMessage := bloodMoonMessage(days, hours, minutes, bloodMoonFrequency)

			return fmt.Sprint( playersMessage, " ", timeMessage, " ", bloodMoonMessage)
		}
		return ""
	})
}

func bloodMoonMessage(days, hours, minutes, bloodMoonFrequency int) string {
	msg := fmt.Sprintf("The next bloodmoon will be on day %v.",((days/bloodMoonFrequency)+1)*bloodMoonFrequency)

	if (days % 7) == 0 {
		msg = "The next bloodmoon will be today."

		if hours >= 22 {
			msg =  "A bloodmoon is active!"
		}
	}

	if (days % 7) == 1 {
		if hours < 4 {
			msg =  "A bloodmoon is active!"
		}
	}


	return msg
}
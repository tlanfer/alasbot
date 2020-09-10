package main

import (
	"fmt"
	"github.com/tlanfer/alasbot"
	"github.com/tlanfer/alasbot/plugins/discord"
	"github.com/tlanfer/alasbot/plugins/sevendays"
	"os"
	"time"
)

var (
	sevenDaysServer = os.Getenv("SEVEN_DAYS_SERVER")
	botToken        = os.Getenv("BOT_TOKEN")
)
func main() {

	if sevenDaysServer == "" {
		panic("SEVEN_DAYS_SERVER missing")
	}

	if botToken == "" {
		panic("BOT_TOKEN missing")
	}


	game := sevendays.New(sevenDaysServer)
	chat, err := discord.New(botToken)

	if err != nil {
		panic(err)
	}

	bot := alasbot.Bot{
		Game: game,
		Chat: chat,
	}

	bot.Start()
	for _ = range time.Tick(30 * time.Second) {
		fmt.Println("still alive...")
	}
}

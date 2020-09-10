package main

import (
	"fmt"
	"github.com/tlanfer/alasbot"
	"github.com/tlanfer/alasbot/plugins/discord"
	"github.com/tlanfer/alasbot/plugins/sevendays"
	"os"
	"time"
)

func main() {

	game := sevendays.New(os.Getenv("SEVEN_DAYS_SERVER"))
	chat, _ := discord.New(os.Getenv("BOT_TOKEN"))

	bot := alasbot.Bot{
		Game: game,
		Chat: chat,
	}

	bot.Start()
	for _ = range time.Tick(30 * time.Second) {
		fmt.Println("still alive...")
	}
}

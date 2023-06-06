package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tlanfer/alasbot"
	"log"
)

func New(token string) (alasbot.Chat, error) {
	dg, err := discordgo.New(token)

	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return &discordClient{
		dg: dg,
	}, nil
}

type discordClient struct {
	dg *discordgo.Session
}

func (d *discordClient) OnMessage(f func(author, message string) string) {
	d.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		user := m.Author.Username
		message := m.Content

		response := f(user, message)

		if response != "" {
			_, err := s.ChannelMessageSend(m.ChannelID, response)

			if err != nil {
				log.Println("failed to respond:", err)
			}
		}
	})
}

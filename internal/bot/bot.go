package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	client *tgbotapi.BotAPI
}

func New(token string) (*Bot, error) {
	client, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("create telegram bot client: %w", err)
	}

	return &Bot{
		client: client,
	}, nil
}

func (b *Bot) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.client.GetUpdatesChan(u)

	log.Printf("Authorized on account %s", b.client.Self.UserName)

	for update := range updates {
		if err := b.handleUpdate(update); err != nil {
			log.Println(err)
		}
	}

	return nil
}

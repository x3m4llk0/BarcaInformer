package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleUpdate(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.Text = "Рад приветствовать тебя в этом боте"
			b.client.Send(msg)

		case "match":
			err := b.sendMatchMessage(update.Message.Chat.ID)
			if err != nil {
				return err
			}

		default:
			msg.Text = "I don't know that command"
			b.client.Send(msg)
		}

	}

	_ = update
	return nil
}

package bot

import (
	"BarcaInformer/internal/provider"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"html"
	"log"
	"strings"
	"time"
)

func (b *Bot) sendMatchMessage(chatID int64) error {
	matchInfo, err := provider.GetInfo()

	message, err := buildMatchMessage(matchInfo)
	if err != nil {
		log.Println(err)
	}
	msg := tgbotapi.NewMessage(chatID, "")

	msg.ParseMode = tgbotapi.ModeHTML

	msg.Text = message
	b.client.Send(msg)

	return nil
}

func buildMatchMessage(matchInfo provider.MatchInfo) (string, error) {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", fmt.Errorf("cannot load timezone Europe/Moscow: %w", err)
	}
	kickoffMSK := matchInfo.UtcDate.In(moscow)
	kickoff := kickoffMSK.Format("02.01.2006 15:04")

	var b strings.Builder
	b.WriteString("⚽ <b>Следующий матч Барселоны</b>\n")
	fmt.Fprintf(&b, "🕒 <b>Время (МСК):</b> <code>%s</code>\n", kickoff)
	fmt.Fprintf(&b, "🏠 <b>Дома:</b> %s\n", html.EscapeString(matchInfo.HomeTeam.Name))
	fmt.Fprintf(&b, "🚌 <b>Гости:</b> %s\n", html.EscapeString(matchInfo.AwayTeam.Name))
	fmt.Fprintf(&b, "🏆 <b>Турнир:</b> %s", html.EscapeString(matchInfo.Competition.Name))

	return b.String(), nil
}

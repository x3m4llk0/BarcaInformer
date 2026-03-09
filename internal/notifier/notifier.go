package notifier

import (
	"BarcaInformer/internal/provider"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"
)

type Notifier struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

func SendMessage(matchInfo provider.MatchInfo) error {

	apiToken := "1861694539:AAFCmesajfWX0289V_jxEJ1wlyfYhTv_-nI"
	chatID := "983413077"

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken)
	message, err := buildMessage(matchInfo)
	if err != nil {
		return err
	}

	body := Notifier{
		ChatID:                chatID,
		Text:                  message,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram status: %s: %s", resp.Status, strings.TrimSpace(string(respBody)))
	}

	return nil

}

func buildMessage(matchInfo provider.MatchInfo) (string, error) {
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

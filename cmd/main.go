package main

import (
	"BarcaInformer/internal/bot"
	"log"
)

func main() {
	telegramBot, err := bot.New("1861694539:AAFCmesajfWX0289V_jxEJ1wlyfYhTv_-nI")
	if err != nil {
		log.Fatal(err)
	}

	if err := telegramBot.Run(); err != nil {
		log.Fatal(err)
	}
}

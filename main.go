package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found\n")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	// Создание нового клиента Max Bot API
	api, _ := maxbot.New(os.Getenv("BOT_TOKEN"))
	
	info, _ := api.Bots.GetBot(ctx)
	log.Printf("Bot %s is launched!", info.Name)

	for upd := range api.GetUpdates(ctx) { // Чтение из канала с обновлениями
		timestamp := time.Now().UnixMilli()
		switch upd := upd.(type) { // Определение типа пришедшего обновления
		case *schemes.BotStartedUpdate, *schemes.MessageCreatedUpdate:
			// Формирование текста
			text := fmt.Sprintf(
				"UserId: <code>%s</code>",
				strconv.FormatInt(upd.GetUserID(), 10),
			)

			// Формирование сообщения
			message := maxbot.NewMessage().
				SetChat(upd.GetChatID()).
				SetText(text).
				SetFormat("html")

			// Отправка сообщения
			err := api.Messages.Send(ctx, message)
			if err != nil {
				log.Printf(
					"Error sending message (ChatId: %d; UserId: %d): %s",
					upd.GetChatID(), upd.GetUserID(), err.Error(),
				)
			}
		}
		log.Printf(
			"Update from user %d handeled in %d milliseconds",
			upd.GetUserID(),
			time.Now().UnixMilli()-timestamp,
		)
	}
}

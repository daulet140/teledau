package routes

import (
	"cwntelegram/internal/client"
	"cwntelegram/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Router struct {
	TelegramClient client.TelegramClient
}

func NewRouter(TelegramClient client.TelegramClient) Router {
	return Router{
		TelegramClient: TelegramClient,
	}
}
func (rtr *Router) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Received webhook payload:", string(body))

	var update models.Update

	err = json.Unmarshal(body, &update)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
		return
	}

	if update.Message != nil {
		fmt.Println("Received message:", update.Message.Text)
		rtr.textMessageHandler(update.Message)
	}

	if update.ChatMember != nil {
		fmt.Println("Received chat member:", update.ChatMember.NewChatMember.Status)
		rtr.chatMemberHandler(update.ChatMember)
	}
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Received"))

}

func (rtr *Router) textMessageHandler(message *models.Message) {
	if message.Text == "/start" {
		rtr.TelegramClient.SendMessage(models.MessageRequest{
			ChatId: message.Chat.Id,
			Text:   `Байқауға қатысу үшін "Қатысу" батырмасын басыңыз`,
			ReplyMarkup: models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{Text: "Қатысу"},
					},
				},
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
			},
		})
	}
	if message.Text == "Қатысу" {

		rtr.TelegramClient.GenerateInviteLinks(models.CreateChatInviteLinkRequest{
			ChatID: "qazaqibol",
			Name:   message.Chat.Username,
		})

		rtr.TelegramClient.SendMessage(models.MessageRequest{
			ChatId: message.Chat.Id,
			Text:   `Сіздің байқауға қатысу үшін сілтемеңіз"`,
			ReplyMarkup: models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{Text: "Статистика"},
					},
				},
				ResizeKeyboard: true,
			},
		})
	}
	if message.Text == "Статистика" {
		rtr.TelegramClient.SendMessage(models.MessageRequest{
			ChatId: message.Chat.Id,
			Text:   `Сіздің байқауға қатысу үшін сілтемеңіз"`,
			ReplyMarkup: models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{Text: "Статистика"},
					},
				},
				ResizeKeyboard: true,
			},
		})
	}
}

func (rtr *Router) chatMemberHandler(member *models.ChatMember) {

}

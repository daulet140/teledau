package teledau

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

var telegramClient *TelegramClient

func TestNewTelegramClient(t *testing.T) {
	telegramClient = NewTelegramClient("6699186697:AAFntrjkj36iY-l5ZcqEBxAvf40sSiGCzpk", context.Background())
	if telegramClient == nil {
		t.Error("Telegram client is nil")
	}
}

func TestTelegramClient_SendMessage(t *testing.T) {
	telegramClient = NewTelegramClient("6699186697:AAFntrjkj36iY-l5ZcqEBxAvf40sSiGCzpk", context.Background())
	if telegramClient == nil {
		t.Error("Telegram client is nil")
	}
	_, err := telegramClient.SendMessage(MessageRequest{
		ChatId:    75504797,
		Text:      "_Test message_",
		ParseMode: "MarkdownV2",
	})
	if err != nil {
		t.Error(err)
	}
	_, err = telegramClient.SendMessage(MessageRequest{
		ChatId:    75504797,
		Text:      "<b>Test message</b>",
		ParseMode: "HTML",
	})
	if err != nil {
		t.Error(err)
	}

	_, err = telegramClient.SendMessage(MessageRequest{
		ChatId: 75504797,
		Text:   "Test message",
		ReplyMarkup: InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{Text: "Option 1", CallbackData: "option1"},
					{Text: "Option 2", CallbackData: "option2"},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	_, err = telegramClient.SendMessage(MessageRequest{
		ChatId: 75504797,
		Text:   "Test message",
		ReplyMarkup: ReplyKeyboardMarkup{
			Keyboard: [][]KeyboardButton{
				{
					{Text: "Option 1"},
				},
			},
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		},
	})
	if err != nil {
		t.Error(err)
	}

}

func TestTelegramClient_SendSticker(t *testing.T) {
	telegramClient = NewTelegramClient("6699186697:AAFntrjkj36iY-l5ZcqEBxAvf40sSiGCzpk", context.Background())
	if telegramClient == nil {
		t.Error("Telegram client is nil")
	}
	//read file 1_1.webp and convert to base64 string
	bytes, err := ioutil.ReadFile("..\\..\\1_1.webp")
	if err != nil {
		t.Error(err)
	}

	base64String := base64.StdEncoding.EncodeToString(bytes)
	_, err = telegramClient.SendSticker(75504797, base64String)
	if err != nil {
		t.Error(err)
	}
}

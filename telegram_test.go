package teledau

import (
	"context"
	"log"
	"testing"
)

var telegramClient *TelegramClient

//	func TestNewTelegramClient(t *testing.T) {
//		telegramClient = NewTelegramClient("6699186697:-", context.Background())
//		if telegramClient == nil {
//			t.Error("Telegram client is nil")
//		}
//	}
//
//	func TestTelegramClient_SendMessage(t *testing.T) {
//		telegramClient = NewTelegramClient("6699186697:-l5ZcqEBxAvf40sSiGCzpk", context.Background())
//		if telegramClient == nil {
//			t.Error("Telegram client is nil")
//		}
//		_, err := telegramClient.SendMessage(MessageRequest{
//			ChatId:    "75504797",
//			Text:      "_Test message_",
//			ParseMode: "MarkdownV2",
//		})
//		if err != nil {
//			t.Error(err)
//		}
//		_, err = telegramClient.SendMessage(MessageRequest{
//			ChatId:    75504797,
//			Text:      "<b>Test message</b>",
//			ParseMode: "HTML",
//		})
//		if err != nil {
//			t.Error(err)
//		}
//
//		_, err = telegramClient.SendMessage(MessageRequest{
//			ChatId: 75504797,
//			Text:   "Test message",
//			ReplyMarkup: InlineKeyboardMarkup{
//				InlineKeyboard: [][]InlineKeyboardButton{
//					{
//						{Text: "Option 1", CallbackData: "option1"},
//						{Text: "Option 2", CallbackData: "option2"},
//					},
//				},
//			},
//		})
//		if err != nil {
//			t.Error(err)
//		}
//		_, err = telegramClient.SendMessage(MessageRequest{
//			ChatId: 75504797,
//			Text:   "Test message",
//			ReplyMarkup: ReplyKeyboardMarkup{
//				Keyboard: [][]KeyboardButton{
//					{
//						{Text: "Option 1"},
//					},
//				},
//				ResizeKeyboard:  true,
//				OneTimeKeyboard: true,
//			},
//		})
//		if err != nil {
//			t.Error(err)
//		}
//
// }
//
//	func TestTelegramClient_SendSticker(t *testing.T) {
//		telegramClient = NewTelegramClient("6699186697:-l5ZcqEBxAvf40sSiGCzpk", context.Background())
//		if telegramClient == nil {
//			t.Error("Telegram client is nil")
//		}
//		//read file 1_1.webp and convert to base64 string
//		bytes, err := ioutil.ReadFile("..\\..\\1_1.webp")
//		if err != nil {
//			t.Error(err)
//		}
//
//		base64String := base64.StdEncoding.EncodeToString(bytes)
//		_, err = telegramClient.SendSticker(75504797, base64String)
//		if err != nil {
//			t.Error(err)
//		}
//	}
//func TestTelegramClient_SendMedia(t *testing.T) {
//	telegramClient = NewTelegramClient(":", context.Background())
//	if telegramClient == nil {
//		t.Error("Telegram client is nil")
//	}
//	//read file 1_1.webp and convert to base64 string
//	media, err := getImageBase64FromURL("https://pbs.twimg.com/media/GQ0qcymXsAAbLcU?format=jpg&name=small")
//
//	resp, err := telegramClient.SendMedia("@", media, "*Test message* \n_test_", "MarkdownV2")
//	if err != nil {
//		t.Error(err)
//	}
//	log.Printf("%v", resp.Result.MessageId)
//}
//
//func getImageBase64FromURL(url string) (string, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("failed to fetch image: %s", resp.Status)
//	}
//
//	imageBytes, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
//	return base64Image, nil
//}

func TestTelegramClient_GetChat(t *testing.T) {
	telegramClient = NewTelegramClient("", context.Background())
	if telegramClient == nil {
		t.Error("Telegram client is nil")
	}
	chat, err := telegramClient.GetChat("@")
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", chat)

}

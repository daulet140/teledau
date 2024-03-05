package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"cwntelegram/internal/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Telegram interface {
	SendMessage(message models.MessageRequest) (models.Message, error)
	//SendMessageWithImg(chatId string, message models.MessageRequest, img []string)

	//SendPost(chatId string, title string, text string, img []string)

	SendPoll(poolRequest models.PollRequest) (models.PollResponse, error)

	SendMedia(chatId string, media string)

	SendSticker(chatId int64, media string) (models.StikerResponse, error)

	GenerateInviteLinks(invite models.CreateChatInviteLinkRequest) ([]string, error)
}

type TelegramClient struct {
	BotToken   string
	HttpClient http.Client
	Ctx        context.Context
}

func NewTelegramClient(botToken string, ctx context.Context) *TelegramClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &TelegramClient{
		BotToken: botToken,
		HttpClient: http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		},
		Ctx: ctx,
	}
}

func (t *TelegramClient) SendMessage(message models.MessageRequest) (models.Message, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendMessage"

	messageData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return models.Message{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(messageData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return models.Message{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return models.Message{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return models.Message{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return models.Message{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var createdApplicant models.Message
	if err := json.Unmarshal(bodyBytes, &createdApplicant); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return models.Message{}, err
	}

	return createdApplicant, nil
}

func (t *TelegramClient) SendPoll(poolRequest models.PollRequest) (models.PollResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendPoll"

	pollData, err := json.Marshal(poolRequest)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return models.PollResponse{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(pollData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return models.PollResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return models.PollResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return models.PollResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return models.PollResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var pollResponse models.PollResponse
	if err := json.Unmarshal(bodyBytes, &pollResponse); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)

		return models.PollResponse{}, err
	}

	return pollResponse, nil
}

func (t *TelegramClient) SendMedia(chatId string, media string) {

}

func (t *TelegramClient) SendSticker(chatId int64, media string) (models.StikerResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendSticker"
	filePath := "/path/to/decoded/sticker.webp"

	imageBytes, err := base64.StdEncoding.DecodeString(media)
	if err != nil {
		fmt.Printf("err: %v", err)
		return models.StikerResponse{}, err
	}

	var buffer bytes.Buffer
	_, err = buffer.Write(imageBytes)
	if err != nil {
		return models.StikerResponse{}, err
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part, err := writer.CreateFormFile("sticker", filepath.Base(filePath))
	if err != nil {
		log.Printf("create form file err: %v", err)
		return models.StikerResponse{}, err
	}

	_, err = io.Copy(part, &buffer)
	if err != nil {
		log.Printf("io copy err: %v", err)
		return models.StikerResponse{}, err
	}

	_ = writer.WriteField("chat_id", strconv.FormatInt(chatId, 10))
	writer.Close()

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return models.StikerResponse{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return models.StikerResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return models.StikerResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return models.StikerResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var stikerResponse models.StikerResponse
	if err := json.Unmarshal(bodyBytes, &stikerResponse); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return models.StikerResponse{}, err

	}

	return stikerResponse, nil
}

func (t *TelegramClient) GenerateInviteLinks(invite models.CreateChatInviteLinkRequest) ([]string, error) {
	requestBody, err := json.Marshal(invite)
	if err != nil {
		return nil, err
	}

	apiURL := "https://api.telegram.org/bot" + t.BotToken + "/createChatInviteLink"

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var inviteLinks []string
	if err := json.Unmarshal(bodyBytes, &inviteLinks); err != nil {
		return nil, err
	}

	return inviteLinks, nil
}

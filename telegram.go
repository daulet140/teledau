package teledau

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Telegram interface {
	GetChat(chatID string) (GetChatResponse, error)
	SendPoll(poolRequest PollRequest) (PollResponse, error)
	SendMedia(chatId string, media, message string, parseMode string) (*SendMessageResponse, error)
	SendSticker(chatId string, media string) (StikerResponse, error)
	SendMessage(message MessageRequest) (SendMessageResponse, error)
	EditMessage(message EditMessageRequest) (SendMessageResponse, error)
	EditCaption(message EditMessageRequest) (SendMessageResponse, error)
	GetFilePath(fileID string) (string, error)
	DownloadByte(filePath string) ([]byte, error)
	DownloadFile(fileName, filePath string) error
	DownloadStrBase64(filePath string) (string, error)
	GenerateInviteLinks(invite CreateChatInviteLinkRequest) (*InviteLinks, error)
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

func (t *TelegramClient) GetChat(chatID string) (GetChatResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/getChat?chat_id=" + chatID
	req, err := http.NewRequestWithContext(t.Ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return GetChatResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return GetChatResponse{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return GetChatResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return GetChatResponse{}, err
	}

	var chat GetChatResponse
	err = json.Unmarshal(bodyBytes, &chat)
	if err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return GetChatResponse{}, err
	}

	return chat, nil
}

func (t *TelegramClient) SendMessage(message MessageRequest) (SendMessageResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendMessage"

	messageData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return SendMessageResponse{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(messageData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return SendMessageResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return SendMessageResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return SendMessageResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return SendMessageResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var createdApplicant SendMessageResponse
	if err := json.Unmarshal(bodyBytes, &createdApplicant); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return SendMessageResponse{}, err
	}

	return createdApplicant, nil
}
func (t *TelegramClient) EditMessage(message EditMessageRequest) (SendMessageResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/editMessageText"

	messageData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return SendMessageResponse{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(messageData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return SendMessageResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return SendMessageResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return SendMessageResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return SendMessageResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var createdApplicant SendMessageResponse
	if err := json.Unmarshal(bodyBytes, &createdApplicant); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return SendMessageResponse{}, err
	}

	return createdApplicant, nil
}
func (t *TelegramClient) EditCaption(message EditMessageRequest) (SendMessageResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/editMessageCaption"

	messageData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return SendMessageResponse{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(messageData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return SendMessageResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return SendMessageResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return SendMessageResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return SendMessageResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var createdApplicant SendMessageResponse
	if err := json.Unmarshal(bodyBytes, &createdApplicant); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return SendMessageResponse{}, err
	}

	return createdApplicant, nil
}

func (t *TelegramClient) SendPoll(poolRequest PollRequest) (PollResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendPoll"

	pollData, err := json.Marshal(poolRequest)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return PollResponse{}, err
	}

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, bytes.NewBuffer(pollData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return PollResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return PollResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return PollResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return PollResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var pollResponse PollResponse
	if err := json.Unmarshal(bodyBytes, &pollResponse); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)

		return PollResponse{}, err
	}

	return pollResponse, nil
}

func (t *TelegramClient) SendMedia(chatId string, media, message, parseMode string) (*SendMessageResponse, error) {
	imgData, err := base64.StdEncoding.DecodeString(media)
	response := new(SendMessageResponse)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return response, err
	}
	tempFile, err := ioutil.TempFile("", "image*.jpeg")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return response, err
	}
	defer os.Remove(tempFile.Name()) // Clean up temporary file

	if _, err := tempFile.Write(imgData); err != nil {
		fmt.Println("Error writing image data to file:", err)
		return response, err
	}

	if err := tempFile.Close(); err != nil {
		fmt.Println("Error closing temporary file:", err)
		return response, err
	}

	// Open the temporary file
	file, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println("Error opening temporary file:", err)
		return response, err
	}
	defer file.Close()

	// Create a new HTTP request with the file
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%s", t.BotToken, chatId), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return response, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("photo", "image.jpeg")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return response, err
	}

	if _, err := io.Copy(part, file); err != nil {
		fmt.Println("Error copying file data:", err)
		return response, err
	}

	writer.WriteField("caption", message)
	if len(parseMode) <= 0 {
		writer.WriteField("parse_mode", "MarkdownV2")
	} else {
		writer.WriteField("parse_mode", parseMode)
	}

	// Close the multipart writer
	writer.Close()

	// Set the request body
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = ioutil.NopCloser(body)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return response, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)

		return response, err
	}

	json.Unmarshal(bodyBytes, &response)
	return response, nil
}

func (t *TelegramClient) SendSticker(chatId string, media string) (StikerResponse, error) {
	url := "https://api.telegram.org/bot" + t.BotToken + "/sendSticker"
	filePath := "/path/to/decoded/sticker.webp"

	imageBytes, err := base64.StdEncoding.DecodeString(media)
	if err != nil {
		fmt.Printf("err: %v", err)
		return StikerResponse{}, err
	}

	var buffer bytes.Buffer
	_, err = buffer.Write(imageBytes)
	if err != nil {
		return StikerResponse{}, err
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part, err := writer.CreateFormFile("sticker", filepath.Base(filePath))
	if err != nil {
		log.Printf("create form file err: %v", err)
		return StikerResponse{}, err
	}

	_, err = io.Copy(part, &buffer)
	if err != nil {
		log.Printf("io copy err: %v", err)
		return StikerResponse{}, err
	}

	_ = writer.WriteField("chat_id", chatId)
	writer.Close()

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return StikerResponse{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return StikerResponse{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return StikerResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))
		return StikerResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var stikerResponse StikerResponse
	if err := json.Unmarshal(bodyBytes, &stikerResponse); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return StikerResponse{}, err

	}

	return stikerResponse, nil
}

func (t *TelegramClient) GenerateInviteLinks(invite CreateChatInviteLinkRequest) (*InviteLinks, error) {
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

	inviteLinks := new(InviteLinks)
	if err := json.Unmarshal(bodyBytes, &inviteLinks); err != nil {
		return nil, err
	}

	return inviteLinks, nil
}

func (t *TelegramClient) DeleteMessage(messageId int, chatId int64) error {

	url := "https://api.telegram.org/bot" + t.BotToken + "/deleteMessage"
	requestBody, err := json.Marshal(map[string]string{
		"chat_id":    fmt.Sprintf("%d", chatId),
		"message_id": fmt.Sprintf("%d", messageId),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (t *TelegramClient) ForwardMessage(chatId string, fromChatId string, messageId string) ([]byte, error) {

	url := "https://api.telegram.org/bot" + t.BotToken + "/forwardMessage"

	requestBody, err := json.Marshal(map[string]string{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (t *TelegramClient) DownloadStrBase64(filePath string) (string, error) {
	// Encode file bytes as base64
	fileBytes, err := t.DownloadByte(filePath)
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	return base64String, nil
}
func (t *TelegramClient) DownloadByte(filePath string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", t.BotToken, filePath))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (t *TelegramClient) DownloadFile(fileName, filePath string) error {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", t.BotToken, fileName))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (t *TelegramClient) GetFilePath(fileID string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", t.BotToken, fileID))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var fileResponse struct {
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}
	err = json.Unmarshal(body, &fileResponse)
	if err != nil {
		return "", err
	}

	return fileResponse.Result.FilePath, nil
}

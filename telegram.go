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
	SendMessage(message MessageRequest) (Message, error)
	//SendMessageWithImg(chatId string, message models.MessageRequest, img []string)

	//SendPost(chatId string, title string, text string, img []string)

	SendPoll(poolRequest PollRequest) (PollResponse, error)

	SendMedia(chatId string, media string)

	SendSticker(chatId int64, media string) (StikerResponse, error)

	GenerateInviteLinks(invite CreateChatInviteLinkRequest) ([]string, error)
	DownloadAndEncodeFile(filePath string) (string, error)
	DownloadFile(fileName, filePath string) error
	FetFilePath(fileID string) (string, error)
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

func (t *TelegramClient) SendMedia(chatId string, media, message string) error {
	imgData, err := base64.StdEncoding.DecodeString(media)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return err
	}
	tempFile, err := ioutil.TempFile("", "image*.jpeg")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return err
	}
	defer os.Remove(tempFile.Name()) // Clean up temporary file

	if _, err := tempFile.Write(imgData); err != nil {
		fmt.Println("Error writing image data to file:", err)
		return err
	}

	if err := tempFile.Close(); err != nil {
		fmt.Println("Error closing temporary file:", err)
		return err
	}

	// Open the temporary file
	file, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println("Error opening temporary file:", err)
		return err
	}
	defer file.Close()

	// Create a new HTTP request with the file
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%s", t.BotToken, chatId), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return err
	}

	// Create a new form file field
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", "image.jpeg")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return err
	}

	// Copy the file data to the form file field
	if _, err := io.Copy(part, file); err != nil {
		fmt.Println("Error copying file data:", err)
		return err
	}

	// Add the caption field
	writer.WriteField("caption", "text")

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
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}
	log.Printf("%s", bodyBytes)
	fmt.Println("Photo sent successfully.")
	return nil
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

func (t *TelegramClient) DownloadAndEncodeFile(filePath string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", t.BotToken, filePath))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Encode file bytes as base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	return base64String, nil
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

func (t *TelegramClient) FetFilePath(fileID string) (string, error) {
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

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
	GetChat(chatID string) (*GetChatResponse, error)

	SendMessage(message MessageRequest) (SendMessageResponse, error)
	EditMessage(message EditMessageRequest) (SendMessageResponse, error)

	SendMedia(chatId, media, message string, parseMode string) (*SendMessageResponse, error)
	SendMediaGroups(chatId string, media []string, message, parseMode string) (*SendMessageResponse, error)
	EditCaption(message EditCaptionRequest) (SendMessageResponse, error)

	SendSticker(chatId string, media string) (StikerResponse, error)

	GetFilePath(fileID string) (string, error)
	DownloadByte(filePath string) ([]byte, error)
	DownloadFile(fileName, filePath string) error
	DownloadStrBase64(filePath string) (string, error)

	GenerateInviteLinks(invite CreateChatInviteLinkRequest) (*InviteLinks, error)

	SendPoll(poolRequest PollRequest) (PollResponse, error)
}

type TelegramClient struct {
	Ctx        context.Context
	BotToken   string
	HttpClient http.Client
}

func NewTelegramClient(ctx context.Context, botToken string) *TelegramClient {
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
func NewTelegramClientWithClient(ctx context.Context, botToken string, httpClient http.Client) *TelegramClient {

	return &TelegramClient{
		BotToken:   botToken,
		HttpClient: httpClient,
		Ctx:        ctx,
	}
}

// GetChat retrieves chat information from the Telegram API for a given chat ID.
// It constructs an HTTP GET request using the bot token and chat ID, sends the request,
// and processes the response. If successful, it returns a GetChatResponse containing
// the chat details. In case of errors during request creation, sending, or response
// processing, it logs the error and returns an empty GetChatResponse along with the error.
func (t *TelegramClient) GetChat(chatID string) (*GetChatResponse, error) {

	url := TgBotBaseUrl + t.BotToken + TgBotGetChat + chatID

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)

		return nil, err
	}

	req.Header.Set(HeaderContentType, ApplicationJson)
	resp, err := t.HttpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)

		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)

		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API request failed with status code: %d body %v", resp.StatusCode, string(bodyBytes))

		return nil, err
	}

	var chat GetChatResponse
	err = json.Unmarshal(bodyBytes, &chat)
	if err != nil {
		log.Printf("Error unmarshalling response: %v", err)

		return nil, err
	}

	return &chat, nil
}
func (t *TelegramClient) SendMessage(message MessageRequest) (SendMessageResponse, error) {
	url := TgBotBaseUrl + t.BotToken + TgBotSendMessageUrl

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

	req.Header.Set(HeaderContentType, ApplicationJson)
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
	url := TgBotBaseUrl + t.BotToken + TgBotEditMessageUrl

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

	req.Header.Set(HeaderContentType, ApplicationJson)
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
func (t *TelegramClient) EditCaption(message EditCaptionRequest) (SendMessageResponse, error) {
	url := TgBotBaseUrl + t.BotToken + TgBotEditCaptionUrl

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

	req.Header.Set(HeaderContentType, ApplicationJson)
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
	url := TgBotBaseUrl + t.BotToken + TgBotSendPoolUrl

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

	req.Header.Set(HeaderContentType, ApplicationJson)
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
	tempFile, err := ioutil.TempFile("", TempFileName)
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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(TgFieldMediaType, TempFileName)
	if err != nil {
		fmt.Println("Error creating form file:", err)

		return response, err
	}

	if _, err := io.Copy(part, file); err != nil {
		fmt.Println("Error copying file data:", err)

		return response, err
	}

	writer.WriteField(TgFieldCaption, message)
	if len(parseMode) <= 0 {
		writer.WriteField(TgFieldNameParseMod, TgParseModMarkdownV2)
	} else {
		writer.WriteField(TgFieldNameParseMod, parseMode)
	}

	// Close the multipart writer
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(TgBotSendPhotoUrlSptf, t.BotToken, chatId), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)

		return response, err
	}

	// Set the request body
	req.Header.Set(HeaderContentType, writer.FormDataContentType())
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

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Printf("Error unmarshal response body: %v", err)

		return response, err
	}

	return response, nil
}

func (t *TelegramClient) SendMediaGroup(chatId string, media []string, message, parseMode string) (*SendMessageResponse, error) {
	prefix := time.Now().UnixMilli()
	response := new(SendMessageResponse)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField(TgFieldChatId, chatId)

	var mediaGroups []MediaGroup

	for i, s := range media {
		imgData, err := base64.StdEncoding.DecodeString(s)

		if err != nil {
			fmt.Println("Error decoding base64 string:", err)
			return response, err
		}

		tempFile, err := ioutil.TempFile("", fmt.Sprintf(TempFileNameFmt, prefix, i))
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

		file, err := os.Open(tempFile.Name())
		if err != nil {
			fmt.Println("Error opening temporary file:", err)
			return response, err
		}

		defer file.Close()

		part, err := writer.CreateFormFile(fmt.Sprintf("photo%d", i), fmt.Sprintf(TempFileNameFmt, prefix, i))
		if err != nil {
			fmt.Println("Error creating form file:", err)
			return response, err
		}

		if _, err := io.Copy(part, file); err != nil {
			fmt.Println("Error copying file data:", err)
			return response, err
		}

		mediaG := MediaGroup{Type: TgFieldMediaType, Media: fmt.Sprintf("attach://photo%d", i)}
		if i == 0 {
			mediaG.Caption = message

			if len(parseMode) <= 0 {
				err := writer.WriteField(TgFieldNameParseMod, TgParseModMarkdownV2)
				if err != nil {
					return nil, err
				}
				mediaG.ParseMode = TgParseModMarkdownV2

			} else {

				err := writer.WriteField(TgFieldNameParseMod, parseMode)
				if err != nil {
					return nil, err
				}
				mediaG.ParseMode = parseMode
			}
		}

		mediaGroups = append(mediaGroups, mediaG)

	}
	mediaGroupBytes, err := json.Marshal(mediaGroups)
	if err != nil {
		fmt.Println("Error marshalling mediaGroups:", err)
		return response, err
	}

	err = writer.WriteField("media", string(mediaGroupBytes))
	if err != nil {
		return nil, err
	}

	err = writer.WriteField(TgFieldCaption, message)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(TgBotSendMediaGroupUrlSptf, t.BotToken, chatId), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return response, err
	}

	req.Header.Set(HeaderContentType, writer.FormDataContentType())
	req.Body = ioutil.NopCloser(body)

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

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Printf("Error unmarshal response body: %v", err)

		return response, err
	}

	return response, nil
}
func (t *TelegramClient) SendSticker(chatId string, media string) (StikerResponse, error) {
	url := TgBotBaseUrl + t.BotToken + TgBotSendStickerUrl
	filePath := TempStickerFileName

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

	part, err := writer.CreateFormFile(TgFieldSticker, filepath.Base(filePath))
	if err != nil {
		log.Printf("create form file err: %v", err)
		return StikerResponse{}, err
	}

	_, err = io.Copy(part, &buffer)
	if err != nil {
		log.Printf("io copy err: %v", err)
		return StikerResponse{}, err
	}

	_ = writer.WriteField(TgFieldChatId, chatId)
	writer.Close()

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, url, payload)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return StikerResponse{}, err
	}

	req.Header.Set(HeaderContentType, writer.FormDataContentType())

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

	apiURL := TgBotBaseUrl + t.BotToken + TgBotCreateInviteLinkUrl

	req, err := http.NewRequestWithContext(t.Ctx, http.MethodPost, apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set(HeaderContentType, ApplicationJson)
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

	url := TgBotBaseUrl + t.BotToken + TgBotDeleteMsgUrl

	requestBody, err := json.Marshal(map[string]string{
		TgFieldChatId:    fmt.Sprintf("%d", chatId),
		TgFieldMessageId: fmt.Sprintf("%d", messageId),
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, ApplicationJson, bytes.NewBuffer(requestBody))
	if err != nil {

		return err
	}

	defer resp.Body.Close()

	return nil
}

func (t *TelegramClient) ForwardMessage(chatId string, fromChatId string, messageId string) ([]byte, error) {
	url := TgBotBaseUrl + t.BotToken + TgBotForwardMsgUrl

	requestBody, err := json.Marshal(map[string]string{
		TgFieldChatId:     chatId,
		TgFieldMessageId:  messageId,
		TgFieldFromChatId: fromChatId,
	})
	if err != nil {

		return nil, err
	}

	resp, err := http.Post(url, ApplicationJson, bytes.NewBuffer(requestBody))
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
	fileBytes, err := t.DownloadByte(filePath)
	if err != nil {

		return "", err
	}

	return base64.StdEncoding.EncodeToString(fileBytes), nil
}

func (t *TelegramClient) DownloadByte(filePath string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(TgBotDownloadFileUrl, t.BotToken, filePath))
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
	resp, err := http.Get(fmt.Sprintf(TgBotDownloadFileUrl, t.BotToken, fileName))
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
	var fileResponse FileResponse
	resp, err := http.Get(fmt.Sprintf(TgBotGetFileUrl, t.BotToken, fileID))
	if err != nil {

		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}

	err = json.Unmarshal(body, &fileResponse)
	if err != nil {
		return "", err
	}

	return fileResponse.Result.FilePath, nil
}

package teledau

const (
	ApplicationJson = "application/json"

	HeaderContentType = "Content-Type"

	TgBotBaseUrl               = "https://api.telegram.org/bot"
	TgBotGetChat               = "/getChat?chat_id="
	TgBotSendMessageUrl        = "/sendMessage"
	TgBotEditMessageUrl        = "/editMessageText"
	TgBotEditCaptionUrl        = "/editMessageCaption"
	TgBotSendPoolUrl           = "/sendPoll"
	TgBotForwardMsgUrl         = "/forwardMessage"
	TgBotDeleteMsgUrl          = "/deleteMessage"
	TgBotSendStickerUrl        = "/sendSticker"
	TgBotCreateInviteLinkUrl   = "/createChatInviteLink"
	TgBotSendPhotoUrlSptf      = "https://api.telegram.org/bot%s/sendPhoto?chat_id=%s"
	TgBotSendMediaGroupUrlSptf = "https://api.telegram.org/bot%s/sendMediaGroup?chat_id=%s"
	TgBotDownloadFileUrl       = "https://api.telegram.org/file/bot%s/%s"
	TgBotGetFileUrl            = "https://api.telegram.org/bot%s/getFile?file_id=%s"

	TgParseModMarkdownHTML = "HTML"
	TgParseModMarkdownV1   = "Markdown"
	TgParseModMarkdownV2   = "MarkdownV2"

	TgFieldNameParseMod = "parse_mode"
	TgFieldCaption      = "caption"
	TgFieldChatId       = "chat_id"
	TgFieldSticker      = "sticker"
	TgFieldMediaType    = "photo"
	TgFieldMessageId    = "message_id"
	TgFieldFromChatId   = "from_chat_id"

	TempFileName        = "image*.jpeg"
	TempStickerFileName = "/path/to/decoded/sticker.webp"
	TempFileNameFmt     = "image_%d_%d*.jpeg"
)

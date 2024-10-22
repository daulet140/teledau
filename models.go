package teledau

type Channel struct {
	Id       int
	ChatID   string
	Username string
}

type PollRequest struct {
	ChatId          string   `json:"chat_id"`
	Question        string   `json:"question"`
	Options         []string `json:"options"`
	IsAnonymous     bool     `json:"is_anonymous"`
	Type            string   `json:"type"`
	CorrectOptionId int      `json:"correct_option_id"`
	Explanation     string   `json:"explanation"`
	OpenPeriod      int      `json:"open_period"`
	CloseDate       int      `json:"close_date"`
}

type StikerResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId  int64 `json:"message_id"`
		SenderChat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"sender_chat"`
		Chat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"chat"`
		Date    int `json:"date"`
		Sticker struct {
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			IsAnimated bool   `json:"is_animated"`
			IsVideo    bool   `json:"is_video"`
			Type       string `json:"type"`
			Thumbnail  struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
			} `json:"thumbnail"`
			Thumb struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
			} `json:"thumb"`
			FileId       string `json:"file_id"`
			FileUniqueId string `json:"file_unique_id"`
			FileSize     int    `json:"file_size"`
		} `json:"sticker"`
	} `json:"result"`
}

type PostResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		Chat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"chat"`
		Date     int64 `json:"date"`
		Entities []struct {
			Length int    `json:"length"`
			Offset int    `json:"offset"`
			Type   string `json:"type"`
		} `json:"entities"`
		MessageId  int64 `json:"message_id"`
		SenderChat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"sender_chat"`
		Text string `json:"text"`
	} `json:"result"`
}

type MediaPostResponse struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Chat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"chat"`
		Date     int64 `json:"date"`
		Entities []struct {
			Length int    `json:"length"`
			Offset int    `json:"offset"`
			Type   string `json:"type"`
		} `json:"entities"`
		MessageId  int64 `json:"message_id"`
		SenderChat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"sender_chat"`
		Text string `json:"text"`
	} `json:"result"`
}

type PollResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId  int `json:"message_id"`
		SenderChat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"sender_chat"`
		Chat struct {
			Id       int64  `json:"id"`
			Title    string `json:"title"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"chat"`
		Date int `json:"date"`
		Poll struct {
			Id       string `json:"id"`
			Question string `json:"question"`
			Options  []struct {
				Text       string `json:"text"`
				VoterCount int    `json:"voter_count"`
			} `json:"options"`
			TotalVoterCount       int           `json:"total_voter_count"`
			OpenPeriod            int           `json:"open_period"`
			CloseDate             int           `json:"close_date"`
			IsClosed              bool          `json:"is_closed"`
			IsAnonymous           bool          `json:"is_anonymous"`
			Type                  string        `json:"type"`
			AllowsMultipleAnswers bool          `json:"allows_multiple_answers"`
			CorrectOptionId       int           `json:"correct_option_id"`
			Explanation           string        `json:"explanation"`
			ExplanationEntities   []interface{} `json:"explanation_entities"`
		} `json:"poll"`
	} `json:"result"`
}

type MessageId struct {
	ChatId    string `json:"chat_id"`
	MessageId int    `json:"message_id"`
}

type Post struct {
	ChatId     string   `json:"chat_id"`
	Caption    string   `json:"caption"`
	Photo      []string `json:"photo"` //base64
	IsMarkdown bool     `json:"is_markdown"`
}

type MediaGroup struct {
	Type            string            `json:"type"`
	Media           string            `json:"media"`
	Caption         string            `json:"caption,omitempty"`
	CaptionEntities []CaptionEntities `json:"caption_entities,omitempty"`
	ParseMode       string            `json:"parse_mode,omitempty"`
}

type CaptionEntities struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type PostRequest struct {
	Id               int64    `json:"id"`
	Title            string   `json:"title"`
	ChannelId        int      `json:"channel_id"`
	ChatUsername     string   `json:"chat_id"`
	Text             string   `json:"text"`
	Img              []string `json:"img"`
	ParseMode        string   `json:"parse_mode"`
	PostType         int64    `json:"post_type"`
	MessageId        int64    `json:"message_id"`
	Status           int64    `json:"status"`
	PostedAt         string   `json:"posted_at"`
	ReplyToMessageId int64    `json:"reply_to_message_id"`
}

type Update struct {
	UpdateId   int         `json:"update_id"`
	ChatMember *ChatMember `json:"chat_member,omitempty"`
	Message    *Message    `json:"message,omitempty"`
}

type SendMessageResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat       Chat    `json:"chat"`
		SenderChat Chat    `json:"sender_chat"`
		Date       int     `json:"date"`
		Text       string  `json:"text"`
		Photo      []Photo `json:"photo"`
	} `json:"result"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type Message struct {
	MessageId int `json:"message_id"`
	From      struct {
		Id           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		Id        int64  `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date           int             `json:"date"`
	Text           string          `json:"text"`
	ReplyToMessage *ReplyToMessage `json:"reply_to_message,omitempty"`
	Photo          []Photo         `json:"photo"`
}

type Photo struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type ReplyToMessage struct {
	MessageId int `json:"message_id"`
	From      struct {
		Id           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		Id        int    `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

type ChatMember struct {
	Chat struct {
		Id       int64  `json:"id"`
		Title    string `json:"title"`
		Username string `json:"username"`
		Type     string `json:"type"`
	} `json:"chat"`
	From struct {
		Id        int64  `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name,"`
		LastName  string `json:"last_name,omitempty"`
		Username  string `json:"username"`
	} `json:"from"`
	Date          int `json:"date"`
	OldChatMember struct {
		User struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name,omitempty"`
			Username  string `json:"username"`
		} `json:"user"`
		Status string `json:"status"`
	} `json:"old_chat_member"`
	NewChatMember struct {
		User struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name,omitempty"`
			Username  string `json:"username"`
		} `json:"user"`
		Status string `json:"status"`
	} `json:"new_chat_member"`
	InviteLink struct {
		InviteLink string `json:"invite_link"`
		Name       string `json:"name"`
		Creator    struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"creator"`
		CreatesJoinRequest bool `json:"creates_join_request"`
		IsPrimary          bool `json:"is_primary"`
		IsRevoked          bool `json:"is_revoked"`
	} `json:"invite_link,omitempty"`
}

type InviteLinks struct {
	Ok     bool `json:"ok"`
	Result struct {
		InviteLink string `json:"invite_link"`
		Name       string `json:"name"`
		Creator    struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"creator"`
		CreatesJoinRequest bool `json:"creates_join_request"`
		IsPrimary          bool `json:"is_primary"`
		IsRevoked          bool `json:"is_revoked"`
	} `json:"result"`
}

type MediaResponse struct {
	Ok bool `json:"ok"`
}

type GetChatResponse struct {
	Ok      bool    `json:"ok"`
	GetChat GetChat `json:"result"`
}
type GetChat struct {
	Id                int64      `json:"id"`
	Title             string     `json:"title"`
	Username          string     `json:"username"`
	Type              string     `json:"type"`
	ActiveUsernames   []string   `json:"active_usernames"`
	Description       string     `json:"description"`
	InviteLink        string     `json:"invite_link"`
	HasVisibleHistory bool       `json:"has_visible_history"`
	Photo             *ChatPhoto `json:"photo"`
	MaxReactionCount  int        `json:"max_reaction_count"`
	AccentColorId     int        `json:"accent_color_id"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

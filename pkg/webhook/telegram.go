package webhook

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func FormatTelegramMessage(chatID, text string) TelegramMessage {
	return TelegramMessage{ChatID: chatID, Text: text}
}

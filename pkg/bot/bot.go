package bot

type BotMessage struct {
	ChatID string
	URL    string
}

type BotDaemon struct {
	TelegramToken string
	DiscordToken  string
}

func NewBotDaemon(tgToken, dcToken string) *BotDaemon {
	return &BotDaemon{TelegramToken: tgToken, DiscordToken: dcToken}
}

func (b *BotDaemon) HandleMessage(msg BotMessage) string {
	if msg.URL == "" {
		return "Please send a valid media URL."
	}
	return "Queued download for: " + msg.URL
}

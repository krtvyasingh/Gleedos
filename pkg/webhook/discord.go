package webhook

type DiscordEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
}

type DiscordPayload struct {
	Content string         `json:"content"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

func FormatDiscordMessage(title, file string) DiscordPayload {
	return DiscordPayload{
		Content: "Download completed: " + title,
		Embeds: []DiscordEmbed{
			{Title: title, Description: file, Color: 0x00FF00},
		},
	}
}

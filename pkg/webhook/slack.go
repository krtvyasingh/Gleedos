package webhook

type SlackMessage struct {
	Text string `json:"text"`
}

func FormatSlackMessage(title, file string) SlackMessage {
	return SlackMessage{Text: "*Download Finished*: " + title + " (" + file + ")"}
}

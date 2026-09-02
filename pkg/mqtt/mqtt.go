package mqtt

import "fmt"

type DownloadEvent struct {
	URL        string
	Title      string
	OutputPath string
}

func FormatMQTTTopic(topicPrefix, eventType string) string {
	return fmt.Sprintf("%s/%s", topicPrefix, eventType)
}

package dlna

import "fmt"

func FormatDIDLItem(id, parentID, title, uri, mimeType string) string {
	return fmt.Sprintf(`<item id="%s" parentID="%s" restricted="1"><dc:title>%s</dc:title><res protocolInfo="http-get:*:%s:*">%s</res></item>`, id, parentID, title, mimeType, uri)
}

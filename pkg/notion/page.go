package notion

import "fmt"

func FormatNotionPageJSON(title, url string) string {
	return fmt.Sprintf(`{"properties":{"Name":{"title":[{"text":{"content":"%s"}}]},"URL":{"url":"%s"}}}`, title, url)
}

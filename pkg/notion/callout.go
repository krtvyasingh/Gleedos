package notion

import "fmt"

func FormatCalloutBlock(emoji, text string) string {
	return fmt.Sprintf("> [!NOTE]\n> %s %s\n", emoji, text)
}

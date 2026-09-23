package arweave

import "fmt"

func FormatArweaveManifest(txID, mimeType string) string {
	return fmt.Sprintf(`{"arweave_tx":"%s","content_type":"%s"}`, txID, mimeType)
}

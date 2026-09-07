package dlna

import "fmt"

func BuildSSDPOK(location, usn string) string {
	return fmt.Sprintf("HTTP/1.1 200 OK\r\nCACHE-CONTROL: max-age=1800\r\nLOCATION: %s\r\nUSN: %s\r\n\r\n", location, usn)
}

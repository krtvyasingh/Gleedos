package safety

import "regexp"

var secretPattern = regexp.MustCompile(`(?i)(api[_-]?key|secret|password)\s*[:=]\s*['"][^'"]+['"]`)

func DetectSecrets(text string) bool {
	return secretPattern.MatchString(text)
}

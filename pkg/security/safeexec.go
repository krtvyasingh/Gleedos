package security

import (
	"errors"
	"strings"
)

func ValidateArgument(arg string) error {
	if strings.Contains(arg, ";") || strings.Contains(arg, "&&") || strings.Contains(arg, "|") {
		return errors.New("command chaining character detected")
	}
	return nil
}

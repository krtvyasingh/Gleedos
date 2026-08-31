package ui

import "fmt"

func FormatRow(col1, col2 string) string {
	return fmt.Sprintf("%-20s : %s", col1, col2)
}

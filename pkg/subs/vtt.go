package subs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func VTTtoSRT(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	idx := 1
	var buffer []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "WEBVTT") || strings.HasPrefix(line, "NOTE") {
			continue
		}
		if line == "" {
			if len(buffer) > 0 {
				fmt.Fprintf(w, "%d\n", idx)
				for _, l := range buffer {
					fmt.Fprintln(w, l)
				}
				fmt.Fprintln(w)
				idx++
				buffer = nil
			}
			continue
		}

		if strings.Contains(line, "-->") {
			line = strings.ReplaceAll(line, ".", ",")
		}
		buffer = append(buffer, line)
	}

	if len(buffer) > 0 {
		fmt.Fprintf(w, "%d\n", idx)
		for _, l := range buffer {
			fmt.Fprintln(w, l)
		}
		fmt.Fprintln(w)
	}

	return scanner.Err()
}

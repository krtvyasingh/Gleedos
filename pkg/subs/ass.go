package subs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const ASSHeader = `[Script Info]
Title: Gleedos Subtitles
ScriptType: v4.00+
Collisions: Normal
PlayDepth: 0

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Default,Arial,20,&H00FFFFFF,&H000000FF,&H00000000,&H00000000,0,0,0,0,100,100,0,0,1,1,0,2,10,10,10,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`

func SRTtoASS(r io.Reader, w io.Writer) error {
	fmt.Fprint(w, ASSHeader)
	scanner := bufio.NewScanner(r)

	var timing string
	var textLines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if timing != "" && len(textLines) > 0 {
				times := strings.Split(timing, "-->")
				if len(times) == 2 {
					start := formatASSTime(strings.TrimSpace(times[0]))
					end := formatASSTime(strings.TrimSpace(times[1]))
					text := strings.Join(textLines, "\\N")
					fmt.Fprintf(w, "Dialogue: 0,%s,%s,Default,,0,0,0,,%s\n", start, end, text)
				}
				timing = ""
				textLines = nil
			}
			continue
		}

		if strings.Contains(line, "-->") {
			timing = line
		} else if timing != "" {
			textLines = append(textLines, line)
		}
	}

	if timing != "" && len(textLines) > 0 {
		times := strings.Split(timing, "-->")
		if len(times) == 2 {
			start := formatASSTime(strings.TrimSpace(times[0]))
			end := formatASSTime(strings.TrimSpace(times[1]))
			text := strings.Join(textLines, "\\N")
			fmt.Fprintf(w, "Dialogue: 0,%s,%s,Default,,0,0,0,,%s\n", start, end, text)
		}
	}

	return scanner.Err()
}

func formatASSTime(srtTime string) string {
	srtTime = strings.ReplaceAll(srtTime, ",", ".")
	parts := strings.Split(srtTime, ":")
	if len(parts) == 3 {
		secParts := strings.Split(parts[2], ".")
		if len(secParts) == 2 && len(secParts[1]) > 2 {
			parts[2] = secParts[0] + "." + secParts[1][:2]
		}
		return fmt.Sprintf("%s:%s:%s", parts[0], parts[1], parts[2])
	}
	return srtTime
}

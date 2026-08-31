package ui

import "fmt"

type BarSlot struct {
	ID      int
	Percent float64
	Label   string
}

func RenderSlot(s BarSlot) string {
	return fmt.Sprintf("[%d] %-15s %5.1f%%", s.ID, s.Label, s.Percent)
}

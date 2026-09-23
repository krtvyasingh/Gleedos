package matrixgrid

import "fmt"

func RenderSlotStatus(slotID int, title string, progressPercent float64) string {
	return fmt.Sprintf("[Slot %d] %s: %.1f%%", slotID, title, progressPercent)
}

package main

import (
	"fmt"

	"github.com/krtvysingh/gleedos/pkg/palette"
	"github.com/krtvysingh/gleedos/pkg/spsc"
	"github.com/krtvysingh/gleedos/pkg/themes"
)

func handleThemeDemo(name string) {
	th := themes.GetTheme(name)
	fmt.Printf("Theme: %s | Background: %s | Accent: %s\n", th.Name, th.Background, th.Accent)
}

func handlePaletteDemo() {
	fmt.Printf("Extracted Dominant Color: %s\n", palette.FormatRGBHex(137, 180, 250))
}

func handleSPSCDemo() {
	rb := spsc.NewRingBuffer(4)
	rb.Push(0xAA)
	val, _ := rb.Pop()
	fmt.Printf("Lock-Free SPSC Ring Buffer Initialized: Read byte 0x%X\n", val)
}

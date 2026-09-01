package main

import (
	"fmt"

	"github.com/krtvysingh/gleedos/pkg/metrics"
	"github.com/krtvysingh/gleedos/pkg/tui"
	"github.com/krtvysingh/gleedos/pkg/vault"
)

func handleVaultKeyDerive(pass string) {
	key := vault.DeriveVaultKey(pass)
	fmt.Printf("Vault 256-bit Key Derived: %x\n", key)
}

func handleTUIDemo() {
	wf := tui.RenderWaveform([]float64{0.1, 0.3, 0.7, 0.9, 0.4, 0.2})
	fmt.Printf("Gleedos Audio Spectrum Waveform: %s\n", wf)
}

func handleExportGrafana() {
	fmt.Println(metrics.GetDashboardTemplate())
}

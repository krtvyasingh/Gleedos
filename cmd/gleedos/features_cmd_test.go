package main

import "testing"

func TestFeaturesCommandHandlers(t *testing.T) {
	handleVaultKeyDerive("secret")
	handleTUIDemo()
	handleExportGrafana()
}

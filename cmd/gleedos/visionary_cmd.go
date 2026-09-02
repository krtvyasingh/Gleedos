package main

import (
	"fmt"

	"github.com/krtvysingh/gleedos/pkg/batch"
	"github.com/krtvysingh/gleedos/pkg/pqc"
	"github.com/krtvysingh/gleedos/pkg/wireframe3d"
)

func handlePQCDemo() {
	kp, err := pqc.GenerateKyberKeypair()
	if err != nil {
		fmt.Printf("PQC Kyber Error: %v\n", err)
		return
	}
	fmt.Printf("Kyber-1024 Post-Quantum Keypair Initialized: %d bytes pub key\n", len(kp.PublicKey))
}

func handleWireframeDemo() {
	fmt.Println(wireframe3d.Render3DCube(0))
}

func handleUniversalExportDemo(outPath, format string) {
	records := []batch.MediaExportRecord{
		{URL: "https://example.com/demo.mp4", Title: "Gleedos Visionary Release", Format: "mp4", Quality: "4K"},
	}
	_ = batch.ExportPlaylist(records, format, outPath)
}

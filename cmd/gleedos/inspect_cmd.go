package main

import (
	"fmt"

	"github.com/krtvysingh/gleedos/pkg/probe"
)

func handleInspect(path string) {
	info, err := probe.InspectFile(path)
	if err != nil {
		fmt.Printf("Inspect error: %v\n", err)
		return
	}
	fmt.Printf("File: %s\nSize: %d bytes\nExt: %s\n", info.Path, info.SizeBytes, info.Extension)
}

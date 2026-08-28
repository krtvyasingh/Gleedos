package main

import "github.com/krtvysingh/gleedos/pkg/doctor"

func handleDoctor() {
	results := doctor.RunDiagnostics()
	doctor.PrintReport(results)
}

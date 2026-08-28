package doctor

import (
	"fmt"
	"os/exec"
)

type CheckResult struct {
	Name    string
	Found   bool
	Version string
}

func CheckTool(name string) CheckResult {
	path, err := exec.LookPath(name)
	if err != nil {
		return CheckResult{Name: name, Found: false}
	}
	out, _ := exec.Command(path, "--version").Output()
	return CheckResult{
		Name:    name,
		Found:   true,
		Version: string(out),
	}
}

func RunDiagnostics() []CheckResult {
	tools := []string{"yt-dlp", "ffmpeg", "ffprobe"}
	results := make([]CheckResult, len(tools))
	for i, t := range tools {
		results[i] = CheckTool(t)
	}
	return results
}

func PrintReport(results []CheckResult) {
	fmt.Println("Gleedos System Diagnostics:")
	for _, r := range results {
		if r.Found {
			fmt.Printf("  ✓ %-10s : Installed\n", r.Name)
		} else {
			fmt.Printf("  ✗ %-10s : Not Found (Optional for native streams)\n", r.Name)
		}
	}
}

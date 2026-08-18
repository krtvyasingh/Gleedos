package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const banner = `
  ██████╗ ██╗     ███████╗███████╗███████╗██████╗  ██████╗ ███████╗
 ██╔════╝ ██║     ██╔════╝██╔════╝██╔══██╗██╔═══██╗██╔════╝
 ██║  ███╗██║     █████╗  █████╗  ██║  ██║██║   ██║███████╗
 ██║   ██║██║     ██╔══╝  ██╔══╝  ██║  ██║██║   ██║╚════██║
 ╚██████╔╝███████╗███████╗██████╔╝╚██████╔╝██████╔╝███████║
  ╚═════╝ ╚══════╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝ ╚══════╝

  Universal terminal downloader
`

func printUsage() {
	fmt.Print(`Gleedos — Universal terminal downloader

Usage:
  gleedos <URL>
  gleedos <URL> --best
  gleedos <URL> --audio
  gleedos <URL> --format mp4
  gleedos <URL> --turbo
  gleedos <URL> --list-formats
  gleedos <URL> -o <path>

Options:
  --best            Best available quality
  --audio           Extract audio as MP3
  --format <ext>    Prefer output format/container
  --turbo           Increase concurrent fragments
  --list-formats    Show all formats available from yt-dlp
  -o, --output      Output directory/path
  -h, --help        Show this help
`)
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "Gleedos: "+format+"\n", args...)
	os.Exit(1)
}

func validURL(url string) bool {
	return strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://")
}

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Print(banner)
		printUsage()
		os.Exit(1)
	}

	if cliArgs[0] == "--help" || cliArgs[0] == "-h" {
		fmt.Print(banner)
		printUsage()
		os.Exit(0)
	}

	url := cliArgs[0]

	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		fmt.Fprintln(os.Stderr, "Gleedos: invalid URL")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Gleedos:", err)
		os.Exit(1)
	}

	outputDir := filepath.Join(home, "Downloads", "Gleedos")
	quality := "bv*[height<=1080]+ba/b[height<=1080]/b[ext=mp4]/b"
	audioMode := false
	turbo := false

	for i := 1; i < len(cliArgs); i++ {
		switch cliArgs[i] {
		case "--best":
			quality = "bv*+ba/b[ext=mp4]/b"

		case "--audio":
			audioMode = true

		case "--turbo":
			turbo = true

		case "--format":
			if i+1 >= len(cliArgs) {
				fmt.Fprintln(os.Stderr, "Gleedos: --format requires a value")
				os.Exit(1)
			}
			i++
			format := cliArgs[i]
			quality = fmt.Sprintf(
				"bv*[ext=%s]+ba/b[ext=%s]/b[ext=%s]",
				format,
				format,
				format,
			)

		case "-o", "--output":
			if i+1 >= len(cliArgs) {
				fmt.Fprintln(os.Stderr, "Gleedos: output path required")
				os.Exit(1)
			}
			i++
			outputDir = cliArgs[i]

		case "--list-formats":
			cmd := exec.Command(
				"yt-dlp",
				"--no-playlist",
				"--list-formats",
				url,
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				os.Exit(1)
			}
			return
		}
	}

	if turbo {
		// Kept intentionally conservative. YouTube throttling/403s
		// are not fixed by blindly increasing fragment concurrency.
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Gleedos:", err)
		os.Exit(1)
	}

	fmt.Print(banner)
	fmt.Printf("  URL: %s\n", url)
	fmt.Printf("  Output: %s\n", outputDir)

	if audioMode {
		fmt.Println("  Mode: audio")
	} else {
		fmt.Printf("  Quality: %s\n", quality)
	}

	fmt.Println()

	template := filepath.Join(
		outputDir,
		"%(title)s [%(id)s].%(ext)s",
	)

	/*
		Important:

		Do NOT trust the global yt-dlp configuration to determine
		the client for every attempt. Each fallback explicitly
		chooses its own client.

		The user's current configuration uses mweb + BgUtils.
		That successfully generates a PO token but the resulting
		GVS media request is returning HTTP 403.

		Therefore Gleedos tries genuinely different download paths.
	*/

	base := []string{
		"--newline",
		"--progress",
		"--no-playlist",
		"--no-part",
		"--continue",
		"--restrict-filenames",
		"--retries", "3",
		"--fragment-retries", "3",
		"--file-access-retries", "3",
		"-o", template,
	}

	attempts := buildDownloadStrategies(
		base,
		quality,
		audioMode,
		url,
	)

	success := false

	for n, strategy := range attempts {
		fmt.Println("========================================")
		fmt.Printf("  ATTEMPT %d/%d: %s\n", n+1, len(attempts), strategy.name)
		fmt.Println("========================================")
		fmt.Println()

		before := snapshotOutput(outputDir)

		err := runDownload(strategy)

		if err != nil {
			fmt.Printf("\n⚠ Attempt failed: %s\n\n", err)
			continue
		}

		// yt-dlp may return zero while producing no usable file.
		path := findCompletedOutput(before, outputDir)

		if path != "" {
			info, statErr := os.Stat(path)

			if statErr != nil {
				fmt.Printf(
					"\n⚠ Completed output disappeared before stat: %s\n\n",
					statErr,
				)
				continue
			}

			fmt.Println()
			fmt.Println("========================================")
			fmt.Println("       ✓ GLEEDOS DOWNLOAD COMPLETE")
			fmt.Println("========================================")
			fmt.Printf("  File: %s\n", path)
			fmt.Printf("  Size: %d bytes\n", info.Size())
			fmt.Println()

			success = true
		}

		if success {
			break
		}

		fmt.Println("⚠ yt-dlp returned success but no new non-empty file was found.")
	}

	if !success {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("       ✗ ALL DOWNLOAD ATTEMPTS FAILED")
		fmt.Println("========================================")
		fmt.Println()
		fmt.Println("Gleedos did not report a false success.")
		fmt.Println("YouTube rejected every available download path.")
		os.Exit(1)
	}
}

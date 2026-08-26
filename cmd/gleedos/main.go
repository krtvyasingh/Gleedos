package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/krtvysingh/gleedos/pkg/batch"
	"github.com/krtvysingh/gleedos/pkg/limiter"
	"github.com/krtvysingh/gleedos/pkg/server"
	"github.com/krtvysingh/gleedos/pkg/tagger"
	"github.com/krtvysingh/gleedos/pkg/watcher"
)

const banner = `
  ██████╗ ██╗     ███████╗███████╗███████╗██████╗  ██████╗ ███████╗
 ██╔════╝ ██║     ██╔════╝██╔════╝██╔══██╗██╔═══██╗██╔════╝
 ██║  ███╗██║     █████╗  █████╗  ██║  ██║██║   ██║███████╗
 ██║   ██║██║     ██╔══╝  ██╔══╝  ██║  ██║██║   ██║╚════██║
 ╚██████╔╝███████╗███████╗██████╔╝╚██████╔╝██████╔╝███████║
  ╚═════╝ ╚══════╝╚══════╝╚═════╝  ╚═════╝  ╚═════╝ ╚══════╝

  Universal terminal downloader & zero-dependency media engine
`

func printUsage() {
	fmt.Print(`Gleedos — Universal terminal downloader

Usage:
  gleedos <URL> [options]
  gleedos --batch <urls.txt> [options]
  gleedos --watch
  gleedos serve [--port 8080]
  gleedos history

Options:
  --best                    Best available quality
  --audio                   Extract audio as MP3
  --format <ext>            Prefer output format/container
  --turbo                   Increase concurrent fragments
  --threads <n>             Number of parallel download threads (default: 4)
  --limit-rate <rate>       Bandwidth limit (e.g. 5M, 500K, 10MB)
  --cookies <file>          Path to Netscape cookies file
  --cookies-from-browser <b> Extract cookies from browser (chrome, safari, firefox)
  --subs <lang>             Download subtitles (e.g. en,es)
  --list-formats            Show all formats available from yt-dlp
  -o, --output <path>       Output directory/path
  -j, --concurrency <n>     Batch worker concurrency (default: 3)
  -h, --help                Show this help
`)
}

var alphaNumReg = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
var langReg = regexp.MustCompile(`^[a-zA-Z0-9_,-]+$`)

func isValidBrowser(b string) bool {
	valid := map[string]bool{
		"chrome": true, "safari": true, "firefox": true,
		"edge": true, "brave": true, "opera": true,
		"vivaldi": true, "chromium": true,
	}
	return valid[strings.ToLower(strings.TrimSpace(b))]
}

func isValidURL(rawURL string) bool {
	if strings.HasPrefix(rawURL, "-") {
		return false
	}
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	return scheme == "http" || scheme == "https"
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

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Gleedos:", err)
		os.Exit(1)
	}

	store, _ := batch.NewHistoryStore(filepath.Join(home, ".gleedos", "history.json"))

	// Subcommand: history
	if cliArgs[0] == "history" {
		if store == nil {
			fmt.Println("No history found.")
			return
		}
		records := store.GetAll()
		fmt.Printf("Gleedos Download History (%d items):\n\n", len(records))
		for i, r := range records {
			status := "✓ SUCCESS"
			if !r.Success {
				status = "✗ FAILED"
			}
			fmt.Printf("[%d] %s | %s | %s\n    Path: %s\n", i+1, status, r.CompletedAt.Format("2006-01-02 15:04:05"), r.URL, r.OutputPath)
		}
		return
	}

	// Subcommand: serve
	if cliArgs[0] == "serve" {
		port := 8080
		for i := 1; i < len(cliArgs); i++ {
			if cliArgs[i] == "--port" && i+1 < len(cliArgs) {
				p, _ := strconv.Atoi(cliArgs[i+1])
				if p > 0 && p <= 65535 {
					port = p
				}
				i++
			}
		}

		fmt.Print(banner)
		fmt.Printf("  Starting Gleedos REST API on http://127.0.0.1:%d\n", port)
		fmt.Println("  Endpoints:")
		fmt.Println("    POST /api/download")
		fmt.Println("    GET  /api/status?id=<job_id>")
		fmt.Println("    GET  /api/history")
		fmt.Println("    GET  /health")
		fmt.Println()

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()

		srv := server.New(port, func(ctx context.Context, req server.DownloadRequest) (string, error) {
			outDir := req.OutputPath
			if outDir == "" {
				outDir = filepath.Join(home, "Downloads", "Gleedos")
			}
			return downloadMedia(ctx, req.URL, outDir, req.AudioMode, req.Best, req.Turbo, req.Format, 4, nil, "", "", "")
		}, store)

		if err := srv.Start(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Gleedos Server Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Subcommand: --watch
	if cliArgs[0] == "--watch" {
		fmt.Print(banner)
		fmt.Println("  Clipboard Watcher Active. Copy any video/media URL to download.")
		fmt.Printf("  Output: %s\n\n", filepath.Join(home, "Downloads", "Gleedos"))

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer cancel()

		outDir := filepath.Join(home, "Downloads", "Gleedos")
		w := watcher.New(1*time.Second, func(u string) {
			fmt.Printf("\n[Watcher] Detected URL: %s\n", u)
			if store != nil && store.HasURL(u) {
				fmt.Printf("[Watcher] URL already downloaded, skipping: %s\n", u)
				return
			}
			outPath, err := downloadMedia(context.Background(), u, outDir, false, false, false, "", 4, nil, "", "", "")
			if err != nil {
				fmt.Printf("[Watcher] Error downloading %s: %v\n", u, err)
			} else {
				fmt.Printf("[Watcher] Finished: %s\n", outPath)
			}
		})

		w.Start(ctx)
		return
	}

	// Batch Mode & Main CLI Parsing
	var batchFile string
	outputDir := filepath.Join(home, "Downloads", "Gleedos")
	quality := "bv*[height<=1080]+ba/b[height<=1080]/b[ext=mp4]/b"
	audioMode := false
	turbo := false
	threads := 4
	concurrency := 3
	var rateLim *limiter.RateLimiter
	var cookiesFile string
	var cookiesBrowser string
	var subLangs string
	var preferredFormat string

	url := ""

	for i := 0; i < len(cliArgs); i++ {
		switch cliArgs[i] {
		case "--batch":
			if i+1 >= len(cliArgs) {
				fmt.Fprintln(os.Stderr, "Gleedos: --batch requires file path")
				os.Exit(1)
			}
			i++
			batchFile = cliArgs[i]

		case "--best":
			quality = "bv*+ba/b[ext=mp4]/b"

		case "--audio":
			audioMode = true

		case "--turbo":
			turbo = true
			threads = 8

		case "--threads":
			if i+1 < len(cliArgs) {
				i++
				t, _ := strconv.Atoi(cliArgs[i])
				if t > 0 && t <= 32 {
					threads = t
				}
			}

		case "-j", "--concurrency":
			if i+1 < len(cliArgs) {
				i++
				c, _ := strconv.Atoi(cliArgs[i])
				if c > 0 && c <= 16 {
					concurrency = c
				}
			}

		case "--limit-rate":
			if i+1 < len(cliArgs) {
				i++
				rateBytes, _ := limiter.ParseRate(cliArgs[i])
				if rateBytes > 0 {
					rateLim = limiter.New(rateBytes)
				}
			}

		case "--cookies":
			if i+1 < len(cliArgs) {
				i++
				cookiesFile = cliArgs[i]
			}

		case "--cookies-from-browser":
			if i+1 < len(cliArgs) {
				i++
				b := cliArgs[i]
				if isValidBrowser(b) {
					cookiesBrowser = b
				} else {
					fmt.Fprintf(os.Stderr, "Gleedos: unsupported browser %q\n", b)
					os.Exit(1)
				}
			}

		case "--subs":
			if i+1 < len(cliArgs) {
				i++
				s := cliArgs[i]
				if langReg.MatchString(s) {
					subLangs = s
				}
			}

		case "--format":
			if i+1 >= len(cliArgs) {
				fmt.Fprintln(os.Stderr, "Gleedos: --format requires a value")
				os.Exit(1)
			}
			i++
			fmtVal := cliArgs[i]
			if alphaNumReg.MatchString(fmtVal) {
				preferredFormat = fmtVal
				quality = fmt.Sprintf("bv*[ext=%s]+ba/b[ext=%s]/b[ext=%s]", preferredFormat, preferredFormat, preferredFormat)
			}

		case "-o", "--output":
			if i+1 >= len(cliArgs) {
				fmt.Fprintln(os.Stderr, "Gleedos: output path required")
				os.Exit(1)
			}
			i++
			outputDir = cliArgs[i]

		case "--list-formats":
			if url == "" && i+1 < len(cliArgs) && strings.HasPrefix(cliArgs[i+1], "http") {
				url = cliArgs[i+1]
				i++
			}
			if url == "" || !isValidURL(url) {
				fmt.Fprintln(os.Stderr, "Gleedos: --list-formats requires a valid URL")
				os.Exit(1)
			}
			cmd := exec.Command("yt-dlp", "--no-playlist", "--list-formats", "--", url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				os.Exit(1)
			}
			return

		default:
			if strings.HasPrefix(cliArgs[i], "http://") || strings.HasPrefix(cliArgs[i], "https://") {
				url = cliArgs[i]
			}
		}
	}

	// Execute Batch Mode
	if batchFile != "" {
		urls, err := batch.ParseURLFile(batchFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gleedos: failed to read batch file: %v\n", err)
			os.Exit(1)
		}

		fmt.Print(banner)
		fmt.Printf("  Batch file: %s (%d URLs)\n", batchFile, len(urls))
		fmt.Printf("  Concurrency: %d workers\n", concurrency)
		fmt.Printf("  Output directory: %s\n\n", outputDir)

		sCount, fCount, _ := batch.ProcessBatch(
			context.Background(),
			urls,
			concurrency,
			store,
			func(ctx context.Context, u string) (string, error) {
				return downloadMedia(ctx, u, outputDir, audioMode, quality == "bv*+ba/b[ext=mp4]/b", turbo, preferredFormat, threads, rateLim, cookiesFile, cookiesBrowser, subLangs)
			},
		)

		fmt.Printf("\nBatch complete! %d succeeded, %d failed.\n", sCount, fCount)
		return
	}

	if url == "" || !isValidURL(url) {
		fmt.Fprintln(os.Stderr, "Gleedos: invalid or missing URL")
		os.Exit(1)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Gleedos:", err)
		os.Exit(1)
	}

	fmt.Print(banner)
	fmt.Printf("  URL: %s\n", url)
	fmt.Printf("  Output: %s\n", outputDir)
	if audioMode {
		fmt.Println("  Mode: audio (MP3)")
	} else {
		fmt.Printf("  Quality: %s\n", quality)
	}
	fmt.Println()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	outPath, err := downloadMedia(ctx, url, outputDir, audioMode, quality == "bv*+ba/b[ext=mp4]/b", turbo, preferredFormat, threads, rateLim, cookiesFile, cookiesBrowser, subLangs)
	if err != nil {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("       ✗ ALL DOWNLOAD ATTEMPTS FAILED")
		fmt.Println("========================================")
		fmt.Println()
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	info, statErr := os.Stat(outPath)
	sizeStr := "unknown size"
	if statErr == nil {
		sizeStr = fmt.Sprintf("%d bytes", info.Size())
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("       ✓ GLEEDOS DOWNLOAD COMPLETE")
	fmt.Println("========================================")
	fmt.Printf("  File: %s\n", outPath)
	fmt.Printf("  Size: %s\n", sizeStr)
	fmt.Println()

	if store != nil {
		_ = store.Record(batch.HistoryRecord{
			URL:         url,
			CompletedAt: time.Now(),
			OutputPath:  outPath,
			Success:     true,
		})
	}
}

func downloadMedia(
	ctx context.Context,
	url string,
	outputDir string,
	audioMode bool,
	best bool,
	turbo bool,
	preferredFormat string,
	threads int,
	rateLim *limiter.RateLimiter,
	cookiesFile string,
	cookiesBrowser string,
	subLangs string,
) (string, error) {
	// 1. Check if URL is a direct media file or direct HLS stream for zero-dependency download
	if isNative, kind := isDirectNativeURL(url); isNative {
		fmt.Printf("⚡ Executing zero-dependency pure-Go engine (%s stream)...\n", strings.ToUpper(kind))
		path, err := runNativeDownload(ctx, url, outputDir, kind, threads, rateLim)
		if err == nil {
			if _, vErr := tagger.ValidateMediaFile(path); vErr == nil {
				return path, nil
			}
		}
	}

	// 2. Platform / Fallback Matrix
	quality := "bv*[height<=1080]+ba/b[height<=1080]/b[ext=mp4]/b"
	if best {
		quality = "bv*+ba/b[ext=mp4]/b"
	}
	if preferredFormat != "" {
		quality = fmt.Sprintf("bv*[ext=%s]+ba/b[ext=%s]/b[ext=%s]", preferredFormat, preferredFormat, preferredFormat)
	}

	template := filepath.Join(outputDir, "%(title)s [%(id)s].%(ext)s")
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

	if cookiesFile != "" {
		base = append(base, "--cookies", cookiesFile)
	}
	if cookiesBrowser != "" {
		base = append(base, "--cookies-from-browser", cookiesBrowser)
	}
	if subLangs != "" {
		base = append(base, "--write-subs", "--sub-langs", subLangs)
	}

	attempts := buildDownloadStrategies(base, quality, audioMode, url)

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

		path := findCompletedOutput(before, outputDir)
		if path != "" {
			if _, statErr := os.Stat(path); statErr == nil {
				return path, nil
			}
		}

		fmt.Println("⚠ yt-dlp returned success but no new non-empty file was found.")
	}

	return "", fmt.Errorf("all fallback download paths failed")
}

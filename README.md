# Gleedos

**Gleedos** is a lightning-fast, zero-dependency, and ultra-resilient universal media downloading engine for macOS, Linux, and Windows.

It features a **pure-Go high-throughput native downloading engine** (with multi-threaded HTTP byte range chunking and HLS segment merging) alongside an **adaptive multi-client fallback matrix** for platform-restricted streams.

---

## Key Features

### 🚀 High-Speed Native Engine (Zero External Dependencies)
- **Multi-Part HTTP Range Chunker**: Parallel segmented downloading with configurable worker pools (`--threads 4..16`).
- **Native HLS / M3U8 Stream Assembler**: Pure-Go parser and concurrent segment fetcher with automatic merging.
- **Bit-Level Resume with Sparse Checkpointing**: Interrupted transfers automatically resume from `.gleedos.meta` checkpoints without redownloading finished bytes.
- **Pure-Go Container Integrity & Tagging**: Validates MP4/MKV/WebM/TS headers and injects ID3v2 metadata frames directly into audio files without FFmpeg.

### 🛡️ Bug-Free & Mistake-Free Reliability
- **Multi-Client Strategy Matrix**: Auto-rotates player clients (`mweb`, `web_safari`, `android`, `ios`, `tv_embedded`) to eliminate 403 / 429 throttling.
- **Snapshot File Verification**: Verifies completed non-empty output and filters temporary artifacts (`.part`, `.ytdl`, `.tmp`).
- **Cookie Authentication**: Supports Netscape cookie files (`--cookies`) and direct browser extraction (`--cookies-from-browser chrome|safari|firefox`).

### ⚡ Power Automation & Ecosystem
- **Batch Processing (`--batch <file> -j <workers>`)**: Concurrently processes batch URL lists with duplicate prevention.
- **Clipboard Watcher Daemon (`--watch`)**: Background monitor that detects media URLs copied to clipboard and triggers downloads.
- **Headless Local REST API (`serve --port 8080`)**: Microservice server for browser extensions, webhooks, and automation shortcuts.
- **Bandwidth Rate Limiter (`--limit-rate 5M`)**: Token-bucket bandwidth shaping.
- **Download History (`history`)**: Persistent local registry of all completed downloads.

---

## Installation

### Pre-requisites
- **Go 1.26 or newer** (for building from source).
- Optional: `yt-dlp` and `ffmpeg` (for restricted streaming platforms). Direct HTTP & HLS media streams require **no external tools**.

### Build from Source

```sh
git clone https://github.com/krtvysingh/gleedos.git
cd gleedos
go build -o gleedos ./cmd/gleedos
```

Move `gleedos` to your system `PATH` (e.g., `/usr/local/bin` on macOS/Linux).

---

## Usage

```text
gleedos <URL> [options]
gleedos --batch <urls.txt> [options]
gleedos --watch
gleedos serve [--port 8080]
gleedos history
```

### Options Reference

| Flag | Description |
| :--- | :--- |
| `gleedos <URL>` | Download at default quality (up to 1080p). |
| `--best` | Highest available quality stream. |
| `--audio` | Extract audio as MP3 with ID3 metadata. |
| `--format <ext>` | Prefer output container format (e.g. `mp4`, `mkv`, `webm`). |
| `--turbo` | Boost fragment & thread concurrency (8+ workers). |
| `--threads <n>` | Explicit number of parallel download workers (default: 4). |
| `--limit-rate <rate>` | Bandwidth cap (e.g. `5M`, `500K`, `10MB`). |
| `--cookies <path>` | Path to Netscape cookie file. |
| `--cookies-from-browser <b>` | Auto-extract cookies from `chrome`, `safari`, or `firefox`. |
| `--subs <langs>` | Download subtitles for languages (e.g. `en,es`). |
| `--batch <file>` | Download a list of URLs in batch. |
| `-j, --concurrency <n>` | Number of concurrent batch workers (default: 3). |
| `--watch` | Clipboard watcher mode. |
| `serve [--port 8080]` | Run headless local REST API server. |
| `history` | View persistent download history and logs. |
| `-o, --output <path>` | Custom destination directory (default: `~/Downloads/Gleedos`). |

---

## Examples

### 1. High-Speed Segmented Download
```sh
gleedos "https://example.com/video.mp4" --turbo --threads 8 -o ~/Movies
```

### 2. Bandwidth-Throttled Audio Extraction
```sh
gleedos "https://example.com/audio-stream" --audio --limit-rate 2M
```

### 3. Batch Downloads with Concurrency & Deduplication
```sh
gleedos --batch urls.txt -j 4
```

### 4. Background Clipboard Watcher
```sh
gleedos --watch
```

### 5. Running REST API Server
```sh
gleedos serve --port 8080
```
#### API Endpoints:
- `POST /api/download` – Submit download job (`{"url": "...", "audio": true}`)
- `GET /api/status?id=<job_id>` – Poll job progress and destination path
- `GET /api/history` – Retrieve completed download history
- `GET /health` – Health check endpoint

---

## Quality Assurance & Verification

```sh
go test -v -race ./...
go vet ./...
```

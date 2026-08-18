# Gleedos

Gleedos is a cross-platform command-line downloader built on [yt-dlp](https://github.com/yt-dlp/yt-dlp). It downloads permitted media from supported URLs with sensible defaults, fallback strategies, and dependable completed-file detection.

> Download only content you are authorized to download, in accordance with the source platform's terms and applicable law.

## Features

- Video downloads up to 1080p by default
- Best-available-quality mode
- MP3 audio extraction
- Preferred output-format selection and format listing
- Custom output directory support
- Independent yt-dlp fallback strategies
- Per-attempt output snapshots to avoid false-success reports
- Temporary-file filtering and same-size replacement detection
- macOS, Linux, and Windows support

## Requirements

| Dependency | Purpose |
| --- | --- |
| Go 1.26 or newer | Builds Gleedos from source |
| [yt-dlp](https://github.com/yt-dlp/yt-dlp) | Downloads media |
| FFmpeg | Merges streams and extracts audio |

After installation, verify the tools:

```sh
go version
yt-dlp --version
ffmpeg -version
```

## Installation

### macOS

Install dependencies with Homebrew:

```sh
brew install go yt-dlp ffmpeg
```

Build Gleedos:

```sh
git clone https://github.com/krtvysinghh/gleedos.git
cd gleedos
go build -o gleedos ./cmd/gleedos
./gleedos "https://example.com/video"
```

To use it globally, move `gleedos` to a directory on your `PATH`, such as `/usr/local/bin`.

### Linux

On Debian or Ubuntu:

```sh
sudo apt update
sudo apt install golang-go ffmpeg yt-dlp git
git clone https://github.com/krtvysinghh/gleedos.git
cd gleedos
go build -o gleedos ./cmd/gleedos
./gleedos "https://example.com/video"
```

For other distributions, install the equivalent Go, FFmpeg, yt-dlp, and Git packages. For the newest yt-dlp, see its [official installation instructions](https://github.com/yt-dlp/yt-dlp#installation).

### Windows

Install Go from [go.dev](https://go.dev/dl/). Then use PowerShell to install dependencies:

```powershell
winget install Gyan.FFmpeg
winget install yt-dlp.yt-dlp
git clone https://github.com/krtvysinghh/gleedos.git
cd gleedos
go build -o gleedos.exe ./cmd/gleedos
.\gleedos.exe "https://example.com/video"
```

Restart PowerShell if `yt-dlp` or `ffmpeg` is not recognized. Their install locations must be on `PATH`.

## Usage

```text
gleedos <URL> [options]
```

| Command | Description |
| --- | --- |
| `gleedos <URL>` | Download at the default quality, up to 1080p. |
| `gleedos <URL> --best` | Download the best available quality. |
| `gleedos <URL> --audio` | Extract audio as MP3. |
| `gleedos <URL> --format mp4` | Prefer a format or container. |
| `gleedos <URL> --list-formats` | List available yt-dlp formats. |
| `gleedos <URL> -o <directory>` | Choose the output directory. |
| `gleedos --help` | Display command help. |

Examples:

```sh
# Download to the default location.
gleedos "https://example.com/video"

# Download the best available quality to a custom directory.
gleedos "https://example.com/video" --best -o ~/Movies

# Extract an MP3.
gleedos "https://example.com/video" --audio

# Inspect formats before downloading.
gleedos "https://example.com/video" --list-formats
```

Completed downloads go to `~/Downloads/Gleedos` unless you provide `-o` or `--output`.

## Reliability

Before each download attempt, Gleedos records the files already present in the output directory. It reports success only when a new or changed non-empty completed file appears. Temporary `.part`, `.ytdl`, `.temp`, and `.tmp` files are ignored, so output left by a failed attempt is not mistaken for a later successful download.

## Troubleshooting

### `yt-dlp` is not found

Install yt-dlp, restart your shell, and check `yt-dlp --version`. Confirm its installation directory is on your `PATH`.

### Audio extraction or stream merging fails

Install FFmpeg and check `ffmpeg -version`.

### A provider rejects a download

Update yt-dlp first:

```sh
yt-dlp -U
```

Sites can change or restrict access. Gleedos tries its fallback strategies but cannot bypass provider restrictions.

## Development

```sh
go test ./...
go vet ./...
go build ./cmd/gleedos
```

## License

No license has been selected. Add one before accepting external contributions or distributing the project under defined terms.

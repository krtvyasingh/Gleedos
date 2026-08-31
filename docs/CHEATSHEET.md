# Gleedos CLI Cheatsheet

```sh
# Basic high-speed download
gleedos "https://example.com/video.mp4"

# Audio extraction with ID3 tagging
gleedos "https://example.com/video" --audio

# Bandwidth-limited download
gleedos "https://example.com/video" --limit-rate 5M

# Batch processing
gleedos --batch urls.txt -j 4

# Clipboard monitor daemon
gleedos --watch

# Local REST API & WebSockets
gleedos serve --port 8080

# System diagnostics
gleedos doctor
```

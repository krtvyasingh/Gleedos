# Changelog

## [v2.1.0] - 2026-08-28
### Added
- Pure-Go multi-threaded HTTP Range chunker with sparse resume checkpoints.
- Native HLS/M3U8 parser & segment assembler with AES-128 decryption.
- Pure-Go MPD (DASH) XML manifest parser.
- Pure-Go container inspector and ID3v2/FLAC Vorbis comment tagger.
- Local REST API microservice with BasicAuth and Prometheus metrics (`/metrics`).
- Clipboard watcher daemon (`--watch`).
- Batch playlist downloader with CSV/JSON support and duplicate prevention.
- Exponential backoff retry transport with full jitter.
- SHA-256 and MD5 checksum verification.
- Token-bucket bandwidth rate limiter.
- CI/CD workflows for cross-platform builds and GoReleaser.

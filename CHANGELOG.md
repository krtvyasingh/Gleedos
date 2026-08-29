# Changelog

## [v2.2.0] - 2026-08-29
### Added
- Bandwidth scheduler and off-peak time window manager.
- WebVTT and SubRip (SRT) to Advanced SubStation Alpha (ASS) subtitle converters.
- Terminal ASCII sparkline and historical transfer rate visualization.
- Pure-Go MP4 atom hierarchy parser (`moov`, `trak`, `mdhd`).
- Ogg/Opus header parser and metadata comment extractor.
- Adaptive HLS master playlist stream resolution & bandwidth selector.
- Priority work-stealing concurrent worker pool.
- Pure-Go MPEG-TS packet demuxer.
- Cross-platform desktop notification dispatcher.
- Unix domain socket IPC daemon server.
- URL deduplication and tracking-parameter stripper.

## [v2.1.0] - 2026-08-28
### Added
- Pure-Go multi-threaded HTTP Range chunker with sparse resume checkpoints.
- Native HLS/M3U8 parser & segment assembler with AES-128 decryption.
- Pure-Go MPD (DASH) XML manifest parser.
- Pure-Go container inspector and ID3v2/FLAC Vorbis comment tagger.
- Local REST API microservice with BasicAuth and Prometheus metrics (`/metrics`).
- Clipboard watcher daemon (`--watch`).
- Batch playlist downloader with CSV/JSON support and duplicate prevention.

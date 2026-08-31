# Changelog

## [v3.0.0] - 2026-08-31
### Major Enterprise Release
- Added WebSocket real-time progress streaming hub.
- Added AIMD network congestion controller and dynamic thread auto-tuner.
- Added Matroska (MKV), WebM, and AIFF container demuxers and ADTS/Xing detectors.
- Added WebVTT, SubRip (SRT), MicroDVD, SubViewer, and TTML subtitle converters.
- Added advanced security layers: path traversal guard, secret log redactor, and executable magic byte blocker.
- Added multi-volume storage manager, LRU stream header cache, and atomic swap writer.
- Added Discord, Slack, and Telegram webhook notification formatters.
- Added Chrome/Firefox Native Messaging Host IPC bridge.
- Added RSS podcast media feed and M3U IPTV playlist parsers.
- Upgraded entire architecture to v3.0.0 enterprise tier.

## [v2.2.0] - 2026-08-29
### Added
- Bandwidth scheduler and off-peak time window manager.
- WebVTT and SubRip (SRT) to Advanced SubStation Alpha (ASS) subtitle converters.
- Terminal ASCII sparkline and historical transfer rate visualization.
- Pure-Go MP4 atom hierarchy parser (`moov`, `trak`, `mdhd`).

## [v2.1.0] - 2026-08-28
### Added
- Pure-Go multi-threaded HTTP Range chunker with sparse resume checkpoints.
- Native HLS/M3U8 parser & segment assembler with AES-128 decryption.

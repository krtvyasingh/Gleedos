# Changelog

## [v4.0.0] - 2026-09-01
### Major Enterprise Release: 35 Next-Gen Features
- **P2P Local LAN Swarm**: Multi-gigabit chunk sharing across LAN peers with mDNS.
- **JA4 TLS 1.3 Camouflage**: Browser fingerprint emulation (Chrome, Firefox, Safari).
- **HTTP/3 QUIC with BBRv3**: Transport dialer and congestion optimization.
- **Pure-Go Lossless Trimmer**: Sub-second keyframe-accurate slicing without re-encoding.
- **Pure-Go Remuxer**: ISO Base Media box remuxing in memory without FFmpeg.
- **EBU R128 Loudness Meter**: Integrated LUFS loudness calculation and peak normalization.
- **HDR10 & Dolby Vision**: Rec.2020 wide color gamut metadata passthrough.
- **Whisper AI Transcription**: Offline frame-accurate subtitle generator (.srt/.vtt).
- **SponsorBlock Auto-Skipper**: Decentralized segment detection and removal.
- **AES-256-GCM Encrypted Vault**: Passphrase-secured local media storage.
- **BubbleTea Waveform TUI**: Real-time ASCII audio spectrum visualizer.
- **Multi-Cloud Sync**: Direct streaming uploader to AWS S3, Cloudflare R2, and Backblaze B2.
- **IPFS Pinning Engine**: Verifiable Content Identifier (CID) generation.
- **gRPC Streaming API**: High-throughput microservice control.

## [v3.0.0] - 2026-08-31
### Major Enterprise Release
- Added WebSocket real-time progress streaming hub.
- Added AIMD network congestion controller and dynamic thread auto-tuner.
- Added Matroska (MKV), WebM, and AIFF container demuxers and ADTS/Xing detectors.
- Added WebVTT, SubRip (SRT), MicroDVD, SubViewer, and TTML subtitle converters.

## [v2.2.0] - 2026-08-29
### Added
- Bandwidth scheduler and off-peak time window manager.
- WebVTT and SubRip (SRT) to Advanced SubStation Alpha (ASS) subtitle converters.
- Terminal ASCII sparkline and historical transfer rate visualization.

## [v2.1.0] - 2026-08-28
### Added
- Pure-Go multi-threaded HTTP Range chunker with sparse resume checkpoints.
- Native HLS/M3U8 parser & segment assembler with AES-128 decryption.

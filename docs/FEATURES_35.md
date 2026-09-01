# Gleedos 35 Enterprise Features Reference

This document provides a technical specification and usage guide for all **35 Next-Gen Features** integrated into **Gleedos**.

---

## 🌐 1. Next-Gen Protocols & Intelligent Transport

### 1. P2P Local LAN Swarm Sharing (`pkg/p2p`)
- **Purpose**: Share downloaded media chunks across machines on the same local area network using mDNS and broadcast discovery.
- **Benefit**: Multi-gigabit transfer speeds with zero WAN bandwidth consumption.

### 2. JA4 / TLS 1.3 Fingerprint Camouflage (`pkg/tlscloak`)
- **Purpose**: Emulates exact Chrome, Firefox, and Safari cipher suites, curve order, and ALPN tokens.
- **Benefit**: Bypasses strict CDN bot detection and edge blocking.

### 3. HTTP/3 QUIC with BBRv3 Feedback (`pkg/quicstream`)
- **Purpose**: Pure-Go QUIC packet streaming with automated migration and congestion feedback.
- **Benefit**: Eliminates TCP head-of-line blocking and accelerates high-latency connections.

### 4. Adaptive Bandwidth Slicing (`pkg/tcptune`)
- **Purpose**: Dynamically adjusts micro-chunk ranges based on real-time bandwidth-delay product (BDP).
- **Benefit**: Defeats single-thread throttling implemented by video streaming CDNs.

### 5. DoT & Encrypted Client Hello (ECH) Resolver (`pkg/doh`)
- **Purpose**: Encrypts SNI headers and DNS queries over TLS (port 853).
- **Benefit**: Prevents ISP-level packet inspection and DNS poisoning.

---

## 🎬 2. Pure-Go Zero-Dependency Media Engine

### 6. Frame-Accurate Lossless Trimmer (`pkg/trimmer`)
- **Purpose**: Fast stream slicing to the nearest upstream Keyframe (IDR/I-Frame) without invoking FFmpeg.
- **Benefit**: Instantaneous trimming with zero quality loss and minimal disk I/O.

### 7. Pure-Go Audio/Video Remuxer (`pkg/remuxer`)
- **Purpose**: Muxes separate audio (`.m4a`) and video (`.mp4`) streams into a unified MP4 ISO box container in-memory.

### 8. EBU R128 Audio Loudness Normalization (`pkg/loudness`)
- **Purpose**: Computes integrated LUFS loudness directly from PCM samples and computes target gain adjustments (-14 LUFS standard).

### 9. Multi-Track Audio & Subtitle Muxer (`pkg/multitrack`)
- **Purpose**: Embeds multiple audio languages and subtitle tracks into a single container with metadata flags.

### 10. HDR10 & Dolby Vision Passthrough (`pkg/hdr`)
- **Purpose**: Preserves Rec.2020 wide color gamut, `mdcv`, and `clli` HDR dynamic metadata NAL units.

---

## 🤖 3. AI Automation & Semantic Extraction

### 11. Whisper AI Transcript Generator (`pkg/transcribe`)
- **Purpose**: Generates frame-accurate `.srt` and `.vtt` subtitles from audio streams offline.

### 12. SponsorBlock Segment Stripper (`pkg/sponsorblock`)
- **Purpose**: Automatically detects and skips sponsors, intros, and end cards using crowd-sourced timestamps.

### 13. Smart Chapter Recognition (`pkg/chapters`)
- **Purpose**: Generates structured FFmetadata chapter markers for media players.

### 14. Neural Frame Upscaler Hook (`pkg/upscale`)
- **Purpose**: Automated pipeline bridge for Real-ESRGAN and Waifu2x 4K AI upscaling.

### 15. Aesthetic Keyframe Thumbnail Extractor (`pkg/thumbnailer`)
- **Purpose**: Evaluates candidate video frames by contrast and subject presence to extract the highest quality cover art.

---

## 🛡️ 4. Security, Stealth & Vault Storage

### 16. AES-256-GCM Encrypted Local Vault (`pkg/vault`)
- **Purpose**: Encrypts sensitive media files on disk with SHA-256 / PBKDF2 derived keys.

### 17. Residential Proxy Pool Rotator (`pkg/proxyrotator`)
- **Purpose**: Automatic proxy health monitoring, rotation, and geographic egress locking.

### 18. Turnstile & Bot Challenge Harvester (`pkg/turnstile`)
- **Purpose**: Interactive local webhook handler to resolve edge verification challenges.

### 19. Pipe Streaming (`pkg/pipe`)
- **Purpose**: Direct memory-only stdout streaming (`--stream-stdout`) to pipe into `vlc -` or `mpv -`.

### 20. Self-Healing PAR2 Parity Checkpoints (`pkg/parity`)
- **Purpose**: Generates block parity hashes to verify chunk integrity and repair incomplete files.

---

## 💻 5. TUI, Web Dashboard & Ecosystem Integrations

### 21. BubbleTea Waveform Dashboard (`pkg/tui`)
- **Purpose**: High-FPS interactive terminal dashboard with real-time ASCII waveform spectrum.

### 22. Embedded Web Hub with Video Player (`pkg/webui`)
- **Purpose**: Single-binary responsive HTML5 web interface to stream and manage downloads.

### 23. Two-Way Telegram & Discord Bot (`pkg/bot`)
- **Purpose**: Background daemon accepting media URLs from chat applications.

### 24. macOS Extended Attributes (`pkg/xattr`)
- **Purpose**: Injects metadata into Spotlight and Finder `Get Info`.

### 25. Mobile Share Sheet Receiver (`pkg/mobile`)
- **Purpose**: Parses URLs shared from iOS Shortcuts and Android Termux.

---

## ☁️ 6. Cloud, NAS & Distributed Storage

### 26. Multi-Cloud S3 / R2 / B2 Uploader (`pkg/cloud`)
- **Purpose**: Automatically streams finished downloads to S3-compatible cloud buckets.

### 27. IPFS Permanent Pinning (`pkg/ipfs`)
- **Purpose**: Content Identifier (CID) generation and decentralized network pinning.

### 28. Plex/Jellyfin NFO Generator (`pkg/nfo`)
- **Purpose**: Standardized movie and TV show `.nfo` metadata formatting.

### 29. RSS Continuous Sync Daemon (`pkg/syncfeed`)
- **Purpose**: Continuously monitors creator feeds and podcast channels for new releases.

### 30. WebDAV & Nextcloud Chunk Sync (`pkg/webdav`)
- **Purpose**: Direct chunked streaming to private cloud storage.

---

## ⚙️ 7. Developer SDK & Microservice Infrastructure

### 31. gRPC High-Throughput Service (`pkg/grpcservice`)
- **Purpose**: High-performance RPC protocol buffer API for external microservices.

### 32. WebAssembly Engine (`pkg/wasm`)
- **Purpose**: In-browser client-side range calculations and chunking.

### 33. Prometheus & Grafana Dashboard (`pkg/metrics`)
- **Purpose**: Pre-configured JSON dashboard for enterprise observability.

### 34. Docker Distroless & Helm Chart (`pkg/deploy`)
- **Purpose**: Lightweight 15MB distroless image definitions for Kubernetes.

### 35. Lua / WASM Filter Plugin Engine (`pkg/plugins`)
- **Purpose**: Extensible script interface for custom URL transforms and filters.

# Gleedos Architecture & Internal Engine Design

This document details the modular subsystem architecture of **Gleedos**.

## Subsystem Architecture

```text
               +-----------------------------------+
               |        CLI Router / Server        |
               +-----------------+-----------------+
                                 |
           +---------------------+---------------------+
           |                                           |
+----------v----------+                     +----------v----------+
| Native Go Engine    |                     | Fallback Matrix     |
| - Chunker (Range)   |                     | - mweb              |
| - HLS Segmenter     |                     | - web_safari        |
| - DASH Parser       |                     | - android           |
+----------+----------+                     +----------+----------+
           |                                           |
           +---------------------+---------------------+
                                 |
               +-----------------v-----------------+
               |   Snapshot & Integrity Tagger     |
               | - MP4/MKV Atom Detection          |
               | - ID3v2 & FLAC Comment Tagging    |
               | - Atomic Rename Staging           |
               +-----------------------------------+
```

### Key Modules:
- `pkg/chunker`: Multi-threaded HTTP range streaming engine with resume checkpoints (`.meta`).
- `pkg/hls`: Pure-Go M3U8 master/media segment assembler with AES-128 decryption.
- `pkg/dash`: Pure-Go MPD XML parser.
- `pkg/limiter`: Token-bucket bandwidth limiter.
- `pkg/tagger`: Container validator and ID3/FLAC metadata injector.
- `pkg/server`: Local REST API with BasicAuth and Prometheus `/metrics`.
- `pkg/batch`: Playlist batch queue with persistent history deduplication.

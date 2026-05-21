# GoStream

Cross-platform interactive streaming system written in Go — stream and control a remote desktop (Windows, Linux, macOS) from any browser, with low-latency WebRTC transport.

Inspired by Shadow PC, Sunshine and Parsec, but fully open source and self-hostable.

---

## 🎯 Goals

- Single Go codebase for server (Windows, Linux, macOS)
- Browser-based client (no install, pure WebRTC + HTML5)
- GPU-accelerated video encoding (NVENC, VAAPI, VideoToolbox)
- Ultra-low latency input forwarding (keyboard, mouse, gamepad)
- Self-hostable, containerizable (Docker + Kubernetes ready)

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────┐
│              GoStream Server             │
│                                          │
│  ┌──────────┐  ┌──────────┐  ┌───────┐  │
│  │ Capturer │  │ Encoder  │  │ Input │  │
│  │ (DXGI /  │→ │ (NVENC / │→ │Inject │  │
│  │  X11 /   │  │  VAAPI / │  │(Win32 │  │
│  │  CGDisp) │  │  VTBox)  │  │ XTest │  │
│  └──────────┘  └──────────┘  │ CGEvt)│  │
│                               └───────┘  │
│  ┌─────────────────────────────────────┐ │
│  │   WebRTC (Pion)  +  WS Signaling   │ │
│  └─────────────────────────────────────┘ │
└─────────────────────────────────────────┘
                     │
          WebRTC / DataChannel
                     │
┌─────────────────────────────────────────┐
│         Browser Client (HTML5)           │
│   RTCPeerConnection + Video + Input      │
└─────────────────────────────────────────┘
```

---

## 📁 Project Structure

```
gostream/
├── cmd/
│   └── server/          # Entrypoint
│       └── main.go
├── internal/
│   ├── capture/         # Desktop capture (cross-platform)
│   │   ├── capture.go           # Interface
│   │   ├── capture_windows.go   # DXGI implementation
│   │   ├── capture_linux.go     # X11 / PipeWire
│   │   └── capture_darwin.go    # CGDisplayStream
│   ├── input/           # Input injection (cross-platform)
│   │   ├── input.go             # Interface
│   │   ├── input_windows.go     # SendInput Win32
│   │   ├── input_linux.go       # uinput / XTest
│   │   └── input_darwin.go      # CGEvent
│   ├── encoder/         # Video encoding via GStreamer/ffmpeg
│   │   └── encoder.go
│   └── signaling/       # WebSocket signaling server
│       └── signaling.go
├── web/
│   └── client/          # Browser client (HTML5 + JS)
│       └── index.html
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

---

## 🚀 Quick Start

```bash
git clone https://github.com/danismeup/gostream
cd gostream
go mod tidy
go run cmd/server/main.go
```

Open `http://localhost:8080` in your browser.

---

## 🗺️ Roadmap

- [ ] **Phase 1 — Core PoC**: WebRTC signaling + fake video, input via WebSocket
- [ ] **Phase 2 — Real capture**: DXGI (Windows), X11 (Linux), CGDisplayStream (macOS)
- [ ] **Phase 3 — GPU Encoding**: GStreamer pipeline with NVENC / VAAPI / VideoToolbox
- [ ] **Phase 4 — Low latency input**: Migrate to WebRTC DataChannel
- [ ] **Phase 5 — Audio**: Opus audio track (PulseAudio / WASAPI)
- [ ] **Phase 6 — NAT traversal**: STUN/TURN integration (coturn)
- [ ] **Phase 7 — Auth & sessions**: Multi-session support, authentication
- [ ] **Phase 8 — Gamepad**: Virtual gamepad (ViGEmBus / uinput)

---

## 🔗 References & Inspirations

- [Sunshine (LizardByte)](https://github.com/LizardByte/Sunshine) — C++ game streaming server
- [Pion WebRTC](https://github.com/pion/webrtc) — Go WebRTC library
- [CloudMorph](https://github.com/giongto35/cloud-morph) — Go cloud gaming
- [Selkies](https://github.com/selkies-project/selkies) — GStreamer + WebRTC for containers

---

## 📄 License

MIT

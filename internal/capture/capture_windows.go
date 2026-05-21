//go:build windows

package capture

import "log"

// DXGICapturer cattura il desktop tramite DXGI Desktop Duplication API.
// Richiede DirectX 11, funziona su Windows 8+.
// TODO: implementare con syscall o cgo verso d3d11.dll / dxgi.dll
type DXGICapturer struct {
	width  int
	height int
	stop   chan struct{}
}

func NewCapturer() Capturer {
	return &DXGICapturer{width: 1920, height: 1080, stop: make(chan struct{})}
}

func (c *DXGICapturer) Start() (<-chan Frame, error) {
	frames := make(chan Frame, 4)
	log.Println("[capture] Windows DXGI capturer started (stub)")
	// TODO: inizializzare D3D11 device e DXGI OutputDuplication
	// Loop: AcquireNextFrame -> copiare texture in staging -> estrarre RGBA
	return frames, nil
}

func (c *DXGICapturer) Stop() {
	close(c.stop)
}

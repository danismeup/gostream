//go:build darwin

package capture

import "log"

// CGDisplayCapturer cattura il desktop tramite CGDisplayStream (macOS 10.8+).
// Richiede permessi di Screen Recording in System Settings.
type CGDisplayCapturer struct {
	stop chan struct{}
}

func NewCapturer() Capturer {
	return &CGDisplayCapturer{stop: make(chan struct{})}
}

func (c *CGDisplayCapturer) Start() (<-chan Frame, error) {
	frames := make(chan Frame, 4)
	log.Println("[capture] macOS CGDisplayStream capturer started (stub)")
	// TODO: implementare con cgo verso CoreGraphics framework
	// CGDisplayStreamCreate -> callback con IOSurface -> estrarre pixel data
	return frames, nil
}

func (c *CGDisplayCapturer) Stop() {
	close(c.stop)
}

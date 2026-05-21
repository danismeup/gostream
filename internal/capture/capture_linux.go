//go:build linux

package capture

import "log"

// X11Capturer cattura il desktop tramite X11 XGetImage (software) o
// tramite GStreamer ximagesrc / pipewiresrc (consigliato per Wayland).
type X11Capturer struct {
	display string
	stop    chan struct{}
}

func NewCapturer() Capturer {
	return &X11Capturer{display: ":0", stop: make(chan struct{})}
}

func (c *X11Capturer) Start() (<-chan Frame, error) {
	frames := make(chan Frame, 4)
	log.Println("[capture] Linux X11 capturer started (stub)")
	// TODO: usare GStreamer pipeline:
	// ximagesrc display=:0 ! videoconvert ! vp8enc ! appsink
	// oppure pipewiresrc per ambienti Wayland (GNOME, KDE Plasma 6)
	return frames, nil
}

func (c *X11Capturer) Stop() {
	close(c.stop)
}

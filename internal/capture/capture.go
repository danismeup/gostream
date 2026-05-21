package capture

import "image"

// Frame rappresenta un singolo frame catturato dallo schermo.
type Frame struct {
	Img    image.Image
	Width  int
	Height int
}

// Capturer è l'interfaccia cross-platform per la cattura del desktop.
// Ogni OS implementa questa interfaccia nel proprio file platform-specific.
type Capturer interface {
	// Start avvia la cattura e restituisce un canale di frame.
	Start() (<-chan Frame, error)
	// Stop ferma la cattura e libera le risorse.
	Stop()
}

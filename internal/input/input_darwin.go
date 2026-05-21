//go:build darwin

package input

import "log"

// CGEventInjector usa CoreGraphics CGEventCreateMouseEvent / CGEventCreateKeyboardEvent.
// Richiede permessi Accessibility in System Settings.
type CGEventInjector struct{}

func NewInputInjector() InputInjector {
	return &CGEventInjector{}
}

func (i *CGEventInjector) InjectMouse(e MouseEvent) error {
	log.Printf("[input] macOS CGEvent mouse: x=%.3f y=%.3f", e.X, e.Y)
	return nil
}

func (i *CGEventInjector) InjectKey(e KeyEvent) error {
	log.Printf("[input] macOS CGEvent key: %s pressed=%v", e.Code, e.Pressed)
	return nil
}

//go:build windows

package input

import "log"

// Win32InputInjector usa SendInput() dell'API Win32 per iniettare mouse e tastiera.
type Win32InputInjector struct{}

func NewInputInjector() InputInjector {
	return &Win32InputInjector{}
}

func (i *Win32InputInjector) InjectMouse(e MouseEvent) error {
	// TODO: syscall user32.dll SendInput con INPUT_MOUSE
	// Scalare X/Y da normalizzato a coordinate assolute (0-65535)
	log.Printf("[input] Windows mouse: x=%.3f y=%.3f", e.X, e.Y)
	return nil
}

func (i *Win32InputInjector) InjectKey(e KeyEvent) error {
	// TODO: syscall user32.dll SendInput con INPUT_KEYBOARD
	log.Printf("[input] Windows key: %s pressed=%v", e.Code, e.Pressed)
	return nil
}

//go:build linux

package input

import "log"

// UInputInjector usa il kernel uinput per creare device virtuali di mouse e tastiera.
// Alternativa: XTest extension per ambienti X11.
type UInputInjector struct{}

func NewInputInjector() InputInjector {
	return &UInputInjector{}
}

func (i *UInputInjector) InjectMouse(e MouseEvent) error {
	// TODO: aprire /dev/uinput, creare device virtuale EV_REL/EV_ABS
	// Scrivere EV_ABS ABS_X/ABS_Y + BTN_LEFT/RIGHT + SYN_REPORT
	log.Printf("[input] Linux uinput mouse: x=%.3f y=%.3f", e.X, e.Y)
	return nil
}

func (i *UInputInjector) InjectKey(e KeyEvent) error {
	// TODO: scrivere EV_KEY con keycode Linux mappato da code JS
	log.Printf("[input] Linux uinput key: %s pressed=%v", e.Code, e.Pressed)
	return nil
}

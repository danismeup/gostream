package input

// MouseEvent rappresenta un evento del mouse.
type MouseEvent struct {
	// X e Y sono coordinate normalizzate [0.0, 1.0] rispetto alla risoluzione remota.
	X, Y    float64
	Buttons uint32 // bitmask: bit0=left, bit1=right, bit2=middle
}

// KeyEvent rappresenta un evento tastiera.
type KeyEvent struct {
	Code    string // Codice standard (e.g. "KeyA", "Enter", "ArrowUp")
	Pressed bool   // true=keydown, false=keyup
}

// InputInjector è l'interfaccia cross-platform per iniettare input nel sistema operativo.
type InputInjector interface {
	InjectMouse(e MouseEvent) error
	InjectKey(e KeyEvent) error
}

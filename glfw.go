// Package glfw experimentally provides a glfw-like API
// with desktop (via glfw) and browser (via HTML5 canvas) backends.
//
// It is used for creating a GL context and receiving events.
package glfw

// ContextSwitcher is a general mechanism for switching between contexts.
type ContextSwitcher interface {
	// MakeContextCurrent takes a context and makes it current.
	// If given context is nil, then the current context is detached.
	MakeContextCurrent(context interface{})
}

// VidMode describes a single video mode.
type VidMode struct {
	Width       int // The width, in pixels, of the video mode.
	Height      int // The height, in pixels, of the video mode.
	RedBits     int // The bit depth of the red channel of the video mode.
	GreenBits   int // The bit depth of the green channel of the video mode.
	BlueBits    int // The bit depth of the blue channel of the video mode.
	RefreshRate int // The refresh rate, in Hz, of the video mode.
}
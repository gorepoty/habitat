package composer

import "syscall/js"

// Layout
const (
	ScreenLayoutTiling   = "1"
	ScreenLayoutFloating = "2"
	ScreenLayoutMono     = "0"
)

// Screen ...
type Screen struct {
	Size     ScreenSize
	Layout   LayoutTiling
	Windows  map[string][]*Window
	Document *js.Value
}

// ScreenSize ...
type ScreenSize struct {
	Width  int
	Height int
}

// Refresh ...
func (s *Screen) Refresh() {

	//s.Layout.Windows = s.Windows[s.Layout.ID]
	s.Layout.Screen = s
	s.Layout.New()

}

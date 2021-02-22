package composer

import "syscall/js"

// Layout
const (
	ScreenLayoutTiling   = "1"
	ScreenLayoutFloating = "2"
	ScreenLayoutMono     = "0"
)

// WorkSpace  ...
type WorkSpace struct {
	Size     WorkSpaceSize
	Layout   LayoutTiling
	Windows  map[string][]*Window
	Document *js.Value
}

// WorkSpaceSize ...
type WorkSpaceSize struct {
	Width  int
	Height int
}

// Refresh ...
func (w *WorkSpace) Refresh() {

	//s.Layout.Windows = s.Windows[s.Layout.ID]
	w.Layout.WorkSpace = w
	w.Layout.New()

}

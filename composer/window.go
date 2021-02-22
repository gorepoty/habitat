package composer

import (
	"bytes"
	"embed"
	"html/template"
	"syscall/js"
)

// content holds our static web server content.
//go:embed resources/window.html
var windowBody embed.FS

// Level & Focus
const (
	WindowLayerLower       = "0"
	WindowLayerLowerFocus  = "1"
	WindowLayerNormal      = "2"
	WindowLayerNormalFocus = "3"
	WindowLayerAbove       = "4"
	WindowLayerAboveFocus  = "5"
)

// Type
const (
	WindowTypeNormal = "0"
	WindowTypeWidget = "1"
	WindowTypeDialog = "2"
	WindowTypeDock   = "3"
)

// Safety
const (
	WindowSafetyUnsafe  = "0"
	WindowSafetyUnknown = "1"
	WindowSafetyTrusted = "2"
)

// Shield
const (
	WindowShieldNormal  = "0"
	WindowShieldUnknown = "1"
	WindowShieldTrusted = "2"
)

// Layout
const (
	WindowLayoutTiling   = "1"
	WindowLayoutFloating = "2"
	WindowLayoutMono     = "0"
)

// State
const (
	WindowStateExecution = "0"
	WindowStateReady     = "1"
	WindowStateStop      = "2"
	WindowStateDead      = "3"
	WindowStateWait      = "4"
	WindowStateLoad      = "5"
)

// Window ...
type Window struct {
	ID       string
	Title    string
	Subtitle string
	Icon     string
	Src      string
	Focus    string
	Scale    string
	Document *js.Value
	Location WindowLocation
	Size     WindowSize
	Events   Event
}

// WindowSize ...
type WindowSize struct {
	Width, Height int
}

// WindowLocation ...
type WindowLocation struct {
	X, Y int
}

// New ...
func (w *Window) New(id, src string, width, height, x, y int, document *js.Value) {

	// Assignment of values
	w.ID = id
	w.Title = ""
	w.Icon = "https://img.icons8.com/fluent-systems-filled/24/000000/image-not-avialable.png"
	w.Size.Width = width
	w.Size.Height = height
	w.Location.X = x
	w.Location.Y = y
	w.Src = src
	w.Document = document

	// Print the window on the screen
	w.Print()

	// Find window on screen with id
	body := w.Document.Call("getElementById", "root"+w.ID)

	//cb := js.FuncOf(w.Close)

	w.Events.Task = js.FuncOf(w.Close)

	//w.Even["close"] = cb
	//defer cb.Release()
	//println("root" + w.ID)

	EventAdd(js.Global(), body, EventTypeClick, w.Events.Task)

}

// Body ...
func (w *Window) Body() template.HTML {

	// Reads the file via variable embed
	file, _ := windowBody.ReadFile("resources/window.html")

	// Create a body with the right values
	body, _ := template.New("window").Parse(string(file))

	buf := new(bytes.Buffer)
	body.Execute(buf, w)

	// returns a structure html
	return template.HTML(buf.String())
}

// Refresh ...
func (w *Window) Refresh() {

	document := js.Global().Get("document")
	root := document.Call("getElementById", "root"+w.ID)
	root.Call("setAttribute", "style", "top:0px; left:0px;")
}

// Print ...
func (w *Window) Print() {

	// Create a div
	body := w.Document.Call("createElement", "div")

	// Inserts the body
	body.Set("innerHTML", string(w.Body()))

	// Adds the body to the screen
	w.Document.Get("body").Call("appendChild", body)

}

// Close ...
func (w *Window) Close(this js.Value, i []js.Value) interface{} {

	// Find window on screen with id
	body := w.Document.Call("getElementById", "root"+w.ID)
	// Delete the window
	body.Set("innerHTML", "Z")

	println("root" + w.ID)

	// return 0
	return js.ValueOf(0)

}

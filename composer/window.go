package composer

import (
	"bytes"
	"embed"
	"html/template"
	"syscall/js"

	"github.com/PuerkitoBio/goquery"
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
	Layer    string
	Focus    string
	Lang     string
	Scale    string
	Location WindowLocation
	Screen   string
	Size     WindowSize
	Src      string
	Event
}

// WindowSize ...
type WindowSize struct {
	W, H string
}

// WindowLocation ...
type WindowLocation struct {
	X, Y string
}

// New ...
func (w *Window) New(id, width, height, x, y, src string) {

	doc, _ := goquery.NewDocument(src)

	// use CSS selector found with the browser inspector
	// for each, use index and item
	w.ID = id
	w.Title = doc.Find("title").Contents().Text()
	w.Icon = "https://img.icons8.com/fluent-systems-filled/24/000000/image-not-avialable.png"
	w.Size.W = width
	w.Size.H = height
	w.Location.X = x
	w.Location.Y = y
	w.Src = src

	doc.Find("link").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("rel", "") == "icon" {
			w.Icon = item.AttrOr("href", "")
		}
	})

	document := js.Global().Get("document")

	test := document.Call("createElement", "div")
	test.Set("innerHTML", string(w.Body()))

	document.Get("body").Call("appendChild", test) //2. Exposing go functions/values in javascript variables.

	canvas := document.Call("getElementById", "root0")

	cb := js.FuncOf(w.Close)

	//w.Even["close"] = cb
	//defer cb.Release()

	EventAdd(js.Global(), canvas, EventTypeClick, cb)

}

// Body ...
func (w *Window) Body() template.HTML {

	data, _ := windowBody.ReadFile("resources/window.html")

	loadwin, _ := template.New("window").Parse(string(data))

	bufwin := new(bytes.Buffer)
	loadwin.Execute(bufwin, w)

	return template.HTML(bufwin.String())
}

// Close ...
func (w *Window) Close(this js.Value, i []js.Value) interface{} {

	println("root" + w.ID)
	document := js.Global().Get("document")

	//document.Get("body").Set("innerHTML", "Z") //2. Exposing go functions/values in javascript variables.
	canvas := document.Call("getElementById", "root"+w.ID)
	canvas.Set("innerHTML", "Z")
	//document.Get("body").Call("appendChild", canvas)

	println("close")
	return js.ValueOf(0)

}

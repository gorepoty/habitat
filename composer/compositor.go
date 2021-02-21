package composer

import (
	"bytes"
	"embed"
	"html/template"
	//"syscall/js"
)

// content holds our static web server content.
//go:embed resources/base.html
var content embed.FS

// Event
const (
	CompositorEventUpdate  = 0
	CompositorEventChanged = 1
)

type desk Compositor

// Compositor ...
type Compositor struct {
	Title  string
	Window []template.HTML
	Csss   []template.CSS
	Js     []template.JS
}

// Load ...
func (c *Compositor) Load() string {

	c.Title = "WebAssembly"
	data, _ := content.ReadFile("resources/base.html")

	load, _ := template.New("webpage").Parse(string(data))

	buf := new(bytes.Buffer)

	load.Execute(buf, c)

	return buf.String()

}

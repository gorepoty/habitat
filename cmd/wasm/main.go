package main

import (
	"syscall/js"

	"github.com/gorepoty/habitat/composer"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {

	document := js.Global().Get("document")

	win1 := composer.Window{}
	win2 := composer.Window{}

	list := map[string][]*composer.Window{
		"tiling": []*composer.Window{&win1, &win2},
	}

	screen := composer.Screen{
		Size:     composer.ScreenSize{Width: 1680, Height: 800},
		Layout:   composer.LayoutTiling{ID: "tiling"},
		Windows:  list,
		Document: &document,
	}

	screen.Refresh()

	//win1.New("0", "http://localhost:8090", 1000, 500, 0, 0, document)

	//element.Call("setAttribute", "style", width+height+backgroundColor+center+font)

	<-c
}

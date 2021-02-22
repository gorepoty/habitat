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

	win1 := composer.Window{Src: "https://www.twitch.tv/"}
	win2 := composer.Window{Src: "https://www.google.it/"}
	win3 := composer.Window{Src: "https://www.youtube.com/watch?v=7VWxXkmZD_E&list=PLq0D6L5WffcpgUWRYQOq9AK6go7VnOktF&index=259"}

	workspace1 := composer.WorkSpace{
		Size:   composer.WorkSpaceSize{Width: 1920, Height: 1000},
		Layout: composer.LayoutTiling{ID: "tiling"},
		Windows: map[string][]*composer.Window{
			"tiling": []*composer.Window{&win1, &win2, &win3},
		},
		Document: &document,
	}

	workspace1.Refresh()

	//win1.New("0", "http://localhost:8090", 1000, 500, 0, 0, document)

	//element.Call("setAttribute", "style", width+height+backgroundColor+center+font)

	<-c
}

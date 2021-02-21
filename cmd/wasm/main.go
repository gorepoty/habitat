package main

import (
	"github.com/gorepoty/habitat/composer"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {

	//desk := composer.Compositor{}
	win := composer.Window{}

	win.New("0", "100", "100", "0", "0", "http://localhost:8090")

	//element.Call("setAttribute", "style", width+height+backgroundColor+center+font)

	<-c
}

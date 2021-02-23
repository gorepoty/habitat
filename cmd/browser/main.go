package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	ui.Load("http://localhost:8090/")
	ui.SetBounds(lorca.Bounds{WindowState: lorca.WindowStateFullscreen})
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

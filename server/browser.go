package server

//import "github.com/webview/webview"
import "github.com/gorepoty/habitat/webview"

// Browser ...
func Browser() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintFull)
	w.Navigate("http://localhost:8090/")
	w.Window()
	//w.Navigate("https://wasmboy.app/")

	w.Run()
}

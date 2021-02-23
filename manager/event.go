package composer

import (
	"syscall/js"
)

// Type
const (
	EventTypeClick = "click"
)

// Event ...
type Event struct {
	Task   js.Func
	Signal chan bool
}

// EventAdd ...
func EventAdd(this js.Value, ogg js.Value, event string, fun js.Func) {

	ogg.Call("addEventListener", event, fun)

}

// EventSignal ...
func EventSignal(this js.Value, ogg js.Value, event string, fun js.Func) {

	ogg.Call("addEventListener", event, fun)

}

package main

import (
	"github.com/gorepoty/habitat/server"
)

/*
func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "funziona la funzione hello\n")
}*/

func main() {

	//go web.Browser()
	server.Server()

	/*
		for true {

		}*/

}

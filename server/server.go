package server

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8090", "listen address")
	dir    = flag.String("dir", "/home/desktop/", "directory to serve")
)

// Server ...
func Server() {
	//os.Getenv("GOPATH")+
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Printf("dir on %q...", *dir)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		url := "https://www.youtube.com/watch?v=Hq0qPv6OkK0"
		fmt.Printf("HTML code of %s ...\n", url)
		resp, err := http.Get(url)
		// handle the error if there is one
		if err != nil {
			panic(err)
		}
		// do this now so it won't be forgotten
		defer resp.Body.Close()
		// reads html as a slice of bytes
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		// show the HTML code as a string %s
		fmt.Fprintf(w, string(html[:]))
	})

	http.HandleFunc("/habitat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "habitat.Load()")
	})

	log.Printf("listening on http://localhost%s", *listen)
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	http.ListenAndServe(*listen, nil)

}

package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var fs http.Handler

func main() {
	// setup, then parse flags
	var port int
	var host string

	flag.IntVar(&port, "port", 3000, "The port to bind to")
	flag.StringVar(&host, "host", "127.0.0.1", "The host to bind to")
	flag.Parse()

	addr := host + ":" + strconv.Itoa(port)

	fs = http.FileServer(http.Dir("."))
	http.HandleFunc("/", logHandler)

	log.Printf("Listening to %s on port %d...", host, port)

	http.ListenAndServe(addr, nil)
}

// Just logs the request to the console before passing it to the
// http.FileServer
func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)
	fs.ServeHTTP(w, r)
}

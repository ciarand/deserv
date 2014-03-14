package main

import (
	"flag"
	"github.com/ciarand/martini-fileindex"
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	port := parseFlags()

	m := martini.Classic()

	m.Handlers(
		martini.Logger(),
		martini.Static("."),
		fileindex.ListFiles("."),
	)

	http.ListenAndServe(":"+port, m)
}

func parseFlags() (port string) {
	// setup, then parse flags
	flag.StringVar(&port, "port", "3000", "The port to bind to")
	flag.Parse()

	return port
}

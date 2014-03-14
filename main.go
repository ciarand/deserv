package main

import (
	"flag"
	"github.com/ciarand/martini-fileindex"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"os"
)

func main() {
	port, location := parseArgs()

	server := New(location)

	server.Logger.Println("listening on 127.0.0.1:" + port)
	server.Logger.Fatal(http.ListenAndServe(":"+port, server))
}

func parseArgs() (port string, loc string) {
	// setup, then parse flags
	flag.StringVar(&port, "port", "3000", "The port to bind to")
	flag.Parse()

	loc = flag.Arg(0)
	if loc == "" {
		loc = "."
	}

	return port, loc
}

type Deserv struct {
	*martini.ClassicMartini
	Logger *log.Logger
}

func New(location string) *Deserv {
	d := &Deserv{martini.Classic(), log.New(os.Stdout, "[deserv] ", 0)}

	d.Handlers(
		martini.Logger(),
		martini.Static(location),
		fileindex.ListFiles(location),
	)

	d.Map(d.Logger)

	return d
}

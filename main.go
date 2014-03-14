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
	port := parseFlags()

	server := New()

	server.Logger.Println("listening on 127.0.0.1:" + port)
	server.Logger.Fatal(http.ListenAndServe(":"+port, server))
}

func parseFlags() (port string) {
	// setup, then parse flags
	flag.StringVar(&port, "port", "3000", "The port to bind to")
	flag.Parse()

	return port
}

type Deserv struct {
	*martini.ClassicMartini
	Logger *log.Logger
}

func New() *Deserv {
	d := &Deserv{martini.Classic(), log.New(os.Stdout, "[deserv] ", 0)}

	d.Handlers(
		martini.Logger(),
		martini.Static("."),
		fileindex.ListFiles("."),
	)

	d.Map(d.Logger)

	return d
}

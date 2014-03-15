package main

import (
	"github.com/ciarand/martini-fileindex"
	"github.com/codegangsta/martini"
	"log"
	"os"
)

// A parent structure that "extends" the Classic Martini and also provides
// access to the Logger used
type Deserv struct {
	*martini.ClassicMartini
	Logger *log.Logger
}

// Create a new "Deserv" struct and init it
func New(location string) *Deserv {
	// It starts out as just a Classic Martini and a custom logger
	d := &Deserv{martini.Classic(), log.New(os.Stdout, "[deserv] ", 0)}

	// Setup the handlers (middlewars) used in Deserv
	d.Handlers(
		// We want a logger
		martini.Logger(),
		// And to serve static files from the location
		martini.Static(location),
		// And we want a file index when the URI is a dir
		fileindex.ListFiles(location),
		// Finally if anything goes wrong we want it to recover
		martini.Recovery(),
	)

	// Map the created logger to the log.Logger type. That means that whenever
	// a function requests a log.Logger it'll use ours instead of the normal
	// Martini one
	d.Map(d.Logger)

	// Finally return the struct
	return d
}

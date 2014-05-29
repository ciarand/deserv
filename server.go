package main

import (
	"log"
	"os"

	"github.com/ciarand/martini-fileindex"
	"github.com/codegangsta/martini"
)

// A parent structure that "extends" the Classic Martini and also provides
// access to the Logger used
type Server struct {
	*martini.ClassicMartini
	Logger *log.Logger
}

// Create a new "Server" struct and init it
func NewServer(location string) *Server {
	// It starts out as just a Classic Martini and a custom logger
	d := &Server{martini.Classic(), log.New(os.Stdout, "["+NAME+"] ", 0)}

	// Setup the handlers (middlewares) used in Deserv
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

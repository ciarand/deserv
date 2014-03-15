package main

import (
	"flag"
	"github.com/ciarand/martini-fileindex"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"os"
)

// Begin the madness
func main() {
	// Get the port and location from CLI arguments
	port, location := parseArgs()

	// Create a new server instance, heavily based on the Classic Martini, but
	// with some important differences
	server := New(location)

	// Print the start message
	server.Logger.Println("listening on 127.0.0.1:" + port)
	// If the http.ListenAndServe method fails, log it and bail
	server.Logger.Fatal(http.ListenAndServe(":"+port, server))
}

// Parse the CLI arguments and grab the port and location from them
func parseArgs() (port string, loc string) {
	// setup the flags
	flag.StringVar(&port, "port", "3000", "The port to bind to")
	// Parse the args array for the flags
	flag.Parse()

	// Location is the first non-flag argument after parsing
	loc = flag.Arg(0)
	// If we didn't get a location, set it to the current dir
	if loc == "" {
		loc = "."
	}

	// Return the port and location
	return port, loc
}

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

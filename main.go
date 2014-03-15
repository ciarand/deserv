package main

import (
	"flag"
	"net/http"
)

// Begin the madness
func main() {
	// Get the port and location from CLI arguments
	runWithFlags(parseArgs())
}

func runWithFlags(flags flags) {
	// Create a new server instance, heavily based on the Classic Martini, but
	// with some important differences
	server := New(flags.loc)

	// Print the start message
	server.Logger.Println("listening on 127.0.0.1:" + flags.port)
	// If the http.ListenAndServe method fails, log it and bail
	server.Logger.Fatal(http.ListenAndServe(":"+flags.port, server))
}

// Parse the CLI arguments and grab the port and location from them
func parseArgs() flags {
	flags := flags{}

	// Parse the args array for the flags
	flags.port = *flag.String("port", "3000", "The port to bind to")
	flag.Parse()

	// Location is the first non-flag argument after parsing
	flags.loc = flag.Arg(0)
	// If we didn't get a location, set it to the current dir
	if flags.loc == "" {
		flags.loc = "."
	}

	// Return the port and location
	return flags
}

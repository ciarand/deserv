package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

// Begin the madness
func main() {
	app := cli.NewApp()

	app.Name = "deserv"
	app.Usage = "A simple development server written in Go"

	app.Flags = []cli.Flag{
		cli.IntFlag{"port, p", 3000, "The port to bind to"},
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(c *cli.Context) {
	loc := "."
	if len(c.Args()) > 0 {
		loc = c.Args()[0]
	}

	pstr := strconv.Itoa(c.Int("port"))

	server := New(loc)
	server.Logger.Println("listening on 127.0.0.1:" + pstr)
	server.Logger.Fatal(http.ListenAndServe(":"+pstr, server))
}

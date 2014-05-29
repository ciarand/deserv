package main

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/cli"
)

func NewApp() *cli.App {
	app := cli.NewApp()

	app.Name = NAME
	app.Usage = USAGE
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.IntFlag{"port, p", 3000, "The port to bind to"},
	}

	app.Action = serve

	return app
}

func serve(c *cli.Context) {
	loc := "."
	if len(c.Args()) > 0 {
		loc = c.Args()[0]
	}

	pstr := strconv.Itoa(c.Int("port"))

	server := NewServer(loc)
	server.Logger.Println("listening on 127.0.0.1:" + pstr)
	server.Logger.Fatal(http.ListenAndServe(":"+pstr, server))
}

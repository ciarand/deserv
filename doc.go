/*
Deserv provides a lightweight, local development server that serves static files
over HTTP and logs every request to STDOUT.

Package contents:

 - main.go contains the main func, which merely calls the NewApp func

 - app.go contains the application initialization and serving logic

 - server.go contains the stack of middlewares (logging, file listings, etc.)
	used within the server

 - consts.go contain the version, name, and usage string for the app
*/
package main

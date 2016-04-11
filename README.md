This project is deprecated and unmaintained. Proceed with caution!

Deserv
======
>A simple development server written in Go

![screencast][]

[screencast]: screencast.gif

Intended to provide functionality similar to the ones listed [here][gist]. It
provides pretty directory listings, but that's about it in terms of bells and
whistles.

I made it because I work with Windows users who don't necessarily have anything
cool installed on their machines. The compiled binary can be *double-clicked* to
start, and does exactly what they need.

If you need anything heavier, I recommend getting something heavier.

[gist]: https://gist.github.com/willurd/5720255

Notes
-----
1. This is not a production web server.
2. This is not a production web server.
3. This is not a production web server.

Tests
-----
There aren't any. Deal with it. You can still run `make test` if you like.

```
go test
?       github.com/ciarand/deserv   [no test files]
```

Building
--------
Should be `go get`-able.
```
go get github.com/ciarand/deserv
```

To update, run it with the `-u` flag.
```
go get -u github.com/ciarand/deserv
```

Interested in compiling for multiple platforms? I use [`goxc`][goxc]. If you
have it installed and in your path, running `make compile` should create a dump
of all the possible versions in the `snapshots` directory. If you need more fine
grained control then do it yourself, I'm not your mother.

[goxc]: https://github.com/laher/goxc

Usage
-----
```
# port is optional, defaults to 3000
# "public" is the directory to serve from - defaults to .
./deserve -port 8080 public
```

Libraries
---------
Deserv is built with [Go][]. This means that it can be run basically anywhere,
as near as I can tell.

Deserv is also built on the [Martini][] framework. If you're interested in
working on web applications in Go, I highly recommend investigating it.

Finally, the directory index middleware that was developed for this project is
available [here][middleware].

[middleware]: https://github.com/ciarand/martini-fileindex
[martini]: https://github.com/codegangsta/martini
[go]: http://golang.org/

License
-------
MIT, see the [license file][license].

[license]: /LICENSE

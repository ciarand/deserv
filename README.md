Deserv
======
>A simple development server written in Go

Intended to provide functionality similar to the ones listed [here][gist]. It
provides directory listings and that's about it in terms of bells and whistles.

I made it because I work with Windows users who don't necessarily have anything
cool installed on their machines. The compiled binary can be *double-clicked* to
start, and does exactly what they need.

[gist]: https://gist.github.com/willurd/5720255

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

Usage
-----
```
# port is optional, defaults to 3000
# "public" is the directory to serve from - defaults to .
./deserve -port 8080 public
```

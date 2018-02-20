# http2d
Really simple HTTP/2.0 server, inspired by websocketd. Also my first Go project.

# Installation
You require the net/http2 package, which can be obtained by using `go get golang.org/x/net/http2`

Then build with `go build -o http2d main.go`

# Usage
Run `./http2d [script name]` for a simple HTTP/2 server.

```
[sae@wide http2d]$ ./http2d --help
Usage of ./http2d:
  -cert string
        Path to the SSL certificate (default "server.crt")
  -http1
        Run in HTTP/1.1 mode, does not require SSL (default is false)
  -key string
        Path to the SSL key (default "server.key")
  -port int
        Port to run on (default 8080)
```

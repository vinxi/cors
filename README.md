# cors [![Build Status](https://travis-ci.org/vinxi/cors.png)](https://travis-ci.org/vinxi/cors) [![GoDoc](https://godoc.org/github.com/vinxi/cors?status.svg)](https://godoc.org/github.com/vinxi/cors) [![Coverage Status](https://coveralls.io/repos/github/vinxi/cors/badge.svg?branch=master)](https://coveralls.io/github/vinxi/cors?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/vinxi/cors)](https://goreportcard.com/report/github.com/vinxi/cors)

[W3C standard](http://www.w3.org/TR/cors/) compliant full-featured CORS support for your proxies.

## Installation

```bash
go get -u gopkg.in/vinxi/cors.v0
```

## API

See [godoc](https://godoc.org/github.com/vinxi/cors) reference.

## Examples

#### Enable CORS with custom options

```go
package main

import (
  "fmt"

  "gopkg.in/vinxi/cors.v0"
  "gopkg.in/vinxi/vinxi.v0"
)

const port = 3100

func main() {
  // Create a new vinxi proxy
  vs := vinxi.NewServer(vinxi.ServerOptions{Port: port})

  // Enable CORS support for all incoming traffic
  vs.Use(cors.New(cors.Options{
    AllowedOrigins: []string{"http://foo.com"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    AllowCredentials: true,
  }))

  // Target server to forward
  vs.Forward("http://httpbin.org")

  fmt.Printf("Server listening on port: %d\n", port)
  err := vs.Listen()
  if err != nil {
    fmt.Errorf("Error: %s\n", err)
  }
}
```

## License

MIT

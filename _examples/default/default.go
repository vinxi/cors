package main

import (
	"fmt"
	"net/http"

	"gopkg.in/vinxi/log.v0"
	"gopkg.in/vinxi/vinxi.v0"
)

const port = 3100

func main() {
	// Create a new vinxi proxy
	vs := vinxi.NewServer(vinxi.ServerOptions{Port: port})

	// Plugin multiple middlewares writting some logs
	vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		log.Infof("[%s] %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})

	vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		log.Warnf("%s", "foo bar")
		h.ServeHTTP(w, r)
	})

	// Target server to forward
	vs.Forward("http://httpbin.org")

	fmt.Printf("Server listening on port: %d\n", port)
	err := vs.Listen()
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}
}

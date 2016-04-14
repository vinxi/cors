// Package cors is net/http handler to handle CORS related requests
// as defined by http://www.w3.org/TR/cors/
package cors

import (
	"net/http"

	"github.com/rs/cors"
)

// Options is a configuration container to setup the CORS middleware.
type Options struct {
	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters
	// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
	// Only one wildcard can be used per origin.
	// Default value is ["*"]
	AllowedOrigins []string
	// AllowOriginFunc is a custom function to validate the origin. It take the origin
	// as argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowedOrigins is ignored.
	AllowOriginFunc func(origin string) bool
	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET and POST)
	AllowedMethods []string
	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [] but "Origin" is always appended to the list.
	AllowedHeaders []string
	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string
	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool
	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge int
	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool
	// Debugging flag adds additional output to debug server side CORS issues
	Debug bool
}

// Default exposes a predefined CORS handler with default options.
var Default = New(Options{})

// New creates a new CORS handler with the provided options.
func New(opts Options) func(http.Handler) http.Handler {
	// Map config options
	config := cors.Options{
		AllowedOrigins:     opts.AllowedOrigins,
		AllowOriginFunc:    opts.AllowOriginFunc,
		AllowedMethods:     opts.AllowedMethods,
		AllowCredentials:   opts.AllowCredentials,
		ExposedHeaders:     opts.ExposedHeaders,
		MaxAge:             opts.MaxAge,
		OptionsPassthrough: opts.OptionsPassthrough,
		Debug:              opts.Debug,
	}
	return cors.New(config).Handler
}

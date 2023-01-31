package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// bho
// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{
			"content-type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With", "Authorization",
		}),
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(h)
}

package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer(httpDir, port string, authenticator *auth.BasicAuth) {
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

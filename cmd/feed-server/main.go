package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aaronland/beuys"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-flags/flagset"
)

func handlerFunc() http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		rsp.Header().Set("Content-type", "application/xml")
		err := beuys.Feed(ctx, rsp)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}

func main() {

	var server_uri string

	fs := flagset.NewFlagSet("felt")
	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI string.")

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "FELT")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/", handlerFunc())

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	log.Printf("Listen for requests at %s\n", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}

}

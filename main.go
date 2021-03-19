package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/rs/cors"
)

func main() {
	var host string
	var port int
	var responseStatus int

	flag.StringVar(&host, "host", "", "HTTP server host")
	flag.IntVar(&port, "port", 3333, "HTTP server port")
	flag.IntVar(&responseStatus, "response-status", 200, "the fixed response status for the dummy HTTP server")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Listening for HTTP requests at %s\n", addr)
	http.ListenAndServe(addr, cors.AllowAll().Handler(createServerHandler(responseStatus)))
}

func createServerHandler(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+s\n\n\n", string(dump))

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.WriteHeader(status)
	}
}

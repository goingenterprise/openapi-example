package main

import (
	"log"
	"net/http"
	api "open-api-example"
	"open-api-example/internal/transporthttp"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := transporthttp.NewServer()

	r := http.NewServeMux()

	strict := api.NewStrictHandler(server, nil)
	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(strict, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

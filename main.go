package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/forecast/exp/", ExpHandler)
	mux.HandleFunc("/api/forecast/slice/", MiddleSliceHandler)
	mux.HandleFunc("/api/forecast/regression/", RegressionHandler)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":11000", handler)
}

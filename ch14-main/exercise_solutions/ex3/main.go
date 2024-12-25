package main

import (
	"net/http"

	"github.com/Prashant20nov2003/ch14/exercise_solutions/ex3/log"
)

func main() {
	server := http.Server{
		Handler: log.Middleware(http.HandlerFunc(message)),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}

func message(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Log(ctx, log.Debug, "This is a debug message")
	log.Log(ctx, log.Info, "This is an info message")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}

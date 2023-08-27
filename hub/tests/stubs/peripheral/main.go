package main

import (
	"log/slog"
	"math/rand"
	"net/http"
)

func main() {
	log := slog.Default()
	http.HandleFunc("/peripheral/temp", func(w http.ResponseWriter, r *http.Request) {
		log.Info("received request")
		stubbedTemps := []string{"21.5", "22", "21.7"}
		temp := stubbedTemps[rand.Int()%3]
		w.Write([]byte(temp))
	})
	log.Info("starting server")
	http.ListenAndServe("localhost:9999", nil)
}

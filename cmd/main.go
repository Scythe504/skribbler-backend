package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scythe504/skribbler-backend/internals/server"
	"github.com/scythe504/skribbler-backend/internals/websockets"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", server.Healthz)
	r.HandleFunc("/words", server.GetRandomWords)
	r.HandleFunc("/ws/{roomId}", websockets.HandleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", r))
}

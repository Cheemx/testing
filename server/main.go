package main

import (
	"log"
	"net/http"

	"github.com/Cheemx/testing/server/server"
)

func main() {
	store := NewInMemoryPlayerStore()
	svr := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", svr))
}

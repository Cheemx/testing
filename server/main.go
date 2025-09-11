package main

import (
	"log"
	"net/http"

	"github.com/Cheemx/testing/server/server"
)

func main() {
	svr := &server.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", svr))
}

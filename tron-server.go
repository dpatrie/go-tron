package main

import (
	"fmt"
	"github.com/dpatrie/go-tron/tron"
	"log"
	"net/http"
	"time"
)

const (
	DEFAULT_ARENA_SIZE    = 10
	DEFAULT_TURN_DURATION = 1 * time.Second
)

func main() {
	server := tron.NewServer()
	game := tron.NewGame(
		tron.NewRandomId(),
		DEFAULT_TURN_DURATION,
		tron.NewArena(DEFAULT_ARENA_SIZE),
	)

	server.AddGame(game)
	game.Arena().Join(tron.NewMoto(tron.NewRandomColor()))
	game.Arena().Join(tron.NewMoto(tron.NewRandomColor()))

	if err := game.Start(); err != nil {
		log.Println(err)
	} else {
		log.Printf("Started game %s", game.Id())
	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go-tron!")
}

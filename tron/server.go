package tron

import (
	"fmt"
)

type server struct {
	games map[Id]Game
}

func (s server) Serve() error {
	return nil
}

func (s server) AddGame(g Game) error {
	if _, found := s.games[g.Id()]; !found {
		s.games[g.Id()] = g
		return nil
	} else {
		return fmt.Errorf("Game %s already exist", g.Id())
	}
}

func (s server) Game(id Id) (Game, bool) {
	g, found := s.games[id]
	return g, found
}

func NewServer() Server {
	return server{
		games: map[Id]Game{},
	}
}

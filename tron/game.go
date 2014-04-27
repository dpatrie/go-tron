package tron

import (
	"fmt"
	"time"
)

type game struct {
	id           Id
	turnDuration time.Duration
	arena        Arena
	started      bool
	killChan     chan bool
}

func (g game) Id() Id {
	return g.id
}

func (g game) Start() error {
	if g.Arena().MotoCount() < 2 {
		return fmt.Errorf("Game %s can't start until arena has 2 or more players (count: %d)", g.Id(), g.Arena().MotoCount())
	}

	g.started = true

	go func() {
		for _ = range time.Tick(g.turnDuration) {
			select {
			case <-g.killChan:
				return
			default:
				g.arena.ProcessTurn()
			}
		}
	}()

	return nil
}

func (g game) Kill() {
	g.killChan <- true
}

func (g game) Started() bool {
	return g.started
}

func (g game) Finished() bool {
	return g.arena.Winner() != nil
}

func (g game) Arena() Arena {
	return g.arena
}

func NewGame(id Id, turnDuration time.Duration, a Arena) Game {
	return game{
		id:           id,
		arena:        a,
		turnDuration: turnDuration,
		killChan:     make(chan bool),
	}
}

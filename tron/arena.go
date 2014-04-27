package tron

import (
	"fmt"
)

type motoPosition map[Color]Square

type arena struct {
	moto    map[Color]Moto
	pos     motoPosition
	squares [][]Square
	winner  Moto
}

//Loop on each moto. Determine where they are going next
// If two moto collide, crash them both
// If the next position is colored, crash
// Else, color it
//If a moto crashed, remove it's color line
func (a arena) ProcessTurn() error {
	fmt.Println("Processing turn!")
	// for c, m := range a.moto {

	// }

	return nil
}

func (a arena) Join(m Moto) error {
	if _, found := a.moto[m.Color()]; !found {
		a.moto[m.Color()] = m
		return nil
	} else {
		return fmt.Errorf("Moto %s already exist in this arena", m.Color())
	}
}

func (a arena) MotoCount() int {
	return len(a.moto)
}

func (a arena) Winner() Moto {
	return a.winner
}

func NewArena(size int) Arena {
	a := arena{
		moto:    map[Color]Moto{},
		pos:     motoPosition{},
		squares: make([][]Square, size),
	}

	for x := 0; x < size; x++ {
		a.squares[x] = make([]Square, size)
		for y := 0; y < size; y++ {
			a.squares[x][y] = newSquare(x, y)
		}
	}
	return a
}

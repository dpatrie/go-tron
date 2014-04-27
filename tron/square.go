package tron

import (
	"fmt"
)

type square struct {
	x, y  int
	color Color
}

func (s square) X() int {
	return s.x
}

func (s square) Y() int {
	return s.y
}

func (s square) Color() Color {
	return s.color
}

func (s square) Colorize(c Color) error {
	if s.color == "" {
		s.color = c
		return nil
	} else {
		return fmt.Errorf("Square [%d,%d] is already colored in %s. Unable to color it in %s.", s.x, s.y, s.color, c)
	}
}

func newSquare(x, y int) Square {
	return square{
		x: x,
		y: y,
	}
}

package tron

import (
	"fmt"
	"image/color/palette"
	"math/rand"
	"time"
)

type Color string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRandomColor() Color {
	c := palette.WebSafe[rand.Intn(len(palette.WebSafe)-1)]
	r, g, b, _ := c.RGBA()
	return Color(fmt.Sprintf("%02X%02X%02X", uint8(r>>8), uint8(g>>8), uint8(b>>8)))
}

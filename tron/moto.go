package tron

type direction string

var (
	DirectionLeft     = direction("left")
	DirectionRight    = direction("right")
	DirectionStraight = direction("straight")
)

type moto struct {
	color   Color
	dir     direction
	crashed bool
}

func (m moto) Color() Color {
	return m.color
}

func (m moto) Direction() direction {
	return m.dir
}

func (m moto) Steer(dir direction) {
	m.dir = dir
}

func (m moto) Crashed() bool {
	return m.crashed
}

func (m moto) Crash() {
	m.crashed = true
}

func NewMoto(c Color) Moto {
	return moto{
		color: c,
		dir:   DirectionStraight,
	}
}

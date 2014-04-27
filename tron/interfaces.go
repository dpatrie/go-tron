package tron

type Server interface {
	AddGame(g Game) error
	Game(id Id) (Game, error)
	Serve()
}

type Game interface {
	Id() Id
	Start() error
	Kill()
	Started() bool
	Finished() bool
	Arena() Arena
}

type Moto interface {
	Color() Color
	Direction() direction
	Steer(d direction)
	Crash()
	Crashed() bool
}

type Arena interface {
	ProcessTurn() error
	Join(m Moto) error
	MotoCount() int
	Winner() Moto
}

type Square interface {
	X() int
	Y() int
	Colorize(c Color) error
	Color() Color
}

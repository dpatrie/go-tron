package tron

import (
	"github.com/nu7hatch/gouuid"
)

type Id string

func NewRandomId() Id {
	id, _ := uuid.NewV4()
	return Id(id.String())
}

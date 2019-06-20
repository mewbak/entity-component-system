package components

import "github.com/davidsbond/game/internal/component"

const (
	TypeWASDControl = component.Type("wasd-control")
)

type (
	WASDControl struct {
		Position *Position
		Speed    int
	}
)

func NewWASDControl(p *Position, speed int) *WASDControl {
	return &WASDControl{
		Position: p,
		Speed:    speed,
	}
}

func (wasd *WASDControl) GetType() component.Type {
	return TypeWASDControl
}

package components

import "github.com/davidsbond/game/internal/component"

const (
	TypeWASDControl = component.Type("wasd-control")
)

type (
	WASDControl struct {
		Position *Position
		Speed    float64
	}
)

func NewWASDControl(p *Position, speed float64) *WASDControl {
	return &WASDControl{
		Position: p,
		Speed:    speed,
	}
}

func (wasd *WASDControl) GetType() component.Type {
	return TypeWASDControl
}

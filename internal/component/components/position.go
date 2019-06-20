package components

import "github.com/davidsbond/game/internal/component"

const (
	TypePosition = component.Type("position")
)

type (
	Position struct {
		X int
		Y int
	}
)

func NewPosition() *Position {
	return &Position{}
}

func (p *Position) GetType() component.Type {
	return TypePosition
}

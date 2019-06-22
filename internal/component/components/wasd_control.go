package components

import "github.com/davidsbond/game/internal/component"

const (
	// TypeWASDControl is the unique type for the WASDControl
	// component.
	TypeWASDControl = component.Type("wasd-control")
)

type (
	// The WASDControl type is used on entities to allow them to
	// be moved using the W, A, S & D keys on the keyboard. It tracks
	// the current position and speed at which the entity should move.
	WASDControl struct {
		Position *Position
		Speed    float64
	}
)

// NewWASDControl creates a new WASDControl component using the provided
// position and speed value.
func NewWASDControl(p *Position, speed float64) *WASDControl {
	return &WASDControl{
		Position: p,
		Speed:    speed,
	}
}

// GetType returns TypeWASDControl
func (wasd *WASDControl) GetType() component.Type {
	return TypeWASDControl
}

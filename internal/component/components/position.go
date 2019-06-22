package components

import "github.com/davidsbond/game/internal/component"

const (
	// TypePosition is the unique type for the Position component.
	TypePosition = component.Type("position")
)

type (
	// The Position component denotes an entity's position on the
	// screen
	Position struct {
		// Location on the x-axis
		X float64

		// Location on the y-axis.
		Y float64
	}
)

// NewPosition creates a new position at 0:0
func NewPosition() *Position {
	return &Position{}
}

// GetType returns TypePosition
func (p *Position) GetType() component.Type {
	return TypePosition
}

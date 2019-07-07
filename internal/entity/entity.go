// Package entity contains types for managing entities within the game.
package entity

import (
	"github.com/davidsbond/game/internal/component"
)

type (
	// The Entity type contains methods for updating and rendering an entity within
	// the game.
	Entity struct {
		components []component.Component
	}
)

// New creates a new blank entity
func New() *Entity {
	return &Entity{}
}

// AddComponent adds a given component to the entity
func (e *Entity) AddComponent(cmp component.Component) {
	e.components = append(e.components, cmp)
}

// GetComponentsOfType returns all components registered with the entity that match
// a given type.
func (e *Entity) GetComponentsOfType(t component.Type) (out []component.Component) {
	for _, cmp := range e.components {
		if cmp.GetType() == t {
			out = append(out, cmp)
		}
	}

	return
}

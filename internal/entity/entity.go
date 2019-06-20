// Package entity contains types for managing entities within the game.
package entity

import (
	"fmt"

	"github.com/davidsbond/game/internal/component"
)

type (
	// The Entity type contains methods for updating and rendering an entity within
	// the game.
	Entity struct {
		id         string
		components map[string]component.Component
	}
)

func New(id string) *Entity {
	return &Entity{
		id:         id,
		components: make(map[string]component.Component),
	}
}

func (e *Entity) AddComponent(name string, cmp component.Component) error {
	if cmp, ok := e.components[name]; ok {
		return fmt.Errorf("entity %s already has a %T component named %s", e.id, cmp, name)
	}

	e.components[name] = cmp
	return nil
}

func (e *Entity) GetComponentsOfType(t component.Type) (out []component.Component) {
	for _, cmp := range e.components {
		if cmp.GetType() == t {
			out = append(out, cmp)
		}
	}

	return
}

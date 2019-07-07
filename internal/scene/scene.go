// Package scene contains types for managing scenes within the game.
package scene

import (
	"github.com/davidsbond/game/internal/entity"
)

type (
	// The Scene type represents a single in-game space. Each scene has
	// a name, dimensions (which are used to confine the player) and
	// a number of entities that relevant systems will interact with
	Scene struct {
		name     string
		entities []*entity.Entity
	}
)

// New creates a new Scene with the give name, width and
// height.
func New(name string, w, h float64) *Scene {
	return &Scene{
		name: name,
	}
}

// AddEntity adds an entity to the scene.
func (s *Scene) AddEntity(entity *entity.Entity) {
	s.entities = append(s.entities, entity)
}

// GetEntities returns all entities within the scene
func (s *Scene) GetEntities() []*entity.Entity {
	return s.entities
}

// Package scene contains types for managing scenes within the game.
package scene

import (
	"github.com/davidsbond/game/internal/entity"
)

type (
	Scene struct {
		name     string
		width    float64
		height   float64
		entities []*entity.Entity
	}
)

func New(name string, w, h float64) *Scene {
	return &Scene{
		name:   name,
		width:  w,
		height: h,
	}
}

func (s *Scene) AddEntity(entity *entity.Entity) {
	s.entities = append(s.entities, entity)
}

func (s *Scene) GetEntities() []*entity.Entity {
	return s.entities
}

func (s *Scene) GetBoundaries() (minX, minY, maxX, maxY float64) {
	maxX = s.width
	maxY = s.height

	return
}

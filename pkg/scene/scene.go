// Package scene contains types for managing scenes within the game.
package scene

import (
	"fmt"

	"github.com/davidsbond/game/internal/entity"
)

type (
	Scene struct {
		name     string
		width    int
		height   int
		entities map[string]entity.Entity
	}
)

func New(name string, w, h int) *Scene {
	return &Scene{
		name:     name,
		width:    w,
		height:   h,
		entities: make(map[string]entity.Entity),
	}
}

func (s *Scene) AddEntity(name string, entity entity.Entity) error {
	if e, ok := s.entities[name]; ok {
		return fmt.Errorf("scene %s already has a %T entity named %s", s.name, e, name)
	}

	s.entities[name] = entity
	return nil
}

func (s *Scene) GetEntities() (out []entity.Entity) {
	for _, e := range s.entities {
		out = append(out, e)
	}

	return
}

func (s *Scene) GetBoundaries() (minX, minY, maxX, maxY int) {
	maxX = s.width
	maxY = s.height

	return
}

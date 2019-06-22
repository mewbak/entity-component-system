package entities

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/entity"
)

// NewPlayer creates a new player-controlled entity that is represented
// by a cartoon of a sleeping cat.
func NewPlayer() (*entity.Entity, error) {
	pl := entity.New()

	pos := components.NewPosition()
	wasd := components.NewWASDControl(pos, 1)
	spr, err := components.NewSprite("assets/gopher.png", pos)

	if err != nil {
		return nil, err
	}

	pl.AddComponent(pos)
	pl.AddComponent(wasd)
	pl.AddComponent(spr)

	return pl, nil
}

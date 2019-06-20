package entities

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/entity"
	"github.com/hashicorp/go-multierror"
)

func NewPlayer() (*entity.Entity, error) {
	pl := entity.New("player")

	pos := components.NewPosition()
	wasd := components.NewWASDControl(pos, 1)
	spr, err := components.NewSprite("test.png", pos)

	if err != nil {
		return nil, err
	}

	return pl, multierror.Append(
		pl.AddComponent("position", pos),
		pl.AddComponent("wasd-control", wasd),
		pl.AddComponent("sprite", spr),
	).ErrorOrNil()
}

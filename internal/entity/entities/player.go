package entities

import (
	"image/color"

	"github.com/davidsbond/game/internal/controller"
	"github.com/davidsbond/game/internal/controller/controllers"
	"github.com/davidsbond/game/internal/entity"
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/hajimehoshi/ebiten"
)

type (
	player struct {
		x int
		y int

		movement controller.Controller
	}
)

func NewPlayer() entity.Entity {
	pl := &player{}
	pl.movement = controllers.NewMovementController(&pl.x, &pl.y, 1)

	return pl
}

func (p *player) Draw(screen *ebiten.Image) error {
	screen.Set(p.x, p.y, color.White)
	return nil
}

func (p *player) Update(state *input.State, info *scene.Info) error {
	if err := p.movement.Update(state, info); err != nil {
		return err
	}

	return nil
}

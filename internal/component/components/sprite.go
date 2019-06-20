package components

import (
	"github.com/davidsbond/game/internal/component"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	TypeSprite = component.Type("sprite")
)

type (
	Sprite struct {
		Image    *ebiten.Image
		Position *Position
	}
)

func NewSprite(filename string, pos *Position) (*Sprite, error) {
	img, _, err := ebitenutil.NewImageFromFile(filename, ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	return &Sprite{
		Image:    img,
		Position: pos,
	}, nil
}

func (s *Sprite) GetType() component.Type {
	return TypeSprite
}

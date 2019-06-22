package components

import (
	"github.com/davidsbond/game/internal/component"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	// TypeSprite is the unique type for the Sprite component.
	TypeSprite = component.Type("sprite")
)

type (
	// The Sprite component represents an image at a certain position
	// within the game screen.
	Sprite struct {
		Image    *ebiten.Image
		Position *Position
	}
)

// NewSprite creates a new sprite. The 'filename' parameter will be the image
// that is loaded from disk and eventually rendered onto the screen. The 'pos'
// parameter determines where on the screen the sprite will be rendered.
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

// GetType returns TypeSprite
func (s *Sprite) GetType() component.Type {
	return TypeSprite
}

package systems

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/system"
	"github.com/davidsbond/game/pkg/input"
	"github.com/davidsbond/game/pkg/scene"
	"github.com/hajimehoshi/ebiten"
	"github.com/hashicorp/go-multierror"
)

type (
	spriteRenderer struct{}
)

func NewSpriteRenderer() system.System {
	return &spriteRenderer{}
}

func (r *spriteRenderer) Run(screen *ebiten.Image, _ *input.State, scene *scene.Scene) error {
	var errs []error

	for _, e := range scene.GetEntities() {
		cmps := e.GetComponentsOfType(components.TypeSprite)

		for _, cmp := range cmps {
			sprite := cmp.(*components.Sprite)

			if err := r.renderSprite(screen, sprite); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return multierror.Append(nil, errs...).ErrorOrNil()
}

func (r *spriteRenderer) renderSprite(screen *ebiten.Image, sprite *components.Sprite) error {
	return screen.DrawImage(sprite.Image, &ebiten.DrawImageOptions{
		GeoM: ebiten.TranslateGeo(sprite.Position.X, sprite.Position.Y),
	})
}

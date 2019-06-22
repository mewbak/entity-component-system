package systems

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/pkg/input"
	"github.com/hajimehoshi/ebiten"
	"github.com/hashicorp/go-multierror"
)

// The SpriteRenderer system finds components of type components.TypeSprite and renders them
// onto the screen at their specified position.
func SpriteRenderer(screen *ebiten.Image, _ *input.State, scene *scene.Scene) error {
	var errs []error

	for _, e := range scene.GetEntities() {
		cmps := e.GetComponentsOfType(components.TypeSprite)

		for _, cmp := range cmps {
			sprite := cmp.(*components.Sprite)

			err := screen.DrawImage(sprite.Image, &ebiten.DrawImageOptions{
				GeoM: ebiten.TranslateGeo(sprite.Position.X, sprite.Position.Y),
			})

			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return multierror.Append(nil, errs...).ErrorOrNil()
}

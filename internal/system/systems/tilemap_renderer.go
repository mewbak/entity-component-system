package systems

import (
	"image"

	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/pkg/input"
	"github.com/hajimehoshi/ebiten"
	"github.com/hashicorp/go-multierror"
)

// TilemapRenderer obtains the current tile map from the scene and renders it. This should be ran
// before all other rendering operations to avoid sprites being rendered behind the scene's tile map.
func TilemapRenderer(screen *ebiten.Image, _ *input.State, scene *scene.Scene) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	xNum := scene.Width / scene.Tilemap.TileSize

	var errs []error
	for _, l := range scene.Tilemap.Layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64((i%xNum)*scene.Tilemap.TileSize),
				float64((i/xNum)*scene.Tilemap.TileSize),
			)

			sx := (t % scene.Tilemap.NumX) * scene.Tilemap.TileSize
			sy := (t / scene.Tilemap.NumX) * scene.Tilemap.TileSize

			if err := screen.DrawImage(
				scene.Tilemap.Image.SubImage(
					image.Rect(sx, sy, sx+scene.Tilemap.TileSize, sy+scene.Tilemap.TileSize)).(*ebiten.Image),
				op); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return multierror.Append(nil, errs...).ErrorOrNil()
}

// Package tile contains methods/types for creating tile maps.
package tile

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type (
	// The Map type represents a tile map
	Map struct {
		// The image to use to render tiles
		Image *ebiten.Image

		// The layers of tiles in the tilemap, each entry in the array is another
		// layer. Layers are rendered starting from the 0th layer.
		Layers []Layer

		// The tilemap width
		Width int

		// The tilemap height
		Height int

		// The size of the tiles, for example: 32, 64 etc.
		TileSize int

		// The number of tiles on the X-axis, once reached, the Y value
		// will be incremented to draw further tiles on the next row.
		NumX int
	}

	// The MapOptions type contains the options to pass into tiles.NewMap when creating
	// a new map. Their definitions can be found on the Map type.
	MapOptions struct {
		Width    int
		Height   int
		TileSize int
		NumX     int
	}

	Layer []int
)

// NewMap creates a new tile map. It loads the background-image from the provided 'src'
// parameter and may return an error if image loading fails.
func NewMap(src string, opts MapOptions) (*Map, error) {
	img, _, err := ebitenutil.NewImageFromFile(src, ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	return &Map{
		Image:    img,
		Width:    opts.Width,
		Height:   opts.Height,
		TileSize: opts.TileSize,
		NumX:     opts.NumX,
	}, nil
}

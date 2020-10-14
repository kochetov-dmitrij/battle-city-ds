package game

import (
	"github.com/faiface/pixel"
)

const (
	tileEmpty = '.'
	tileBrick = '#'
	tileSteel = '@'
	tileGrass = '%'
	tileWater = '~'
	tileFroze = '-'
)

type tile struct {
	sprite *pixel.Sprite
	x      int64
	y      int64
}

func (t *tile) draw(target pixel.Target) {
	t.sprite.Draw(target, pixel.IM.Moved(pixel.V(float64(t.x), float64(t.y))))
}

func tilePixelPosition(x int64, y int64) (int64, int64) {
	y = (26-y)*8 - 4
	x = x*8 + 4
	return x, y
}

func tileWorldMapByPixel(x int64, y int64) (int64, int64) {
	x = (x - 4) / 8
	y = 26 - (y+4)/8
	return x, y
}

func (g *game) draw() error {
	for i := 0; i < len(g.world.worldMap); i++ {
		for j := 0; j < len(g.world.worldMap[i]); j++ {
			if g.world.worldMap[i][j] == tileEmpty {
				continue
			}
			x, y := tilePixelPosition(int64(j), int64(i))
			t := tile{
				sprite: g.sprites.tiles[g.world.worldMap[i][j]],
				x:      x,
				y:      y,
			}
			t.draw(g.canvas)
		}
	}
	return nil
}

func (g *game) drawLevel(level [26][26]byte) {
	g.canvas.Draw(g.window, pixel.IM)
}

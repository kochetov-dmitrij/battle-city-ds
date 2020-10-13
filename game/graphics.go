package game

import "github.com/faiface/pixel"

const (
	tileEmpty = '.'
	tileBrick = '#'
	tileSteel = '@'
	tileGrass = '%'
	tileWater = '~'
	tileFroze = '-'
)

func tilePositionMat(x int, y int) {
	x = (26 - x) * 16
	y = y * 16

}

func (g *game) drawLevel(level [26][26]byte) {
	g.canvas.Draw(g.window, pixel.IM)
}

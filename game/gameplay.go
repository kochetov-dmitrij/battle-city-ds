package game

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/faiface/pixel"
)

type Direction int

const (
	pixelsPerSecond           = 40
	up              Direction = iota
	down
	left
	right
)

type world struct {
	tanks    []*tank
	worldMap [26][26]byte
}

type tank struct {
	direction Direction
	sprite    pixel.Sprite
	x, y      int64
	size      [2]int64
}

func (g *game) loadTank(sprite *pixel.Sprite, changeColor bool) (t *tank) {
	if changeColor {
		// TODO add coloring
	}
	t = &tank{
		direction: up,
		sprite:    *sprite,
		x:         rand.Int63n(20),
		y:         rand.Int63n(20),
		size: [2]int64{int64(sprite.Frame().Max.X-sprite.Frame().Min.X) / 2,
			int64(sprite.Frame().Max.Y-sprite.Frame().Min.Y) / 2},
	}
	return t
}

func (t *tank) draw(target pixel.Target) {
	// fmt.Println(t.x, t.y)
	mat := pixel.IM
	switch t.direction {
	case left:
		mat = mat.Rotated(pixel.ZV, math.Pi/2)
	case down:
		mat = mat.Rotated(pixel.ZV, math.Pi)
	case right:
		mat = mat.Rotated(pixel.ZV, 3*math.Pi/2)
	}
	mat = mat.Moved(pixel.V(float64(t.x), float64(t.y)))

	t.sprite.Draw(target, mat)
}

func (t *tank) canMove(direction Direction, movedPixels int64) bool {
	if direction == right {
		fmt.Printf("(%d+%d = %d) < %d\n", t.x, movedPixels, t.x+movedPixels, gameH-t.size[0])
		return t.x+movedPixels < gameH-t.size[0]
	}
	if direction == left {
		return t.x-movedPixels > t.size[0]
	}
	if direction == up {
		fmt.Printf("(%d+%d = %d) < %d\n", t.y, movedPixels, t.y+movedPixels, gameH-t.size[0])
		return t.y+movedPixels < gameH-t.size[1]
	}
	if direction == down {
		return t.y-movedPixels > t.size[1]
	}
	return false
}

func (t *tank) getNewPos(direction Direction) (int64, int64) {
	movedPixels := int64(1)
	if !t.canMove(direction, movedPixels) {
		return t.x, t.y
	}
	if direction == right {
		return t.x + movedPixels, t.y
	}
	if direction == left {
		return t.x - movedPixels, t.y
	}
	if direction == up {
		return t.x, t.y + movedPixels
	}
	if direction == down {
		return t.x, t.y - movedPixels
	}
	return t.x, t.y
}

func (t *tank) update(time int64, direction Direction) {
	t.direction = direction
	t.x, t.y = t.getNewPos(direction)
}

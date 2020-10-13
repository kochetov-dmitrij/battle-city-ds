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
}

func loadTank(sprite *pixel.Sprite, changeColor bool) (t *tank) {
	if changeColor {
		// TODO add coloring
	}
	t = &tank{
		direction: up,
		sprite:    *sprite,
		x:         rand.Int63n(400),
		y:         rand.Int63n(350),
	}
	return t
}

func (t *tank) draw(target pixel.Target) {
	fmt.Println(t.x, t.y)
	mat := pixel.IM
	switch t.direction {
	case left:
		mat = mat.Rotated(pixel.ZV, math.Pi/2)
	case down:
		mat = mat.Rotated(pixel.ZV, math.Pi)
	case right:
		mat = mat.Rotated(pixel.ZV, 3*math.Pi/2)
	}
	mat = mat.Moved(pixel.V(float64(t.x)/10, float64(t.y)/10))

	t.sprite.Draw(target, mat)
}

func (t *tank) getNewPos(time int64, direction Direction) (int64, int64) {
	movedPixels := int64(20)
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
	t.x, t.y = t.getNewPos(time, direction)
}

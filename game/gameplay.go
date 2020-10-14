package game

import (
	"math"

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
	size := [2]int64{int64(sprite.Frame().Max.X-sprite.Frame().Min.X) / 2,
		int64(sprite.Frame().Max.Y-sprite.Frame().Min.Y) / 2}
	t = &tank{
		direction: up,
		sprite:    *sprite,
		x:         size[0], // + rand.Int63n(20),
		y:         size[1], // + rand.Int63n(20),
		size:      size,
	}
	return t
}

func (t *tank) draw(target pixel.Target) {
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

func checkBlockingTile(g *game, position [2]int64, size [2]int64, direction Direction) bool {
	checkXLeft, checkYBottom := position[0]-size[0], position[1]-size[1]
	checkXRight, checkYTop := position[0]+size[0], position[1]+size[1]

	closestTilesCenters := [12]int64{
		checkXLeft - (checkXLeft % 4),     // leftCenterX
		checkXLeft + 4 - (checkXLeft % 4), // leftCenterX
		checkXRight - (checkXRight % 4),   // rightCenterX
		checkXRight + 4 - (checkXRight % 4),
		position[0] - (position[0] % 4),       // rightCenterX
		position[0] + 4 - (position[0] % 4),   // rightCenterX
		checkYBottom - (checkYBottom % 4),     // bottomCenterY
		checkYBottom + 4 - (checkYBottom % 4), // topCenterY
		checkYTop - (checkYTop % 4),           // bottomCenterY
		checkYTop + 4 - (checkYTop % 4),       // topCenterY
		position[1] - (position[1] % 4),       // rightCenterX
		position[1] + 4 - (position[1] % 4),   // rightCenterX
	}

	blocking := false
	for i := 0; i < 6; i++ {
		for j := 6; j < 12; j++ {
			blocking = false
			x := closestTilesCenters[i]
			y := closestTilesCenters[j]
			yMap, xMap := tileWorldMapByPixel(x, y)
			if (xMap == 26) || (yMap == 26) {
				continue
			}
			switch g.world.worldMap[xMap][yMap] {
			case tileBrick:
				blocking = true
			case tileSteel:
				blocking = true
			case tileWater:
				blocking = true
			}
			if !blocking {
				continue
			}
			rightIntersection := (position[0] <= x && ((position[0] + size[0]) >= (x - 3)))
			leftIntersection := (position[0] >= x && ((position[0])-size[0]+1) < (x+3))
			insideHorizontal := checkXLeft <= x && checkXRight >= (x-2)
			insideVertical := checkYBottom <= y && checkYTop >= (y-2)
			topIntersection := (position[1] <= y && ((position[1] + size[1]) >= (y - 3)))
			bottomIntersection := (position[1] >= y && ((position[1])-size[1]+1) < (y+3))

			if (direction == left) && leftIntersection && insideVertical {
				return false
			}
			if (direction == right) && rightIntersection && insideVertical {
				return false
			}
			if (direction == up) && topIntersection && insideHorizontal {
				return false
			}
			if (direction == down) && bottomIntersection && insideHorizontal {
				return false
			}
		}
	}
	return true
}

func (t *tank) canMove(g *game, direction Direction, movedPixels int64) bool {
	borderMoveAllowed := false
	newX, newY := t.x, t.y
	if direction == right {
		newX = t.x + movedPixels
		borderMoveAllowed = newX < gameH-t.size[0]
	}
	if direction == left {
		newX = t.x - movedPixels
		borderMoveAllowed = newX > t.size[0]
	}
	if direction == up {
		newY = t.y + movedPixels
		borderMoveAllowed = newY < gameH-t.size[1]
	}
	if direction == down {
		newY = t.y - movedPixels
		borderMoveAllowed = newY > t.size[1]
	}
	newPosition := [2]int64{newX, newY}
	return borderMoveAllowed && checkBlockingTile(g, newPosition, t.size, direction)
}

func (t *tank) getNewPos(direction Direction, movedPixels int64) (int64, int64) {
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

func (g *game) updateTank(t *tank, direction Direction, moves bool) {
	movedPixels := int64(1)
	t.direction = direction
	if moves && t.canMove(g, direction, movedPixels) {
		t.x, t.y = t.getNewPos(direction, movedPixels)
	}
	t.draw(g.canvas)
}

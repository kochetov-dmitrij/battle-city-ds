package game

import (
	"math"

	"github.com/faiface/pixel"
)

const (
	bulletSizeX = int64(1)
	bulletSizeY = int64(2)
)

type bullet struct {
	sprite    pixel.Sprite
	direction Direction
	x, y      int64
	size      [2]int64
	state     State
}

func (g *game) loadBullet(x int64, y int64, direction Direction, state State) *bullet {
	sprite := g.sprites.bullet
	b := &bullet{
		sprite:    *sprite,
		direction: direction,
		x:         x,
		y:         y,
		size:      [2]int64{bulletSizeX, bulletSizeY},
		state:     state,
	}
	return b
}

func (t *tank) fire(g *game) {
	if t.bullet != nil {
		return
	}
	x, y := t.x, t.y
	switch t.direction {
	case right:
		x += 5
	case left:
		x -= 5
	case up:
		y += 5
	case down:
		y -= 5
	}
	t.bullet = g.loadBullet(x, y, t.direction, active)
}

func (b *bullet) draw(target pixel.Target) {
	mat := pixel.IM
	switch b.direction {
	case left:
		mat = mat.Rotated(pixel.ZV, math.Pi/2)
	case down:
		mat = mat.Rotated(pixel.ZV, math.Pi)
	case right:
		mat = mat.Rotated(pixel.ZV, 3*math.Pi/2)
	}
	mat = mat.Moved(pixel.V(float64(b.x), float64(b.y)))
	b.sprite.Draw(target, mat)
}

func (b *bullet) checkTankDestroyed(g *game, playerTank *tank) bool {
	bulletSpriteRect := b.sprite.Frame()
	bulletV := pixel.V(float64(b.x), float64(b.y)).Sub(bulletSpriteRect.Min)
	bulletRect := bulletSpriteRect.Moved(bulletV)

	for _, player := range g.players {
		if player == nil {
			continue
		}
		t := player.tank
		if t == playerTank {
			continue
		}
		if t.bullet != nil {
			anotherBulletV := pixel.V(float64(t.bullet.x), float64(t.bullet.y)).Sub(bulletSpriteRect.Min)
			anotherBulletRect := bulletSpriteRect.Moved(anotherBulletV)
			if bulletRect.Intersects(anotherBulletRect) {
				playerTank.bullet = nil
				t.bullet = nil
				return true
			}
		}

		tankRect := t.sprite.Frame()
		tankV := pixel.V(float64(t.x), float64(t.y)).Sub(tankRect.Min)
		tankRect = tankRect.Moved(tankV)

		if !bulletRect.Intersects(tankRect) && !rectContains(tankRect, bulletRect) {
			continue
		}
		if t.state != active {
			continue
		}

		b.state = removed
		t.state = explodingS + 1
		t.bullet = nil
		g.incrementScore(g.players[playerTank.number])
		return true
	}
	return false
}

func (b *bullet) checkBlockingTile(g *game) {
	checkXLeft, checkYBottom := b.x-b.size[0], b.y-b.size[1]
	checkXRight, checkYTop := b.x-b.size[0], b.y+b.size[1]

	closestTilesCenters := [12]int64{
		checkXLeft - (checkXLeft % 4),     // leftCenterX
		checkXLeft + 4 - (checkXLeft % 4), // leftCenterX
		checkXRight - (checkXRight % 4),   // rightCenterX
		checkXRight + 4 - (checkXRight % 4),
		b.x - (b.x % 4),                       // rightCenterX
		b.x + 4 - (b.x % 4),                   // rightCenterX
		checkYBottom - (checkYBottom % 4),     // bottomCenterY
		checkYBottom + 4 - (checkYBottom % 4), // topCenterY
		checkYTop - (checkYTop % 4),           // bottomCenterY
		checkYTop + 4 - (checkYTop % 4),
		b.y - (b.y % 4),
		b.y + 4 - (b.y % 4),
	}

	bulletRect := b.sprite.Frame()
	bulletV := pixel.V(float64(b.x), float64(b.y)).Sub(bulletRect.Min)
	bulletRect = bulletRect.Moved(bulletV)

	for i := 0; i < 6; i++ {
		for j := 6; j < 12; j++ {
			x := closestTilesCenters[i]
			y := closestTilesCenters[j]
			yMap, xMap := tileWorldMapByPixel(x, y)
			if (xMap == 26) || (yMap == 26) ||
				(xMap == -1) || (yMap == -1) ||
				g.world.worldMap[xMap][yMap] == tileEmpty {
				continue
			}
			tileRect := g.sprites.tiles[g.world.worldMap[xMap][yMap]].Frame()
			tileV := pixel.V(float64(x), float64(y)).Sub(tileRect.Min)

			tileRect = tileRect.Moved(tileV)

			switch g.world.worldMap[xMap][yMap] {
			case tileWater:
				continue
			case tileGrass:
				continue
			}

			if !bulletRect.Intersects(tileRect) && !rectContains(tileRect, bulletRect) {
				continue
			}
			b.state = explodingS
			if g.world.worldMap[xMap][yMap] == tileSteel {
				continue
			}
			g.world.worldMap[xMap][yMap] = tileEmpty
		}
	}
}

func (b *bullet) moveBullet(g *game, t *tank) {
	movedPixels := int64(8)
	if b.direction == right {
		if b.x+movedPixels >= gameH {
			b.x = gameH - 1
			b.state = explodingS
			return
		}
		b.x = b.x + movedPixels
	}
	if b.direction == left {
		if b.x-movedPixels <= 0 {
			b.x = 0
			b.state = explodingS
			return
		}
		b.x = b.x - movedPixels
	}
	if b.direction == up {
		if b.y+movedPixels >= gameH {
			b.y = gameH - 1
			b.state = explodingS
			return
		}
		b.y = b.y + movedPixels
	}
	if b.direction == down {
		if b.y-movedPixels <= 0 {
			b.y = 0
			b.state = explodingS
			return
		}
		b.y = b.y - movedPixels
	}
	if b.checkTankDestroyed(g, t) {
		return
	}
	b.checkBlockingTile(g)
}

func (t *tank) updateBullet(g *game) {
	b := t.bullet
	if b.state == active {
		b.moveBullet(g, t)
	}

	if explodingS <= b.state && b.state <= explodingB {
		b.checkBlockingTile(g)
	}

	if b.state == removed {
		b.checkTankDestroyed(g, t)
	}

	switch b.state {
	case explodingS:
		b.state = explodingM
		b.sprite = *g.sprites.explosions[0]
	case explodingM:
		t.bullet = nil
		b.sprite = *g.sprites.explosions[1]
	case removed:
		t.bullet = nil
	}

	b.draw(g.canvas)
}

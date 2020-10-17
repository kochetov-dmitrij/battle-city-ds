package game

import (
	"fmt"
	"image/color"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type score struct {
	atlas  *text.Atlas
	canvas *pixelgl.Canvas
}

func (g *game) initScore(assetsPath string) *score {
	scoreBounds := pixel.Rect{Max: pixel.Vec{X: 2 * (gameW - gameH), Y: 4 * gameH}}
	scoreCanvas := pixelgl.NewCanvas(scoreBounds)
	// atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	face, err := loadTTF(filepath.Join(assetsPath, "fonts", "prstart.ttf"), 10)
	if err != nil {
		panic(err)
	}
	atlas := text.NewAtlas(face, text.ASCII)
	s := &score{
		atlas:  atlas,
		canvas: scoreCanvas,
	}
	return s
}

func (g *game) drawScore() {
	g.score.canvas.Clear(color.RGBA{30, 30, 30, 255})
	moveText := pixel.V(185, 190)
	txt := text.New(moveText, g.score.atlas)
	txt.LineHeight = g.score.atlas.LineHeight() * 4.5
	fmt.Fprintln(txt, "Score")
	mat := pixel.IM.Moved(pixel.V(gameH+gameW, 0))
	g.score.canvas.Draw(g.window, mat)
	d := pixel.V(0, -43)
	move := moveText.Add(g.window.Bounds().Center())
	for _, player := range g.players {
		if player == nil {
			continue
		}
		move = move.Add(d)
		txt.Dot.X += 40
		sprite, color := g.getTankVisuals(player.name)
		sprite.DrawColorMask(g.window, pixel.IM.Moved(move), color)
		fmt.Fprintln(txt, player.score)
	}
	txt.Draw(g.window, pixel.IM.Moved(g.window.Bounds().Center()))

	if g.lastWinner == "" {
		return 
	}

	moveText = pixel.V(185, -120)
	winnerTxt := text.New(moveText, g.score.atlas)
	winnerTxt.Dot.X += 7
	fmt.Fprintln(winnerTxt, "Last")
	winnerTxt.Dot.X -= 2
	winnerTxt.Dot.Y -= 8
	fmt.Fprintln(winnerTxt, "winner")
	move = moveText.Add(g.window.Bounds().Center()).Add(pixel.V(25, -45))
	sprite, color := g.getTankVisuals(g.lastWinner)
	sprite.DrawColorMask(g.window, pixel.IM.Moved(move), color)
	winnerTxt.Draw(g.window, pixel.IM.Moved(g.window.Bounds().Center()))
	
}

package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/gif"
	"math/rand"
	"path/filepath"
	"time"
)

type game struct {
	sprites   *sprites
	titleSize int
	window    *pixelgl.Window
	canvas    *pixelgl.Canvas
	imd       *imdraw.IMDraw
}

func NewGame(assetsPath string) (g *game) {
	sprites := loadSprites(filepath.Join(assetsPath, "sprites"))

	windowBounds := pixel.Rect{Max: pixel.Vec{X: 480, Y: 416}}
	cfg := pixelgl.WindowConfig{
		Title:  "Battle City",
		Bounds: windowBounds,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	canvas := pixelgl.NewCanvas(windowBounds)
	canvas.SetMatrix(pixel.IM.Scaled(pixel.ZV, 2))

	g = &game{
		sprites:   sprites,
		titleSize: 16,
		window:    window,
		canvas:    canvas,
	}
	return g
}

func (g *game) Run() {
	rand.Seed(time.Now().UnixNano())
	fps := 60
	fpsSync := time.Tick(time.Second / time.Duration(fps))
	for !g.window.Closed() {
		ctrl := pixel.ZV
		if g.window.Pressed(pixelgl.KeyLeft) {
			ctrl.X--
		}
		if g.window.Pressed(pixelgl.KeyRight) {
			ctrl.X++
		}
		if g.window.JustPressed(pixelgl.KeyUp) {
			ctrl.Y = 1
		}

		g.window.Clear(colornames.White)
		g.canvas.Clear(colornames.Black)

		g.sprites.arrows[1].Draw(g.canvas, pixel.IM.Moved(g.sprites.arrows[1].Frame().Size().Scaled(0.5)))

		g.canvas.Draw(g.window, pixel.IM.Moved(g.canvas.Bounds().Center()))
		g.window.Update()
		<-fpsSync
	}
}

package game

import (
	_ "image/gif"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type game struct {
	sprites   *sprites
	titleSize int
	window    *pixelgl.Window
	canvas    *pixelgl.Canvas
	levels 	  [][26][26]byte
}

func NewGame(assetsPath string) (g *game) {
	sprites := loadSprites(filepath.Join(assetsPath, "sprites"))
	levels := loadLevels(filepath.Join(assetsPath, "levels"))

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
		levels:    levels,
	}
	return g
}

func (g *game) Run() {
	rand.Seed(time.Now().UnixNano())
	fps := 30
	fpsSync := time.Tick(time.Second / time.Duration(fps))
	for !g.window.Closed() {
		ctrl := pixel.ZV
		if g.window.Pressed(pixelgl.KeyA) {
			ctrl.X--
		}
		if g.window.Pressed(pixelgl.KeyD) {
			ctrl.X++
		}
		if g.window.JustPressed(pixelgl.KeyW) {
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

package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/kochetov-dmitrij/battle-city-ds/game"
)

func run() {
	g := game.NewGame("./assets")
	g.Run()
}

func main() {
	pixelgl.Run(run)
}

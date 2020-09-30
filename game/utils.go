package game

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/gif"
	"os"
	"path/filepath"
)

type sprites struct {
	player *pixel.Sprite
}

func loadSprites(spritesPath string) *sprites {
	spritesFile, err := os.Open(filepath.Join(spritesPath, "sprites.gif"))
	if err != nil {
		panic(err)
	}
	defer spritesFile.Close()
	spritesImg, _, err := image.Decode(spritesFile)
	if err != nil {
		panic(err)
	}
	spriteSheet := pixel.PictureDataFromImage(spritesImg)
	return &sprites{
		player: pixel.NewSprite(spriteSheet, pixel.R(32, 96-15, 32+15, 96)),
	}
}

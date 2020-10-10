package game

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/gif"
	"os"
	"path/filepath"
)

type sprites struct {
	player, flag          *pixel.Sprite
	playerLife, enemyLife *pixel.Sprite
	enemies               []*pixel.Sprite
	arrows                []*pixel.Sprite
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
		player:     pixel.NewSprite(spriteSheet, pixel.R(0, 99, 13, 112)),
		flag:       pixel.NewSprite(spriteSheet, pixel.R(64, 48, 80, 63)),
		playerLife: pixel.NewSprite(spriteSheet, pixel.R(89, 48, 96, 56)),
		enemyLife:  pixel.NewSprite(spriteSheet, pixel.R(81, 48, 88, 55)),
		enemies: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(32, 97, 45, 112)),
			pixel.NewSprite(spriteSheet, pixel.R(48, 97, 61, 112)),
			pixel.NewSprite(spriteSheet, pixel.R(64, 97, 77, 112)),
			pixel.NewSprite(spriteSheet, pixel.R(80, 97, 93, 112)),
		},
		arrows: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(81, 57, 88, 64)),
			pixel.NewSprite(spriteSheet, pixel.R(88, 57, 95, 64)),
		},
	}
}

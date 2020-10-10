package game

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/gif"
	"os"
	"path/filepath"
)

type sprites struct {
	flag, bullet 		  *pixel.Sprite
	playerLife, enemyLife *pixel.Sprite
	enemies               []*pixel.Sprite
	players               []*pixel.Sprite
	arrows                []*pixel.Sprite
	explosions			  []*pixel.Sprite
	tiles				  []*pixel.Sprite
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
		flag:       pixel.NewSprite(spriteSheet, pixel.R(64, 48, 80, 63)),
		bullet:		pixel.NewSprite(spriteSheet, pixel.R(75, 34, 78, 38)),
		playerLife: pixel.NewSprite(spriteSheet, pixel.R(89, 48, 96, 56)),
		enemyLife:  pixel.NewSprite(spriteSheet, pixel.R(81, 48, 88, 55)),
		players: 	[]*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(0, 99, 13, 112)),
			pixel.NewSprite(spriteSheet, pixel.R(16, 99, 29, 112)),
		},
		arrows: 	[]*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(81, 57, 88, 64)),
			pixel.NewSprite(spriteSheet, pixel.R(88, 57, 95, 64)),
		},
		explosions: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(0, 0, 32, 32)),
			pixel.NewSprite(spriteSheet, pixel.R(32, 0, 64, 32)),	
			pixel.NewSprite(spriteSheet, pixel.R(64, 0, 96, 32)),	
		},
		tiles:		[]*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(48, 40, 56, 48)),
			pixel.NewSprite(spriteSheet, pixel.R(48, 32, 56, 40)),
			pixel.NewSprite(spriteSheet, pixel.R(56, 32, 64, 40)),
			pixel.NewSprite(spriteSheet, pixel.R(64, 40, 72, 48)),
			pixel.NewSprite(spriteSheet, pixel.R(72, 40, 80, 48)),
			pixel.NewSprite(spriteSheet, pixel.R(64, 32, 72, 40)),
		},
	}
}
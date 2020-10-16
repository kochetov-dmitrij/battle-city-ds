package game

import (
	"bufio"
	"image"
	_ "image/gif"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type sprites struct {
	flag, bullet          *pixel.Sprite
	playerLife, enemyLife *pixel.Sprite
	enemies               []*pixel.Sprite
	players               []*pixel.Sprite
	arrows                []*pixel.Sprite
	explosions            []*pixel.Sprite
	spawns                []*pixel.Sprite
	tiles                 map[byte]*pixel.Sprite
	sheet                 *pixel.PictureData
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
		bullet:     pixel.NewSprite(spriteSheet, pixel.R(75, 34, 78, 38)),
		playerLife: pixel.NewSprite(spriteSheet, pixel.R(89, 48, 96, 56)),
		enemyLife:  pixel.NewSprite(spriteSheet, pixel.R(81, 48, 88, 55)),
		players: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(0, 99, 13, 112)),
			pixel.NewSprite(spriteSheet, pixel.R(16, 99, 29, 112)),
		},
		arrows: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(81, 57, 88, 64)),
			pixel.NewSprite(spriteSheet, pixel.R(88, 57, 95, 64)),
		},
		explosions: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(0, 0, 32, 32)),
			pixel.NewSprite(spriteSheet, pixel.R(32, 0, 64, 32)),
			pixel.NewSprite(spriteSheet, pixel.R(64, 0, 96, 32)),
		},
		tiles: map[byte]*pixel.Sprite{
			// tileEmpty: pixel.NewSprite(spriteSheet, pixel.R(0, 32, 32, 64)),
			tileBrick: pixel.NewSprite(spriteSheet, pixel.R(48, 40, 56, 48)),
			tileSteel: pixel.NewSprite(spriteSheet, pixel.R(48, 32, 56, 40)),
			tileGrass: pixel.NewSprite(spriteSheet, pixel.R(56, 32, 64, 40)),
			tileWater: pixel.NewSprite(spriteSheet, pixel.R(64, 40, 72, 48)),
			'w':       pixel.NewSprite(spriteSheet, pixel.R(72, 40, 80, 48)),
			tileFroze: pixel.NewSprite(spriteSheet, pixel.R(64, 32, 72, 40)),
		},
		spawns: []*pixel.Sprite{
			pixel.NewSprite(spriteSheet, pixel.R(32, 48, 48, 64)),
			pixel.NewSprite(spriteSheet, pixel.R(48, 48, 64, 64)),
		},
		sheet: spriteSheet,
	}
}

func (g *game) loadTankSprite(number byte) *pixel.Sprite {
	if number%2 == 0 {
		return pixel.NewSprite(g.sprites.sheet, pixel.R(0, 99, 13, 112))
	}
	return pixel.NewSprite(g.sprites.sheet, pixel.R(16, 99, 29, 112))
}

func loadLevels(levelsPath string) [][26][26]byte {
	var levels [][26][26]byte
	err := filepath.Walk(levelsPath, func(path string, info os.FileInfo, err error) error {
		if levelsPath == path {
			return nil
		}
		levels = append(levels, loadLevel(path))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return levels
}

func loadLevel(levelPath string) [26][26]byte {
	levelFile, err := os.Open(levelPath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(levelFile)
	level := [26][26]byte{}
	for i := range level {
		for j := range level[i] {
			tile, err := reader.ReadByte()
			if err != nil {
				panic(err)
			}
			level[i][j] = tile
		}
		reader.ReadByte()
	}
	return level
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func rectContains(r1 pixel.Rect, r2 pixel.Rect) bool {
	return r1.Min.X < r2.Min.X && r1.Min.Y < r2.Min.Y &&
		r1.Max.X > r2.Max.X && r1.Max.Y > r2.Max.Y
}

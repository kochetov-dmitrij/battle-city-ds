package game

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kochetov-dmitrij/battle-city-ds/connection"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	_ "image/gif"
	"log"
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
	score     *score
	levels    [][26][26]byte
	world     *world
	players   [4]*player
	peers     connection.Peers
}

const (
	// gameW = 240
	gameW = 250
	gameH = 208
)

func NewGame(assetsPath string) (g *game) {
	peers := connection.Peers{}
	go connection.Connection(peers, &pb.ComsService{AddMessage: g.AddMessage})

	sprites := loadSprites(filepath.Join(assetsPath, "sprites"))
	levels := loadLevels(filepath.Join(assetsPath, "levels"))

	windowBounds := pixel.Rect{Max: pixel.Vec{X: 2 * gameW, Y: 2 * gameH}}
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
		world:     &world{},
		players:   [4]*player{nil, nil, nil, nil},
		score:     g.initScore(assetsPath),
		peers:     peers,
	}
	return g
}

func (g *game) AddMessage(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	//todo This function is called by other peers to change the state of THIS peer
	//log.Printf("<<<--- received - %s", msg.BulletDirection)
	return &empty.Empty{}, nil
}

func (g *game) Run() {
	rand.Seed(time.Now().UnixNano())
	fps := 10
	fpsSync := time.Tick(time.Second / time.Duration(fps))

	direction := up
	moves := false
	localPlayer := g.loadPlayer("default")
	g.world.worldMap = g.levels[0] // TODO change to loading from menu
	for i := 1; i < 4; i++ {
		g.loadPlayer(string(i))
	}

	for !g.window.Closed() {

		for peerAddress, client := range g.peers {
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			_, err := client.AddMessage(ctx, &pb.Message{
				//todo This calls AddMessage() of all other peers and passes pb.Message
				BulletDirection: pb.Message_RIGHT,
			})
			if err != nil {
				delete(g.peers, peerAddress)
				log.Printf("Peer %s disconnected | %v\n", peerAddress, err)
			}
		}

		moves = false
		if g.window.Pressed(pixelgl.KeyA) {
			direction = left
			moves = true
		}
		if g.window.Pressed(pixelgl.KeyD) {
			direction = right
			moves = true
		}
		if g.window.Pressed(pixelgl.KeyW) {
			direction = up
			moves = true
		}
		if g.window.Pressed(pixelgl.KeyS) {
			direction = down
			moves = true
		}
		if g.window.JustPressed(pixelgl.KeyF) {
			localPlayer.tank.fire(g)
			for _, player := range g.players {
				if player != localPlayer && player != nil {
					player.tank.fire(g)
				}
			}
		}
		if g.window.Pressed(pixelgl.KeyS) {
			direction = down
			moves = true
		}
		// if g.window.JustPressed(pixelgl.KeySpace) {
		// 	playerTank.fire()
		// }

		g.window.Clear(colornames.White)
		g.canvas.Clear(colornames.Black)
		g.draw()

		// last := time.Since(last).Milliseconds()
		for _, player := range g.players {
			if player != nil {
				g.updatePlayer(player, direction, moves)
			}
		}
		// g.sprites.arrows[1].Draw(g.canvas, pixel.IM.Moved(g.sprites.arrows[1].Frame().Size().Scaled(0.5)))

		// g.sprites.tiles[tileEmpty].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(0.5)))
		// g.sprites.tiles[tileBrick].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(1)))
		// g.sprites.tiles[tileSteel].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(2)))
		// g.sprites.tiles[tileWater].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(3)))
		// g.sprites.tiles[tileFroze].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(4)))
		// g.sprites.tiles[tileGrass].Draw(g.canvas, pixel.IM.Moved(g.sprites.tiles[tileEmpty].Frame().Size().Scaled(5)))
		g.canvas.Draw(g.window, pixel.IM.Moved(g.canvas.Bounds().Center()))
		g.drawScore()
		g.window.Update()
		<-fpsSync
	}
}

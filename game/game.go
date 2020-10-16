package game

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kochetov-dmitrij/battle-city-ds/connection"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	_ "image/gif"
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
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
	port      string
}

const (
	gameW      = 250
	gameH      = 208
	maxPlayers = 4
)

func NewGame(assetsPath string) (g *game) {
	peers := connection.Peers{}
	rand.Seed(time.Now().UTC().UnixNano())
	myPort := strconv.Itoa(rand.Intn(13000-12000) + 12000)

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
		port:      myPort,
	}

	go connection.Connection(peers, myPort, &pb.ComsService{AddMessage: g.AddMessage})

	return g
}

func (g *game) AddMessage(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	i := 0
	lastNil := -1
	fmt.Println("Peer ", g.port, ". Receiving message from ", msg.GetHost())
	for ; i < maxPlayers; i++ {
		if g.players[i] == nil {
			lastNil = i
			continue
		}
		if g.players[i].name == msg.GetHost() {
			break
		}
	}
	fmt.Println("Peer ", g.port, ". Trying to work with ", msg.GetHost())
	if i == maxPlayers || g.players[i] == nil || g.players[i].name != msg.GetHost() {
		if lastNil != -1 {
			fmt.Println("Peer ", g.port, ". Adding new player  ", msg.GetHost())
			g.players[lastNil] = g.loadPlayer(msg.GetHost(), false)
			fmt.Println("Peer ", g.port, ". Added new player  ", msg.GetHost())
			i = lastNil
		}
	}

	g.players[i].tank.state = State(msg.GetTankState())
	positionT := msg.GetTankPosition()
	g.players[i].tank.x = int64(positionT.X)
	g.players[i].tank.y = int64(positionT.Y)
	g.players[i].tank.direction = Direction(msg.GetAction()[0].TankDirection - 1)
	if msg.GetBulletState() == removed {
		g.players[i].tank.bullet = nil
		return &empty.Empty{}, nil
	}

	state := State(msg.GetBulletState())
	direction := Direction(msg.GetBulletDirection() - 1)
	positionB := msg.GetBulletPosition()
	x, y := int64(positionB.X), int64(positionB.Y)
	g.players[i].tank.bullet = g.loadBullet(x, y, direction, state)
	return &empty.Empty{}, nil
}

func (g *game) Run() {
	rand.Seed(time.Now().UnixNano())
	fps := 10
	fpsSync := time.Tick(time.Second / time.Duration(fps))

	direction := up
	moves := false
	localPlayer := g.loadPlayer(g.port, true)
	g.world.worldMap = g.levels[0] // TODO change to loading from menu

	for !g.window.Closed() {
		message := &pb.Message{
			Host:         localPlayer.name,
			TankPosition: &pb.Message_TankPosition{X: uint32(localPlayer.tank.x), Y: uint32(localPlayer.tank.y)},
			TankState:    uint32(localPlayer.tank.state),
			Action:		  []*pb.Message_Action { &pb.Message_Action {TankDirection: pb.Message_Direction(localPlayer.tank.direction + 1)}},
			BulletState:  uint32(removed),
			//todo This calls AddMessage() of all other peers and passes pb.Message
		}
		if localPlayer.tank.bullet != nil {
			x, y := localPlayer.tank.bullet.x, localPlayer.tank.bullet.y
			message.BulletDirection = pb.Message_Direction(localPlayer.tank.bullet.direction + 1)
			message.BulletPosition = &pb.Message_BulletPosition{X: uint32(x), Y: uint32(y)}
			message.BulletState = uint32(localPlayer.tank.bullet.state)
		}

		for peerAddress, client := range g.peers {
			fmt.Println("Peer ", g.port, ". Trying to send info to ", peerAddress)
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			fmt.Println("Peer ", g.port, ". Tried to send info to ", peerAddress)
			_, err := client.AddMessage(ctx, message)
			if err != nil {
				fmt.Println("Peer ", g.port, ". Wow an ERROR while sending info to ", peerAddress)
				port := regexp.MustCompile("http:.*:(.*)").FindStringSubmatch(peerAddress)[1]
				for i := 0; i < maxPlayers; i++ {
					if g.players[i] != nil && g.players[i].name == port {
						g.players[i] = nil
						break
					}
				}
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
		}

		g.window.Clear(colornames.White)
		g.canvas.Clear(colornames.Black)
		g.draw()
		g.updatePlayer(localPlayer, direction, moves)
		for i := 0; i < maxPlayers; i++ {
			if g.players[i] != nil && g.players[i] != localPlayer {
				g.updatePlayer(g.players[i], g.players[i].tank.direction, false)
			}
		}

		g.canvas.Draw(g.window, pixel.IM.Moved(g.canvas.Bounds().Center()))
		g.drawScore()
		g.window.Update()
		<-fpsSync
	}
}

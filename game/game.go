package game

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kochetov-dmitrij/battle-city-ds/connection"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	_ "image/gif"
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type game struct {
	sprites    *sprites
	titleSize  int
	window     *pixelgl.Window
	canvas     *pixelgl.Canvas
	score      *score
	levels     [][][]byte
	world      *world
	players    [4]*player
	peers      connection.Peers
	port       string
	address    string
	lastWinner string
}

const (
	gameW        = 250
	gameH        = 208
	maxPlayers   = 4
	maxScore     = 10
	defaultLevel = 0
)

func NewGame(assetsPath string) (g *game) {
	g = &game{
		titleSize: 16,
		world:     &world{},
		players:   [4]*player{nil, nil, nil, nil},
	}
	peers, myAddress, myPort := connection.Connection(&pb.ComsService{AddMessage: g.AddMessage})

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

	g.sprites = sprites
	g.window = window
	g.canvas = canvas
	g.levels = levels
	g.peers = peers
	g.port = myPort
	g.address = myAddress
	g.score = g.initScore(assetsPath)
	g.lastWinner = ""
	return g
}

func (g *game) AddMessage(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	for _, peer := range msg.AllPeers {
		g.peers.Add(peer, g.address)
	}

	i := 0
	firstNil := -1
	//fmt.Println("Peer ", g.port, ". Receiving message from ", msg.GetHost())
	for ; i < maxPlayers; i++ {
		if g.players[i] == nil {
			if firstNil == -1 {
				firstNil = i
			}
			continue
		}
		if g.players[i].name == msg.GetHost() {
			break
		}
	}
	//fmt.Println("Peer ", g.port, ". Trying to work with ", msg.GetHost())
	if i == maxPlayers || g.players[i] == nil || g.players[i].name != msg.GetHost() {
		if firstNil != -1 {
			//fmt.Println("Peer ", g.port, ". Adding new player  ", msg.GetHost())
			g.players[firstNil] = g.loadPlayer(msg.GetHost(), false)
			//fmt.Println("Peer ", g.port, ". Added new player  ", msg.GetHost())
			i = firstNil
		}
	}

	if g.lastWinner == "" {
		g.lastWinner = msg.GetLastWinner()
	}

	for i := range g.world.worldMap {
		for j := range g.world.worldMap[i] {
			if msg.LevelState[i][j] == 46 && g.world.worldMap[i][j] != 46 {
				g.world.worldMap[i][j] = 46
			}
		}
	}

	scores := msg.GetScore()

	for _, player := range g.players {
		if player == nil {
			continue
		}
		for name, score := range scores {
			if player.name == name && byte(score) >= player.score {
				if score >= maxScore {
					g.lastWinner = player.name
					g.resetMap()
					return &empty.Empty{}, nil
				}
				player.score = byte(score)
			}
		}

	}

	g.players[i].tank.state = State(msg.GetTankState())
	positionT := msg.GetTankPosition()
	g.players[i].tank.x = int64(positionT.X)
	g.players[i].tank.y = int64(positionT.Y)
	g.players[i].tank.direction = Direction(msg.GetTankDirection() - 1)
	if msg.GetBulletState() == removed {
		g.players[i].tank.bullet = nil
		return &empty.Empty{}, nil
	}

	state := State(msg.GetBulletState())
	direction := Direction(msg.GetBulletDirection() - 1)
	positionB := msg.GetBulletPosition()
	x, y := int64(positionB.X), int64(positionB.Y)
	g.players[i].tank.bullet = g.loadBullet(x, y, direction, state, g.players[i].tank)

	return &empty.Empty{}, nil
}

func (g *game) LoadMap(lvl int) {
	g.world.worldMap = make([][]byte, 26)
	for rowIdx := range g.levels[lvl] {
		g.world.worldMap[rowIdx] = make([]byte, 26)
		copy(g.world.worldMap[rowIdx], g.levels[lvl][rowIdx])
	}
}

func (g *game) Run() {
	rand.Seed(time.Now().UnixNano())
	fps := 10
	fpsSync := time.Tick(time.Second / time.Duration(fps))

	direction := up
	moves := false
	localPlayer := g.loadPlayer(g.port, true)
	g.LoadMap(defaultLevel)

	for !g.window.Closed() {

		scores := map[string]uint32{}
		for _, player := range g.players {
			if player != nil {
				scores[player.name] = uint32(player.score)
			}
		}

		if localPlayer.score == maxScore {
			g.resetMap()
			g.lastWinner = localPlayer.name
		}

		message := &pb.Message{
			Host:          localPlayer.name,
			TankPosition:  &pb.Message_Position{X: uint32(localPlayer.tank.x), Y: uint32(localPlayer.tank.y)},
			TankState:     uint32(localPlayer.tank.state),
			TankDirection: pb.Message_Direction(localPlayer.tank.direction + 1),
			BulletState:   uint32(removed),
			AllPeers:      append(g.peers.GetList(), g.address),
			LevelState:    g.world.worldMap,
			Score:         scores,
			LastWinner:    g.lastWinner,
		}

		if localPlayer.tank.bullet != nil {
			x, y := localPlayer.tank.bullet.x, localPlayer.tank.bullet.y
			message.BulletDirection = pb.Message_Direction(localPlayer.tank.bullet.direction + 1)
			message.BulletPosition = &pb.Message_Position{X: uint32(x), Y: uint32(y)}
			message.BulletState = uint32(localPlayer.tank.bullet.state)
		}

		for peerAddress, client := range g.peers {
			//fmt.Println("Peer ", g.port, ". Trying to send info to ", peerAddress)
			ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
			//fmt.Println("Peer ", g.port, ". Tried to send info to ", peerAddress)

			// This calls AddMessage() of all other peers and passes pb.Message
			_, err := client.AddMessage(ctx, message)
			if err != nil && !errors.Is(err, context.DeadlineExceeded) {
				//fmt.Println("Peer ", g.port, ". Wow an ERROR while sending info to ", peerAddress)
				port := regexp.MustCompile("http:.*:(.*)").FindStringSubmatch(peerAddress)[1]
				for i := 0; i < maxPlayers; i++ {
					if g.players[i] != nil && g.players[i].name == port {
						g.players[i] = nil
						break
					}
				}
				g.peers.Remove(peerAddress)
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

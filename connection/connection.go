package connection

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/huin/goupnp"
	"github.com/kochetov-dmitrij/battle-city-ds/connection/pb"
	"github.com/koron/go-ssdp"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"time"
)

func getMyIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

type Peers map[string]pb.ComsClient

type connectorP2P struct {
	usn       string
	myAddress string
	peers     Peers
}

func (c *connectorP2P) advertise() {
	ad, err := ssdp.Advertise(
		c.usn,
		c.usn,
		c.myAddress,
		"",
		1800)
	if err != nil {
		panic(err)
	}

	aliveTick := time.Tick(1 * time.Second)

	for {
		select {
		case <-aliveTick:
			_ = ad.Alive()
		}
	}
}

func (c *connectorP2P) discover() {
	discoverTick := time.Tick(2 * time.Second)
	for {
		select {
		case <-discoverTick:
			response, err := goupnp.DiscoverDevices(c.usn)
			if err != nil {
				panic(err)
			}
			for _, r := range response {
				peerAddress := r.Location.String()
				if _, alreadyDiscovered := c.peers[peerAddress]; !alreadyDiscovered && peerAddress != c.myAddress {
					peerAddressWithoutHTP := regexp.MustCompile("^http://").ReplaceAllString(peerAddress, "")
					conn, err := grpc.Dial(peerAddressWithoutHTP, grpc.WithInsecure())
					if err != nil {
						log.Fatalf("Did not connect: %v", err)
					}
					c.peers[peerAddress] = pb.NewComsClient(conn)
				}
			}
		}
		log.Printf("Found peers: %v", c.peers)
		for peerAddress, client := range c.peers {
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			_, err := client.AddMessage(ctx, &pb.Message{BulletDirection: pb.Message_RIGHT})
			if err != nil {
				log.Printf("Peer %s disconnected | %v\n", peerAddress, err)
				delete(c.peers, peerAddress)
			}
		}
	}
}

func startGRPCServer(port string) {
	lis, _ := net.Listen("tcp", ":"+port)
	server := grpc.NewServer()
	log.Printf("Launched a gRPC server on port %s", port)
	pb.RegisterComsService(server, &pb.ComsService{AddMessage: AddMessage})
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func Connection(peers Peers) {
	rand.Seed(time.Now().UTC().UnixNano())

	myPort := strconv.Itoa(rand.Intn(13000-12000) + 12000)
	connectorP2P := &connectorP2P{
		myAddress: fmt.Sprintf("http://%s:%s", getMyIP().String(), myPort),
		usn:       "game:battle-city-ds",
		peers:     peers,
	}
	fmt.Printf("My connection details: %+v\n", *connectorP2P)

	startGRPCServer(myPort)
	go connectorP2P.discover()
	connectorP2P.advertise()
}

func AddMessage(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	log.Printf("<<<--- received - %s", msg.BulletDirection)
	return &empty.Empty{}, nil
}

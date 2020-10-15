package connection

import (
	"fmt"
	"github.com/huin/goupnp"
	"github.com/koron/go-ssdp"
	"log"
	"math/rand"
	"net"
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

type Peers map[string]struct{}

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
				_ = r.Location.String()
				peerAddress := r.Location.String()
				if _, alreadyDiscovered := c.peers[peerAddress]; !alreadyDiscovered && peerAddress != c.myAddress {
					c.peers[peerAddress] = struct{}{}
				}
			}
		}
		log.Printf("Found peers: %v", c.peers)
	}
}

func Connection(peers Peers) {
	rand.Seed(time.Now().UTC().UnixNano())

	connectorP2P := &connectorP2P{
		myAddress: fmt.Sprintf("http://%s:%d", getMyIP().String(), rand.Intn(13000-12000)+12000),
		usn:       "game:battle-city-ds",
		peers:     peers,
	}
	fmt.Printf("My connection details: %+v\n", *connectorP2P)

	go connectorP2P.discover()
	connectorP2P.advertise()
}

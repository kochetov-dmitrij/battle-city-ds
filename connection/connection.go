package connection

import (
	"fmt"
	"github.com/huin/goupnp"
	"github.com/koron/go-ssdp"
	"log"
	"math/rand"
	"net"
	"sync"
)

type connectionDetails struct {
	ip   string
	port int
	st   string
}

func advertise(cd connectionDetails, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := ssdp.Advertise(
		cd.st,
		"",
		fmt.Sprintf("http://%s:%d", cd.ip, cd.port),
		"",
		1800)
	if err != nil {
		panic(err)
	}
}

func discover(cd connectionDetails, wg *sync.WaitGroup) {
	defer wg.Done()

	r, err := goupnp.DiscoverDevices(cd.st)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", r)
}

func getMyIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func Connection() {
	var wg sync.WaitGroup
	wg.Add(2)

	cd := connectionDetails{
		ip:   getMyIP().String(),
		port: rand.Intn(13000-12000) + 12000,
		st:   "battle-city-ds",
	}

	go advertise(cd, &wg)
	go discover(cd, &wg)

	wg.Wait()
}

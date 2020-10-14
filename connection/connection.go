package connection

import (
	"fmt"
	"github.com/huin/goupnp"
	"github.com/huin/goupnp/httpu"
	"github.com/huin/goupnp/ssdp"
	"sync"
)

var wg sync.WaitGroup

func Advertise() {
	client, _ := httpu.NewHTTPUClient()
	r, _ := ssdp.SSDPRawSearch(client, ssdp.SSDPAll, 5, 9999)

	fmt.Printf("%v\n", r)
	wg.Done()
}

func Discover() {

	fmt.Printf("Lol kakoi discover ..\n")

	r, err := goupnp.DiscoverDevices(ssdp.SSDPAll)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", r)

	wg.Done()
}

func Connection() {

	wg.Add(2)

	go Advertise()
	go Discover()

	wg.Wait()
}

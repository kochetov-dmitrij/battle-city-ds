package connection

import (
	"fmt"
	"github.com/huin/goupnp"
	//"github.com/huin/goupnp/ssdp"
	"github.com/koron/go-ssdp"
	"sync"
)

var wg sync.WaitGroup

func Advertise() {

	_, err := ssdp.Advertise(
		"my:device",                        // send as "ST"
		"unique:id",                        // send as "USN"
		"http://192.168.0.1:57086/foo.xml", // send as "LOCATION"
		"go-ssdp sample",                   // send as "SERVER"
		1800)
	if err != nil {
		panic(err)
	}

	wg.Done()
}

func Discover() {

	fmt.Printf("Lol kakoi discover ..\n")

	r, err := goupnp.DiscoverDevices("ssdp:all")
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

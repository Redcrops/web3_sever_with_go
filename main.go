package main

import (
	"time"

	"github.com/Redcrops/web3_sever_with_go/network"
)

// server
// transport => tcp,udp
// block
// tx
// keypair
func main() {
	// fmt.Println("Hello, Web3 with Go!")
	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}

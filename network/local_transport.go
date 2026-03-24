package network

import "sync"

type LocalTransport struct {
	addr NetAddr
	consumeCh chan RPC
	lock sync.RWMutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr: addr,
		consumeCh: make(chan RPC,1024),
		pees: make(map[NetAddr]*LoalTransport),
	}
}
package main

import (
	"log"

	"github.com/elmarsan/havel/protocol"
)

func main() {
	client := Client{
		version: 0xea62,
		net:     protocol.MainNet,
	}

	err := client.AddPeer("127.0.0.1:8333")
	if err != nil {
		log.Fatal(err)
	}
}

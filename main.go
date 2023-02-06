package main

import (
	"log"

	"github.com/elmarsan/havel/protocol"
)

// 1 - Connect to peer and send version message
// 2 - Receive it's version message and answer with verack
// 3 - Ask for a block or something

func main() {
	// 206.189.125.136:8332
	// 51.15.171.227:8333
	// 51.68.36.57:8333

	client := Client{
		version: 0xea62,
		net:     protocol.MainNet,
	}

	err := client.AddPeer("206.189.125.136:8332")
	if err != nil {
		log.Fatal(err)
	}
}

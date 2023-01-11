package main

// Last protocol version 70015

const (
	mainnet  uint32 = 0xD9B4BEF9
	testnet  uint32 = 0xDAB5BFFA
	testnet3 uint32 = 0x0709110B
)

const MsgHeaderSize = 24

const CommandSize = 12

type MsgHeader struct {
	net      uint32 // 4 bytes
	command  uint16 // 12 bytes
	length   uint32 // 4 bytes
	checksum uint32 // 4 bytes
}

const (
	Version = "version"
)

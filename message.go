package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Last protocol version 70015

const (
	mainnet  uint32 = 0xD9B4BEF9
	testnet  uint32 = 0xDAB5BFFA
	testnet3 uint32 = 0x0709110B
)

const MsgHeaderSize = 24

const CommandSize = 12

type MsgHeader struct {
	net      uint32
	command  uint64
	length   uint32
	checksum uint32
}

const (
	Version uint64 = 0x76657273696F6E00
)

func MsgHeaderFromRaw(raw []byte) (*MsgHeader, error) {
	var net uint32
	err := binary.Read(bytes.NewReader(raw), binary.LittleEndian, &net)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error reading msg net")
	}

	if net != mainnet && net != testnet && net != testnet3 {
		return nil, fmt.Errorf("%s", "Invalid net")
	}

	var command uint64
	err = binary.Read(bytes.NewReader(raw[4:16]), binary.BigEndian, &command)
	if err != nil {
		return nil, fmt.Errorf("%s", "Error reading msg command")
	}

	if command != Version {
		return nil, fmt.Errorf("Invalid command 0x%x", command)
	}

	return &MsgHeader{
		net:      net,
		command:  command,
		length:   0x00,
		checksum: 0x00,
	}, nil
}

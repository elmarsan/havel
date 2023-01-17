package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// https://en.bitcoin.it/wiki/Protocol_documentation#version
type MsgVersion struct {
	// Headers represents msg headers.
	Headers MsgHeader
	// Version represents the protocol version used by the node.
	Version uint32
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// AddRecv represents the network address of the node receiving this msg.
	// AddRecv [26]byte
	AddRecv net.IP
	// AddFrom represents the network address of the node sending this msg.
	// AddFrom net.Addr
	// Nonce represents random nonce generated every time a version msg is sent.
	Nonce uint64
	// UserAgent represents information of the node.
	// https://github.com/bitcoin/bips/blob/master/bip-0014.mediawiki
	UserAgent string
	// StartHeight represents the last block received by the emitting node.
	StartHeight uint32
	// Relay represents whether the remote peer should announce relayed transactions or not.
	// https://github.com/bitcoin/bips/blob/master/bip-0037.mediawiki
	Relay bool
}

// TODO: implement
func (msgv *MsgVersion) Encode(b *bytes.Buffer) error {
	return nil
}

// TODO: implement
func (msgv *MsgVersion) Decode(b *bytes.Buffer) error {
	headers := &MsgHeader{}
	err := headers.Decode(b)
	if err != nil {
		return err
	}

	version := binary.LittleEndian.Uint32(b.Bytes()[24:28])
	services := binary.LittleEndian.Uint64(b.Bytes()[28:36])
	unix := binary.LittleEndian.Uint64(b.Bytes()[36:44])
	addrRecv := b.Bytes()[52:70]
	nonce := binary.LittleEndian.Uint64(b.Bytes()[96:104])
	relay := b.Bytes()[b.Len()-1:][0]
	startHeight := binary.LittleEndian.Uint32(b.Bytes()[b.Len()-5 : b.Len()-1])

	msgv.Headers = *headers
	msgv.Version = version
	msgv.Services = services
	msgv.Timestamp = time.Unix(int64(unix), 0)
	msgv.AddRecv = net.IP(addrRecv[:net.IPv6len])
	msgv.Nonce = nonce
	msgv.StartHeight = startHeight

	if relay == 0x00 {
		msgv.Relay = false
	} else {
		msgv.Relay = true
	}

	fmt.Printf("Start height 0x%x\n", msgv.StartHeight)

	return nil
}

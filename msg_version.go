package main

import (
	"bytes"
	"encoding/binary"
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
	AddRecv net.Addr
	// AddFrom represents the network address of the node sending this msg.
	AddFrom net.Addr
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
	headers.Decode(b)

	version := binary.LittleEndian.Uint32(b.Bytes()[24:28])
	services := binary.LittleEndian.Uint64(b.Bytes()[28:36])
	unix := binary.LittleEndian.Uint64(b.Bytes()[36:44])

	msgv.Headers = *headers
	msgv.Version = version
	msgv.Services = services
	msgv.Timestamp = time.Unix(int64(unix), 0)

	return nil
}

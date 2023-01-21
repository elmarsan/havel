package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/elmarsan/havel/encode"
)

// https://en.bitcoin.it/wiki/Protocol_documentation#version
type MsgVersion struct {
	// Header represents msg header.
	Header *MsgHeader
	// Version represents the protocol version used by the node.
	Version uint32
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// Ip represents node's ip.
	RecvIP net.IP
	// RecvPort represent node's port.
	RecvPort uint16
	// FromIP represents node's ip.
	FromIP net.IP
	// FromPort represent node's port.
	FromPort uint16
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

// Encode encodes MsgVersion into w.
// TODO: implement
func (msgv *MsgVersion) Encode(w io.Writer) error {
	return nil
}

// Decode decodes MsgVersion from r.
// TODO: implement
func (msgv *MsgVersion) Decode(r io.Reader) error {
	// Decode headers
	msgv.Header = &MsgHeader{}
	err := msgv.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	// Decode body
	vals := []encode.DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &msgv.Version,
		},
		{
			Order: binary.LittleEndian,
			Val:   &msgv.Services,
		},
	}

	return encode.DecodeBatch(r, vals...)
}

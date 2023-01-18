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
	AddRecv net.Addr
	// AddFrom represents the network address of the node sending this msg.
	// AddFrom net.IP
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

// Encode encode MsgVersion in given Buffer.
// TODO: implement
func (msgv *MsgVersion) Encode(b *bytes.Buffer) error {
	return nil
}

// Decode decodes MsgVersion in given Buffer.
// [0:24] MsgHeaders
// [24:28] Version
// [28:36] Services
// [36:44] Timestamp
// [44:70] AddRecv
// 		[44:52] AddRecv Services
// 		[52:68] AddRecv IP
// 		[68:70] AddRecv Port

// [70:96] AddFrom
// 		[70:78] AddFrom Services
// 		[78:94] AddFrom IP
// 		[94:96] AddFrom Port

func (msgv *MsgVersion) Decode(b *bytes.Buffer) error {
	headers := &MsgHeader{}
	err := headers.Decode(b)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	msgv.Headers = *headers

	msgv.Version = binary.LittleEndian.Uint32(b.Bytes()[24:28])
	msgv.Services = binary.LittleEndian.Uint64(b.Bytes()[28:36])

	unix := binary.LittleEndian.Uint64(b.Bytes()[36:44])
	msgv.Timestamp = time.Unix(int64(unix), 0)

	msgNetAddr := &MsgNetAddr{}
	err = msgNetAddr.Decode(bytes.NewBuffer(b.Bytes()[96:104]))
	if err != nil {
		return fmt.Errorf("Could not decode AddRecv, cause %s", err.Error())
	}

	msgv.AddRecv = msgNetAddr.Addr
	msgv.Nonce = binary.LittleEndian.Uint64(b.Bytes()[96:104])

	return nil
}

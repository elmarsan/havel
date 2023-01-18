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
	// Decode headers
	headers := &MsgHeader{}
	err := headers.Decode(b)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	msgv.Headers = *headers

	// Decode version
	versionBytes := make([]byte, 4)
	_, err = b.Read(versionBytes)
	if err != nil {
		return fmt.Errorf("Could not decode version")
	}

	msgv.Version = binary.LittleEndian.Uint32(versionBytes)

	// Decode service
	serviceBytes := make([]byte, 8)
	_, err = b.Read(serviceBytes)
	if err != nil {
		return fmt.Errorf("Could not decode service")
	}

	msgv.Services = binary.LittleEndian.Uint64(serviceBytes)

	// Decode unix
	unixBytes := make([]byte, 8)
	_, err = b.Read(unixBytes)
	if err != nil {
		return fmt.Errorf("Could not decode timestamp")
	}

	unix := binary.LittleEndian.Uint64(unixBytes)
	msgv.Timestamp = time.Unix(int64(unix), 0)

	// Decode RecvIP and RecvPort
	msgNetAddr := &MsgNetAddr{}
	err = msgNetAddr.Decode(b)
	if err != nil {
		return fmt.Errorf("Could not decode AddRecv, cause %s", err.Error())
	}

	// Decode nonce
	nonceBytes := make([]byte, 8)
	_, err = b.Read(nonceBytes)
	if err != nil {
		return fmt.Errorf("Could not decode nonce")
	}

	msgv.Nonce = binary.LittleEndian.Uint64(nonceBytes)

	return nil
}

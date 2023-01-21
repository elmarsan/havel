package main

import (
	"encoding/binary"
	"fmt"
	"io"
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
	// RecvAddr represents receiver network address.
	RecvAddr *MsgNetAddr
	// FromAddr represents sender network address.
	FromAddr *MsgNetAddr
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
// TODO: finish implementation
func (msgv *MsgVersion) Decode(r io.Reader) error {
	// Decode headers
	msgv.Header = &MsgHeader{}
	err := msgv.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	var unix uint64

	// Decode version, services and timestamp
	vals := []encode.DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &msgv.Version,
		},
		{
			Order: binary.LittleEndian,
			Val:   &msgv.Services,
		},
		{
			Order: binary.LittleEndian,
			Val:   &unix,
		},
	}

	err = encode.DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	msgv.Timestamp = time.Unix(int64(unix), 0)

	// Decode RecvAddr
	recvAddr := &MsgNetAddr{}
	err = recvAddr.Decode(r)
	if err != nil {
		return err
	}

	// Decode FromAddr
	fromAddr := &MsgNetAddr{}
	err = fromAddr.Decode(r)
	if err != nil {
		return err
	}

	// Decode nonce
	vals = []encode.DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &msgv.Nonce,
		},
	}

	err = encode.DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	return nil
}

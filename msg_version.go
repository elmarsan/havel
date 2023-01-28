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
	UserAgent *MsgStr
	// StartHeight represents the last block received by the emitting node.
	StartHeight uint32
	// Relay represents whether the remote peer should announce relayed transactions or not.
	// https://github.com/bitcoin/bips/blob/master/bip-0037.mediawiki
	Relay bool
}

// Encode encodes MsgVersion into w.
func (version *MsgVersion) Encode(w io.Writer) error {
	err := version.Header.Encode(w)
	if err != nil {
		return fmt.Errorf("Unable to encode header, (%s)", err.Error())
	}

	// Encode Version, Services and Timestamp
	var unix uint64 = uint64(version.Timestamp.Unix())
	vals := []encode.EncodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &version.Version,
		},
		{
			Order: binary.LittleEndian,
			Val:   &version.Services,
		},
		{
			Order: binary.LittleEndian,
			Val:   &unix,
		},
	}

	err = encode.EncodeBatch(w, vals...)
	if err != nil {
		return err
	}

	// Encode RecvAddr
	err = version.RecvAddr.Encode(w)
	if err != nil {
		return err
	}

	// Encode FromAddr
	err = version.FromAddr.Encode(w)
	if err != nil {
		return err
	}

	// Encode Nonce
	err = encode.Encode(w, binary.LittleEndian, &version.Nonce)
	if err != nil {
		return err
	}

	// Encode user agent
	err = version.UserAgent.Encode(w)
	if err != nil {
		return err
	}

	// Encode StartHeight
	err = encode.Encode(w, binary.LittleEndian, &version.StartHeight)
	if err != nil {
		return err
	}

	return nil
}

// Decode decodes MsgVersion from r.
func (version *MsgVersion) Decode(r io.Reader) error {
	// Decode headers
	version.Header = &MsgHeader{}
	err := version.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Unable to decode header, (%s)", err.Error())
	}

	var unix uint64

	// Decode version, services and timestamp
	vals := []encode.DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &version.Version,
		},
		{
			Order: binary.LittleEndian,
			Val:   &version.Services,
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

	version.Timestamp = time.Unix(int64(unix), 0)

	// Decode RecvAddr
	recvAddr := &MsgNetAddr{}
	err = recvAddr.Decode(r)
	if err != nil {
		return err
	}

	version.RecvAddr = recvAddr

	// Decode FromAddr
	fromAddr := &MsgNetAddr{}
	err = fromAddr.Decode(r)
	if err != nil {
		return err
	}

	version.FromAddr = fromAddr

	// Decode nonce
	err = encode.Decode(r, binary.LittleEndian, &version.Nonce)
	if err != nil {
		return err
	}

	// Decode user agent
	userAgent := &MsgStr{}
	err = userAgent.Decode(r)
	if err != nil {
		return err
	}

	version.UserAgent = userAgent

	// Decode start height
	err = encode.Decode(r, binary.LittleEndian, &version.StartHeight)
	if err != nil {
		return err
	}

	return nil
}

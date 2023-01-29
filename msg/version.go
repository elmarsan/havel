package msg

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// https://en.bitcoin.it/wiki/Protocol_documentation#version
type Version struct {
	// Header represents msg header.
	Header *Header
	// Version represents the protocol version used by the node.
	Version uint32
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// RecvAddr represents receiver network address.
	RecvAddr *NetAddr
	// FromAddr represents sender network address.
	FromAddr *NetAddr
	// Nonce represents random nonce generated every time a version msg is sent.
	Nonce uint64
	// UserAgent represents information of the node.
	// https://github.com/bitcoin/bips/blob/master/bip-0014.mediawiki
	UserAgent *Str
	// StartHeight represents the last block received by the emitting node.
	StartHeight uint32
	// Relay represents whether the remote peer should announce relayed transactions or not.
	// https://github.com/bitcoin/bips/blob/master/bip-0037.mediawiki
	Relay bool
}

// Encode encodes Version into w.
func (version *Version) Encode(w io.Writer) error {
	err := version.Header.Encode(w)
	if err != nil {
		return fmt.Errorf("Unable to encode header, (%s)", err.Error())
	}

	// Encode Version, Services and Timestamp
	var unix uint64 = uint64(version.Timestamp.Unix())
	vals := []EncodeVal{
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

	err = EncodeBatch(w, vals...)
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
	err = Encode(w, binary.LittleEndian, &version.Nonce)
	if err != nil {
		return err
	}

	// Encode user agent
	err = version.UserAgent.Encode(w)
	if err != nil {
		return err
	}

	// Encode StartHeight
	err = Encode(w, binary.LittleEndian, &version.StartHeight)
	if err != nil {
		return err
	}

	return nil
}

// Decode decodes Version from r.
func (version *Version) Decode(r io.Reader) error {
	// Decode headers
	version.Header = &Header{}
	err := version.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Unable to decode header, (%s)", err.Error())
	}

	var unix uint64

	// Decode version, services and timestamp
	vals := []DecodeVal{
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

	err = DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	version.Timestamp = time.Unix(int64(unix), 0)

	// Decode RecvAddr
	recvAddr := &NetAddr{}
	err = recvAddr.Decode(r)
	if err != nil {
		return err
	}

	version.RecvAddr = recvAddr

	// Decode FromAddr
	fromAddr := &NetAddr{}
	err = fromAddr.Decode(r)
	if err != nil {
		return err
	}

	version.FromAddr = fromAddr

	// Decode nonce
	err = Decode(r, binary.LittleEndian, &version.Nonce)
	if err != nil {
		return err
	}

	// Decode user agent
	userAgent := &Str{}
	err = userAgent.Decode(r)
	if err != nil {
		return err
	}

	version.UserAgent = userAgent

	// Decode start height
	err = Decode(r, binary.LittleEndian, &version.StartHeight)
	if err != nil {
		return err
	}

	return nil
}

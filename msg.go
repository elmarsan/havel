package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/elmarsan/havel/encode"
)

// MsgHeader represents a default information contained in every protocol msg.
// https://en.bitcoin.it/wiki/Protocol_documentation#Message_structure
type MsgHeader struct {
	Magic    *BitcoinNet // Magic represents the value indicating the origin network.
	Cmd      *BitcoinCmd // Cmd represents type of p2p command.
	Length   uint32      // Lenght of payload in number of bytes.
	Checksum uint32      // Checksum holds first 4 bytes of sha256(sha256(payload)).
}

// Encode encodes MsgHeader into w.
func (msgHeader *MsgHeader) Encode(w io.Writer) error {
	magic := uint32(*msgHeader.Magic)
	cmd := make([]byte, 12)
	copy(cmd, msgHeader.Cmd.HexData[:])

	vals := []encode.EncodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &magic,
		},
		{
			Order: binary.BigEndian,
			Val:   &cmd,
		},
		{
			Order: binary.BigEndian,
			Val:   &msgHeader.Length,
		},
		{
			Order: binary.BigEndian,
			Val:   &msgHeader.Checksum,
		},
	}

	return encode.EncodeBatch(w, vals...)
}

// Decode decodes MsgHeader from r.
func (msgHeader *MsgHeader) Decode(r io.Reader) error {
	var magic uint32
	cmd := make([]byte, 12)

	vals := []encode.DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &magic,
		},
		{
			Order: binary.LittleEndian,
			Val:   &cmd,
		},
		{
			Order: binary.LittleEndian,
			Val:   &msgHeader.Length,
		},
		{
			Order: binary.BigEndian,
			Val:   &msgHeader.Checksum,
		},
	}

	err := encode.DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	// Convert magic to BitcoinNet
	btcNet, err := NewBitcoinNet(magic)
	if err != nil {
		return err
	}

	msgHeader.Magic = btcNet
	msgHeader.Cmd = &BitcoinCmd{}

	// Convert cmd to BitcoinCmd
	err = msgHeader.Cmd.FromHex(cmd)
	if err != nil {
		return err
	}

	return err
}

// MsgNetAddr represents network address of node.
// https://en.bitcoin.it/wiki/Protocol_documentation#Network_address
type MsgNetAddr struct {
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Ip represents node's ip.
	Ip net.IP
	// Port represent node's port.
	Port uint16
}

// Decode decodes MsgNetAddr contained in given buffer.
func (msgNetAddr *MsgNetAddr) Decode(b *bytes.Buffer) error {
	// TODO: Check size taking into account msg version does not include timestamp
	unixBytes := make([]byte, 4)
	_, err := b.Read(unixBytes)
	if err != nil {
		return fmt.Errorf("Could not decode timestamp")
	}

	unix := binary.LittleEndian.Uint32(unixBytes)
	msgNetAddr.Timestamp = time.Unix(int64(unix), 0)

	// TODO: Add service
	b.Next(8)

	ipBytes := make([]byte, 16)
	_, err = b.Read(ipBytes)
	if err != nil {
		return fmt.Errorf("Could not decode ip")
	}

	msgNetAddr.Ip = net.IP(ipBytes)

	portBytes := make([]byte, 2)
	_, err = b.Read(portBytes)
	if err != nil {
		return fmt.Errorf("Could not decode port")
	}

	msgNetAddr.Port = binary.BigEndian.Uint16(portBytes)

	return nil
}

// Encode encodes MsgNetAddr in given buffer.
func (msgNetAddr *MsgNetAddr) Encode(b *bytes.Buffer) error {
	return nil
}

package msg

import (
	"encoding/binary"
	"io"
	"net"
)

// NetAddr represents network address of node.
// https://en.bitcoin.it/wiki/Protocol_documentation#Network_address
type NetAddr struct {
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Ip represents node's ip.
	Ip net.IP
	// Port represent node's port.
	Port uint16
}

// Decode decodes NetAddr from r.
func (netAddr *NetAddr) Decode(r io.Reader) error {
	ip := make([]byte, 16)

	vals := []DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &netAddr.Services,
		},
		{
			Order: binary.BigEndian,
			Val:   &ip,
		},
		{
			Order: binary.BigEndian,
			Val:   &netAddr.Port,
		},
	}

	err := DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	netAddr.Ip = ip

	return nil
}

// Encode encodes NetAddr into w.
func (netAddr *NetAddr) Encode(w io.Writer) error {
	ip := make([]byte, 16)
	copy(ip[:], netAddr.Ip)

	vals := []EncodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &netAddr.Services,
		},
		{
			Order: binary.BigEndian,
			Val:   &ip,
		},
		{
			Order: binary.BigEndian,
			Val:   &netAddr.Port,
		},
	}

	return EncodeBatch(w, vals...)
}

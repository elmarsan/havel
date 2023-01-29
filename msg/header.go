package msg

import (
	"encoding/binary"
	"io"

	"github.com/elmarsan/havel/protocol"
)

// Header represents a default information contained in every protocol msg.
// https://en.bitcoin.it/wiki/Protocol_documentation#Message_structure
type Header struct {
	Magic    *protocol.BitcoinNet // Magic represents the value indicating the origin network.
	Cmd      *protocol.BitcoinCmd // Cmd represents type of p2p command.
	Length   uint32               // Lenght of payload in number of bytes.
	Checksum uint32               // Checksum holds first 4 bytes of sha256(sha256(payload)).
}

// Encode encodes Header into w.
func (header *Header) Encode(w io.Writer) error {
	magic := uint32(*header.Magic)
	cmd := make([]byte, 12)
	copy(cmd, header.Cmd.HexData[:])

	vals := []EncodeVal{
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
			Val:   &header.Length,
		},
		{
			Order: binary.LittleEndian,
			Val:   &header.Checksum,
		},
	}

	return EncodeBatch(w, vals...)
}

// Decode decodes Header from r.
func (header *Header) Decode(r io.Reader) error {
	var magic uint32
	cmd := make([]byte, 12)

	vals := []DecodeVal{
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
			Val:   &header.Length,
		},
		{
			Order: binary.LittleEndian,
			Val:   &header.Checksum,
		},
	}

	err := DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	// Convert magic to BitcoinNet
	btcNet, err := protocol.NewBitcoinNet(magic)
	if err != nil {
		return err
	}

	header.Magic = btcNet
	header.Cmd = &protocol.BitcoinCmd{}

	// Convert cmd to BitcoinCmd
	err = header.Cmd.FromHex(cmd)
	if err != nil {
		return err
	}

	return err
}
